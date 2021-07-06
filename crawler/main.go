package main

import (
	"goAc/crawler/engine"
	"goAc/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{Url: "https://www.zhenai.com/zhenghun", ParserFunc: parser.ParserCityList})
}
