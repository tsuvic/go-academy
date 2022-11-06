package main

import "testing"

func Test1(t *testing.T) {
	var num []int = []int{1, 2, 3, 4, 5}
	if !(Sum(num) == 15) {
		t.Error(`miss`)
	}
}

func Test2(t *testing.T) {
	var num []int = []int{1, 3, 5}
	if !(Sum(num) == 9) {
		t.Error(`miss`)
	}
}
