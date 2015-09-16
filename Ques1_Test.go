package main

import "testing"

func TestQ(t *testing.T) {
	//var val int
	//val = 20
	p := v(20).(int)

	if p != 6765 {

		//var x interface{} = 20
		//i := x.(int)

		//val = x
		//if v != 6765 {
		t.Error("Expected 6765,got ", p)
	}
}
