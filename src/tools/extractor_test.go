package tools

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestExtractUrl(t *testing.T) {
	text, _ := ioutil.ReadFile("sogou.html")
	rslt, _ := ExtractUrl(string(text))
	fmt.Println(len(rslt))
	for _, r := range rslt {
		fmt.Println("result", r)
	}
}

/*
func TestExtractPage(t *testing.T) {
	text, _ := ioutil.ReadFile("weixin.html")
	rslt, _ := ExtractPage(string(text))
    fmt.Println(rslt)
}
*/
