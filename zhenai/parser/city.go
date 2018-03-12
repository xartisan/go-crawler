package parser

import (
	"regexp"

	"github.com/xartisan/go-crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		//result.Items = append(result.Items, "User :"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(content []byte) engine.ParseResult {
				return ParseProfile(content, name, url)
			}})
	}
	return result
}
