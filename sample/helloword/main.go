package main

import (
	"fmt"
	"time"
)

func init()  {
	fmt.Println("初始。。。 你好世界！")
}

func main()  {
	fmt.Println("嗨，你好！")
	fmt.Println("当前时间：" + time.Now().String())
}


