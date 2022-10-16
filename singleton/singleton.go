package singleton

import "sync"

type singleton struct {
	name string
}

var instanse *singleton
var once sync.Once

func GetInstanse(name string) *singleton {
	once.Do(func() {
		instanse = &singleton{name: name}
	})
	return instanse
}
