package singleton

import "fmt"

type Singleton2 interface {
	Speak()
}

type singleton2 struct {
}

func (singleton2) Speak() {
	fmt.Println("s")
}

var singletonInstance Singleton2

func NewSingleton2Instance() Singleton2 {
	if nil == singletonInstance {
		singletonInstance = &singleton2{}
	}
	return singletonInstance
}
