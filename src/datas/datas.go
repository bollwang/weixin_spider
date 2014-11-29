package datas

type Keyword struct {
	Keyword string
}

type Keywords struct {
	Keywords []Keyword
}

type SpiderResult struct {
	Url           string `json:"url"`
	PageType      int    `json:"page_type"`
	PageTime      int64  `json:"page_time"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	Author        string `json:"author"`
	ReplyNum      int    `json:"reply_num"`
	CrawlerSource string `json:"crawler_source"`
}

type ServerResult struct {
	Status bool           `json:"status"`
	Data   []SpiderResult `json:"data"`
}
