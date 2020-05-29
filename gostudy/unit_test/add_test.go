package main

import "testing"

func TestAdd(t *testing.T) {
	var a = 10
	var b = 20
	t.Logf("a = %d b = %d\n", a, b)
	c := Add(a, b)
	if c != 30 {
		t.Fatalf("invalid a + b ")
	}
	t.Logf("a = %d b = %d sun:=%d\n", a, b, c)

}
