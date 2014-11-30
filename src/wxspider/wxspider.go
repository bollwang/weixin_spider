package main

import (
	. "datas"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	. "tools"
)

func HandlerPage(url string, resultChan chan SpiderResult) {
	var sr SpiderResult
	page, err := DownloadRetry(url)
	log.Println("downlaod ret", len(page), err)
	if err == nil {
		log.Println("download succ, len=", len(page))
		sr, _ = ExtractPage(page)
		if err == nil {
			sr.Url = url
			sr.PageType = 7
			sr.CrawlerSource = "weixin_spdier"
			log.Println("page extract succ ", sr)
		} else {
			log.Println("page extract fail")
		}
	} else {
		log.Println("download fail ", url)
	}
	resultChan <- sr
}

func HandlerUrl(keyword string, resultChan chan []SpiderResult) {
	var sprs []SpiderResult

	indexurl := "http://weixin.sogou.com/weixin?query=" + keyword + "&_asf=www.sogou.com&_ast&ie=utf8&type=2"
	page, err := DownloadRetry(indexurl)
	if err != nil {
		log.Println("download from sogou fail")
		resultChan <- sprs
		return
	}
	log.Println("download succ, len=", len(page))

	urls, err := ExtractUrl(page)
	if err != nil {
		log.Println("extract from sogou fail")
		resultChan <- sprs
		return
	}
	log.Println("urls extract succ, len=", len(urls))

	channel := make(chan SpiderResult, len(urls))
	for _, url := range urls {
		go HandlerPage(url, channel)
	}
	for i := 0; i < len(urls); i++ {
		item := <-channel
		if strings.HasPrefix(item.Url, "http://mp.weixin.qq.com") {
			sprs = append(sprs, item)
		}
	}
	resultChan <- sprs
}

func Index(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			var result ServerResult
			result.Status = false
			body, _ := json.Marshal(result)
			fmt.Fprintf(response, string(body))
		}
	}()
	request.ParseForm()
	keywords := request.Form["keywords"]
	channel := make(chan []SpiderResult, len(keywords))
	for _, keyword := range keywords {
		log.Println("begin to handler keyword ", keyword)
		go HandlerUrl(keyword, channel)
	}
	var result ServerResult
	result.Status = true
	for i := 0; i < len(keywords); i++ {
		items := <-channel
		if len(items) == 0 {
			panic("download or extract err")
		} else {
			for _, item := range items {
				result.Data = append(result.Data, item)
			}
		}
	}

	body, err := json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(response, string(body))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("RunTimeError:", err)
		}
	}()

	log.Println("Start serving on port 8001")
	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
