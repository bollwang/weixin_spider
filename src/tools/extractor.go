package tools

import (
	. "datas"
	"github.com/opesun/goquery"
	"regexp"
	"strings"
	"time"
)

func ExtractUrl(page string) (ret []string, err error) {
	reg, _ := regexp.Compile(`http://mp.weixin.qq.com/s\?(.*)\#rd`)
	urls := reg.FindAllString(page, -1)

	result := make([]string, 0)
	found := make(map[string]bool)
	for _, val := range urls {
		if _, ok := found[val]; !ok {
			found[val] = true
			result = append(result, val)
		}
	}
	return result, nil
}

func ExtractPage(page string) (ret SpiderResult, err error) {
	query, _ := goquery.ParseString(page)

	// extract title
	title := query.Find("#activity-name").Text()
	title = strings.TrimSpace(title)

	// extract time
	ptime := query.Find("#post-date").Text()
	ptime = strings.TrimSpace(ptime)

	// extract author
	author := query.Find("#post-user").Text()
	author = strings.TrimSpace(author)

	// extract content
	content := query.Find("#js_content").Text()
	content = strings.TrimSpace(content)

	utime := time.Now().Unix()
	the_time, err := time.Parse("2006-01-02", ptime)
	if err == nil {
		utime = the_time.Unix()
	}
	ret = SpiderResult{
		PageTime: utime,
		Title:    title,
		Content:  content,
		Author:   author,
		ReplyNum: 0,
	}
	return ret, nil
}
