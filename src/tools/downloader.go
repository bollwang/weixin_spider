package tools

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
	//iconv "github.com/djimenez/iconv-go"
)

func Download(pageurl string) (rep string, err error) {
	log.Println("begin to download: ", pageurl)
	defer func() {
		if err := recover(); err != nil {
			log.Println("recovered in download", err)
			switch x := err.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			rep = ""
		}
	}()

    // add your proxy here
	proxy, _ := url.Parse("")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(25 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*20)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	request, _ := http.NewRequest("GET", pageurl, nil)
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36")

	response, err := client.Do(request)
	defer response.Body.Close()
	if err == nil {
		if response.StatusCode == 200 {
			//body, _ := ioutil.ReadAll(response.Body)
			//html := make([]byte, len(body))
			//html = html[:]
			//iconv.Convert(body, html, "gbk", "utf-8")
			html, _ := ioutil.ReadAll(response.Body)
			ioutil.WriteFile("sogou.html", html, 0644)
			return string(html), nil
		} else {
			return "", errors.New("download status=" + string(response.StatusCode))
		}
	} else {
		return "", errors.New("download error")
	}
}

func DownloadRetry(pageurl string) (rep string, err error) {
	for i := 0; i < 5; i++ {
		html, err := Download(pageurl)
		if err == nil && len(html) > 5000 {
			return html, err
		} else {
			log.Println("download failed, begin to retry ", pageurl)
		}
	}
	return "", err
}
