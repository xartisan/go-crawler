package scheduler

import (
	"github.com/xartisan/go-crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//log.Printf("submit request %v\n", request)
	go func() { s.workerChan <- request }()
}
