package main

/**
 *  sync.WaitGroup 控制结果题退出
 */

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type server struct {
	quit chan struct{}
	wait sync.WaitGroup
	isRun bool
}

func newServer() *server {
	return &server{
		quit: make(chan struct{}),
		wait:sync.WaitGroup{},
		isRun:true,
	}
}

func (s *server) close() {
	close(s.quit)
	s.wait.Wait()
}

func (s *server) task() {
	s.wait.Add(1)
	defer s.wait.Done()

	for {
		select {
		case <-s.quit:
			// 退出协程
			fmt.Printf("%+v\n", "stop")
			return
		default:

		}

		fmt.Printf("%+v\n", "task")
		go s.subTask()
		time.Sleep(2*time.Second)
	}
}

func (s *server) subTask() {
	s.wait.Add(1)
	defer s.wait.Done()

	select {
	case <-s.quit:
		// 退出协程
		fmt.Printf("%+v\n", "sub-stop")
		return
	default:

	}

	fmt.Printf("%+v\n", "sub-task")
	time.Sleep(2*time.Second)
}

func main()  {
	s := newServer()
	defer s.close()
	go s.task()

	sig := make(chan  os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
}