package singleton

import (
	"testing"
)

func TestGetInstanse(t *testing.T) {
	test1 := GetInstanse("test1")
	test2 := GetInstanse("test2")
	if test1.name == test2.name {
		t.Log("ok")
	} else {
		t.Fail()
	}
}
