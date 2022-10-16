package Singleton

import "sync"

type singleton struct {
	name string
}

var instance *singleton
var once sync.Once

func GetInstance(name string) *singleton {
	once.Do(func() {
		instance = &singleton{name: name}
	})
	return instance
}
