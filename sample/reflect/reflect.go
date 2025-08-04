package main

import (
	"fmt"
	"reflect"
)

// 反射获取interface类型信息
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)

	// kind() 可以获取具体类型
	k := t.Kind()
	fmt.Println(k)

	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

// 反射修改值
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:

		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem() 获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}
func main() {
	var x float64 = 3.4
	reflect_type(x)
	reflect_set_value(&x)
	fmt.Println("main:", x)
}
