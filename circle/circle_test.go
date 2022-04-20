package main 

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	area := CircleArea()
	want := 100 * math.Pi 
	if area != want {
		t.Fatalf("Want %v, but got %v", want, area)
	}
}