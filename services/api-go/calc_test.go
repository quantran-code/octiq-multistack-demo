package api

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Fatalf("add failed")
	}
}

func TestDivide(t *testing.T) {
	if _, ok := Divide(10, 2); !ok {
		t.Fatalf("divide failed")
	}
	if _, ok := Divide(1, 0); ok {
		t.Fatalf("expected divide by zero to fail")
	}
}
