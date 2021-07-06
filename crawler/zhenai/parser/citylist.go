package parser

import (
	"goAc/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(bytes []byte) engine.ParserResult {
	var match = regexp.MustCompile(cityListRe)
	cityList := match.FindAllStringSubmatch(string(bytes), -1)

	result := engine.ParserResult{}
	for _, list := range cityList {
		result.Items = append(result.Items, list[2])
		result.Requests = append(result.Requests, engine.Request{Url: list[1], ParserFunc: engine.NilParser})
	}
	return result
}
