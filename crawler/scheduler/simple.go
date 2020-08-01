package scheduler

import (
	"github.com/kentliuqiao/learngo/crawler/engine"
)

// SimpleScheduler 通过一个管道向worker发送Request
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Submit ...
func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- r }()
}

// WorkerChan return chan of request
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// Run ...
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// WorkerReady ...
func (s *SimpleScheduler) WorkerReady(chan engine.Request) {}
