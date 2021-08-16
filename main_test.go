package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did NOT get an error when we should have")
	}
}

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool 
}{
	{"valid data", 100.0, 10.0, 10.0, false},
	{"valid data", 100.0, 0.0, 0.0, true},
}

func TestDiv(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend,tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but id not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an err but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got )
		}

	}
}