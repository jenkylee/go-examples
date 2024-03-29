package simplefactory

// go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
// NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
// 在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。

import "fmt"

type API interface {
	Say(name string) string
}

// NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 2 {

		return &helloAPI{}
	} else {

		return &hiAPI{}
	}
}

// hiApi is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {

	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {

	return fmt.Sprintf("Hello, %s", name)
}
