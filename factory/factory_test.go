package factory

import "testing"

func TestGetFBook(t *testing.T) {
	r, _ := GetFBook("chinese")
	if r.Write() != "write chinese" {
		t.Fail()
	}
	e, _ := GetFBook("english")
	if e.Write() != "write english" {
		t.Fail()
	}
}
