package shapes

import (
	"testing"
)

func TestGetRectangleArea(t *testing.T) {
	width := 2.0
	height := 3.0
	result := GetRectangleArea(float32(width), float32(height))
	if result != 6 {
		t.Error("wrong logic")
	}
}

func TestGetCircleArea(t *testing.T) {
	radius := 3.0
	result := GetCircleArea(float32(radius))
	if result != 28.274334 {
		t.Error("wrong logic")
	}
}
