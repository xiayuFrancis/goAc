package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"time"
)

func main() {
	urlstr := "http://metalsucks.net"
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}

	c := colly.NewCollector()
	c.SetRequestTimeout(100 * time.Second)
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", u.Host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", u.Host)
		r.Headers.Set("Referer", urlstr)
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
	})

	c.OnResponse(func(resp *colly.Response) {
		// 读取url内容 colly读取的内容传入给goquery
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body))
		if err != nil {
			log.Fatal(err)
		}
		// 找到抓取项
		htmlDoc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
			band := s.Find("a").Text()
			title := s.Find("i").Text()
			fmt.Printf("Review %d: %s - %s\n", i, band, title)
		})
	})

	c.OnError(func(resp *colly.Response, errHttp  error) {
		err = errHttp
	})

	err = c.Visit(urlstr)
}
