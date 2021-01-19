package main

/**
 * 装饰设计模式(添加钩子)
 */

import "fmt"

// 定义抽象的组件
type Component interface {
	Operate()
}

type Component1 struct {

}

func (c Component1) Operate() {
	fmt.Printf("%+v\n", "c1 operate")
}

// 定义抽象的装饰者
type Decorator interface {
	Component
	Do() // 装饰行为，可以有多个
}

// 定义一个具体的装饰器
type Decorator1 struct {
	c Component
}

func (d Decorator1) Do() {
	fmt.Printf("%+v\n", "c1 发生了装饰行为")
}

// 重新实现component 的方法
func (d Decorator1) Operate() {
	d.Do()
	d.c.Operate()
	d.Do()
}

func main() {
	// 无装饰模式
	c1 := &Component1{}
	c1.Operate()

	// 装饰模式
	c2 := &Decorator1{
		c: c1,
	}

	c2.Operate()
}