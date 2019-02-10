package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld("appleboy")
	want := "Hi, appleboy"
	if got != want {
		t.Errorf("Testing fail")
	}
	got2 := HelloWorld("appleboy ")
	if got2 != want {
		t.Errorf("Testing fail")
	}
}
