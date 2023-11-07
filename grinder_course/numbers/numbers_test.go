package main

import "testing"

func TestNewNumbers(t *testing.T) {
	n := newNumbers()

	if len(n) != 11 {
		t.Errorf("Expected length of 11, got %v", len(n))
	}
}

func TestEvenNumbers(t *testing.T) {
	n := newNumbers()

	if n.evenNumbers(0) != true {
		t.Errorf("Expected true, got %v", n.evenNumbers(0))
	}

	if n.evenNumbers(1) != false {
		t.Errorf("Expected false, got %v", n.evenNumbers(1))
	}
}

func TestOddNumbers(t *testing.T) {
	n := newNumbers()

	if n.oddNumbers(0) != false {
		t.Errorf("Expected false, got %v", n.oddNumbers(0))
	}

	if n.oddNumbers(1) != true {
		t.Errorf("Expected true, got %v", n.oddNumbers(1))
	}
}
