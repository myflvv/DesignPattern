package factory

import "testing"

func TestNewBook(t *testing.T) {
	n := NewBook("chinese")
	r := n.Write()
	if r != "write chinese book" {
		t.Fail()
	}
}
