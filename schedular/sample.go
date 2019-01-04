package schedular

import "robin/learngo/crawler/engine"

type SimpleSchedular struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedular) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleSchedular) WorkerReady(chan engine.Request) {
}

func (s *SimpleSchedular) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleSchedular) Submit(r engine.Request) {
	go func() {s.workerChan <- r}()
}


