package main

import "testing"
import "time"

func Test(t *testing.T) {

	t1 := time.Sleep(time.Second * 1)
	t2 := time.After(time.Second * 4)

	if t2 < t1 {
		t.Error("Sleep is more than After")
	}
}
