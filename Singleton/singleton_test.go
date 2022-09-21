package Singleton

import (
	"testing"
)

func TestGetIns(t *testing.T) {
	res := GetInstance("test1")
	res2 := GetInstance("test2")
	if res.name != res2.name {
		t.Fatal()
	}
}
