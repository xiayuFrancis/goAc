package main

import (
	"goAc/crawler/engine"
	"goAc/crawler/zhenai/parser"
)

func main() {
	//engine.Run(engine.Request{Url: "https://www.zhenai.com/zhenghun", ParserFunc: parser.ParserCityList})
	engine.Run(engine.Request{Url: "https://album.zhenai.com/u/1870666262", ParserFunc: parser.ParseProfile})

}
