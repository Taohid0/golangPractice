package main

import "testing"

func TestHelloGo(t *testing.T) {
	expected := "welcome"
	actual := HelloGo()
	if actual!=expected{
		t.Error("Fail")
	}
}
