package concurrency

import (
	"fmt"
	"sync"
)

type (
	service struct {
		sync.Mutex
		stopCh    chan struct{}
		isStarted bool
	}
)

func (s *service) start() {
	if s.isStarted {
		fmt.Println("Service already Started!")
		return
	}
	s.stopCh = make(chan struct{})
	go func() {
		s.Lock()
		defer s.Unlock()
		s.isStarted = true
		<-s.stopCh
	}()
}
func (s *service) stop() {
	s.Lock()
	defer s.Unlock()
	s.isStarted = false
	close(s.stopCh)
}
