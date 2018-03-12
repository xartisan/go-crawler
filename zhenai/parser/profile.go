package parser

import (
	"regexp"
	"strconv"

	"github.com/xartisan/go-crawler/engine"
	"github.com/xartisan/go-crawler/model"
)

var (
	ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
	//heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
	heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">(\d+)CM</span></td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<])+元</td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	xingzuoRe    = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
	hukouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	idUrlRe      = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
)

func ParseProfile(content []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	if age, err := strconv.Atoi(extractContent(content, ageRe)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractContent(content, heightRe)); err == nil {
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractContent(content, weightRe)); err == nil {
		profile.Weight = weight
	}
	profile.Income = extractContent(content, incomeRe)
	profile.Marriage = extractContent(content, marriageRe)
	profile.Education = extractContent(content, educationRe)
	profile.Occupation = extractContent(content, occupationRe)
	profile.Xingzuo = extractContent(content, xingzuoRe)
	profile.Hukou = extractContent(content, hukouRe)
	profile.Gender = extractContent(content, genderRe)
	profile.Car = extractContent(content, carRe)
	profile.House = extractContent(content, houseRe)
	parseResult := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractContent([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	return parseResult
}

func extractContent(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return "--"
}
