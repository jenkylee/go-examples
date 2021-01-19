package main

import (
	"log"
)

// 实现一个上下文的类
type Context struct {
	Strategy
}

// 抽象的策略
type Strategy interface {
	Do()
}

type Strategy1 struct {

}

func (s *Strategy1) Do() {
	log.Printf("%v", "实现策略1")
}

type Strategy2 struct {

}

func (s *Strategy2) Do() {
	log.Printf("%v", "实现策略2")
}

func main()  {
	ctx := Context{}

	ctx.Strategy = &Strategy1{}
	ctx.Do()

	strategy2 := &Strategy2{}
	ctx.Strategy = strategy2
	ctx.Do()
}