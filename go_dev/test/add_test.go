package test

import "testing"

func TestAdd(t *testing.T) {
	sum := add(1, 2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}

func TestAdd1(t *testing.T) {
	t.Errorf("the result is error")
}
