package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm() // 解析参数， 默认是不会解析的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello jenkylee!") // 写入到w的是输出到客户端
}

func main()  {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)

	if  err != nil{
		log.Fatal("Listen And Server:" , err)
	}
}


