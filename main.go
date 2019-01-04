package main

import (
	"robin/learngo/crawler/engine"
	"robin/learngo/crawler/zhenai/parser"
	"robin/learngo/crawler/schedular"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &schedular.QueuedScheduler{},
		WorkerCount:100,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}