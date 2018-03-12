package main

import (
	"github.com/xartisan/go-crawler/engine"
	"github.com/xartisan/go-crawler/persist"
	"github.com/xartisan/go-crawler/scheduler"
	"github.com/xartisan/go-crawler/zhenai/parser"
)

func main() {
	saveWorker, err := persist.SaveWorker("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 100,
		ItemChan:  saveWorker,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
