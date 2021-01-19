package main

/**
 * 选项设计模式
 */

import (
	"log"
	"time"
)

type who struct {
	name    string
	age     int
	timeout time.Duration
	key     int
	value   string
}

var defaultMyFunOptions  =  who{
	name: "cpj",
	age: 0,
	timeout: time.Second,
	value: "idcpj",
}

type whoOption func(options *who)

func WithOptionAge(age int) whoOption {
	return func(options *who) {
		options.age = age
	}
}

func WithOptionKeyAndValue(key int, value string) whoOption {
	return func(options *who) {
		options.key = key
		options.value = value
	}
}

func NewWho(name string, opts ...whoOption)  who {
	options := defaultMyFunOptions
	options.name = name
	for _, o := range opts {
		o(&options)
	}

	return options
}

func (w *who) print() {
	log.Println(w)
}

func main() {
	w := NewWho("idcpj",
		WithOptionAge(10),
		WithOptionKeyAndValue(12, "12"),
	)

	w.print()
}