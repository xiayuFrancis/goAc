package main

import (
	"fmt"
	"goAc/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}
func post(poster Poster) {
	poster.Post("https://www.baidu.com", map[string]string{
		"name":   "xy",
		"course": "go",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(poster RetrieverPoster) {
	fmt.Println(poster.Get("http://www.baidu.com"))
	poster.Post("https://www.baidu.com", map[string]string{
		"name":   "xy",
		"course": "go",
	})
}

func main() {
	var r RetrieverPoster
	r = &real.Retriever{}
	fmt.Printf("%T, %v", r, r)
	retriever, ok := r.(*real.Retriever)
	if ok {
		fmt.Println(retriever)
		session(r)
	}
	//fmt.Println(download(r))
}
