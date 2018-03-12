package parser

import (
	"regexp"

	"github.com/xartisan/go-crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	//limit := 10
	for _, m := range matches {
		//result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity})
		//if limit--; limit <= 0 {
		//	break
		//}
	}
	return result
}
