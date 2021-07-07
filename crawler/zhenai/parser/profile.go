package parser

import (
	"goAc/crawler/engine"
	"goAc/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

var infoRe = regexp.MustCompile(`<div data-v-499ba28c="" class="des f-cl">(.*)</div>`)
var nameRe = regexp.MustCompile(`<h1 data-v-499ba28c="" class="nickName">(.*)</h1>`)
func ParseProfile(content []byte) engine.ParserResult {
	match := infoRe.FindStringSubmatch(string(content))
	profile := model.Profile{}
	if match != nil {
		profile.Name = match[0]
	}

	infoMatch := nameRe.FindStringSubmatch(string(content))
	if infoMatch != nil {
		split := strings.Split(infoMatch[0], "|")
		if len(split) == 6 {
			profile.Occupation =split[0]
			profile.Age, _ = strconv.Atoi(split[1])
			profile.Gender = split[2]
			profile.Marriage = split[3]
			profile.Height, _ = strconv.Atoi(split[4])
			profile.Income = split[5]
		}
	}
	return engine.ParserResult{Items: []interface{}{profile}}
}