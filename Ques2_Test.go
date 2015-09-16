package main

import "testing"

func TestQuesTwo(t *testing.T) {
	var val float64

	c := Circle{20}
	r := Rectangle{30, 50}
	//val1 = c.perimeter()
	//val2 = r.perimeter()
	val = totalPerimeter(&c, &r)

	if val != 285.66370614359175 {
		t.Error("Expected 285.66370614359175,got ", val)
	}
}
