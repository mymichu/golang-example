package main

import ( 
	"testing"
)

func TestAverage(t *testing.T) {
	if Average(10,5) != 7.5 {
		t.Error("expected 7.5")
	}
}