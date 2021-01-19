package main

/**
 * 组合模式
 * 一个具有层级关系的对象由一系列拥有父子关系的对象通过树形结构组成(类似树形式的组件)
 */

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
)

// IComponent 组件接口
type IComponent interface {
	// 添加一个子组件
	Mount(c IComponent, components ...IComponent) error
	// 移除一个组件
	Remove(c IComponent) error
	// 执行组件&子组件
	Do(ctx context.Context) error
}

// BaseComponent 基础组件
// 实现Add：添加一个子组件
// 实现Remove：删除一个子组件
type BaseComponent struct {
	// 子组件列表
	ChildComponents []IComponent
}

// Mount 挂载一个子组件
func (bc *BaseComponent) Mount(c IComponent, components ...IComponent) (err error){
	bc.ChildComponents = append(bc.ChildComponents, c)
	if len(components) == 0 {
		return
	}
	bc.ChildComponents = append(bc.ChildComponents, components...)

	return
}

// Remove 移除一个子组件
func (bc *BaseComponent) Remove(c IComponent) (err error) {
	if len(bc.ChildComponents) == 0 {
		return
	}

	for k, childComponent := range bc.ChildComponents {
		if c == childComponent {
			fmt.Println(runFuncName(), "移除：", reflect.TypeOf(childComponent))
			bc.ChildComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}

	return
}

// Do 执行组件&子组件
func (bc *BaseComponent) Do(ctx context.Context) (err error) {
	// Do nothing
	return
}

// ChildsDo 执行子组件
func (bc *BaseComponent) ChildsDo(ctx context.Context) (err error) {
	// 执行子组件
	for _, childComponent := range bc.ChildComponents {
		if err = childComponent.Do(ctx); err != nil {
			return err
		}
	}

	return
}

type All struct {
	BaseComponent
}

type A1 struct {
	BaseComponent
}

func (a *A1) Do(ctx context.Context) (err error) {
	fmt.Println(runFuncName(), "A1...")
	a.ChildsDo(ctx)

	return
}

type A11 struct {
	BaseComponent
}

func (a *A11) Do(ctx context.Context) (err error) {
	fmt.Println(runFuncName(), "A11...")
	a.ChildsDo(ctx)

	return
}

type A12 struct {
	BaseComponent
}

func (a *A12) Do(ctx context.Context) (err error) {
	fmt.Println(runFuncName(), "A12...")
	a.ChildsDo(ctx)

	return
}

type B1 struct {
	BaseComponent
}

func (a *B1) Do(ctx context.Context) (err error) {
	fmt.Println(runFuncName(), "B1...")
	a.ChildsDo(ctx)

	return
}

type C1 struct {
	BaseComponent
}

func (a *C1) Do(ctx context.Context) (err error) {
	fmt.Println(runFuncName(), "C1...")
	a.ChildsDo(ctx)

	fmt.Printf("%+v\n", ctx.Value("a"))

	return
}

func main()  {
	//
	all := &All{}

	a1 := &A1{}
	a1.Mount(&A11{}, &A12{})

	b1 := &B1{}
	all.Mount(a1, b1, &C1{})

	all.Remove(b1)

	ctx := context.WithValue(context.Background(), "a", "b")
	all.ChildsDo(ctx)
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	return f.Name()
}