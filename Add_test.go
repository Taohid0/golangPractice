package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2,3)
	if result!=5 {
		t.Errorf("expted 5 got", result)
	}
}
