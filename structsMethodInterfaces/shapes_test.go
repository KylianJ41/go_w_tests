package structsmethodinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	hasArea := 40.0

	if got != hasArea {
		t.Errorf("got %.2f hasArea %.2f", got, hasArea)
	}
}

/*
func TestArea(t *testing.T) {

		checkArea := func(t testing.TB, shape Shape, hasArea float64) {
			t.Helper()
			got := shape.Area()
			if got != hasArea {
				t.Errorf("got %g hasArea %g", got, hasArea)
			}
		}

		t.Run("rectangles", func(t *testing.T) {
			rectangle := Rectangle{12.0, 6.0}
			checkArea(t, rectangle, 72.0)
		})
		t.Run("circles", func(t *testing.T) {
			circle := Circle{10}
			checkArea(t, circle, 314.1592653589793)
		})
	}
*/
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea) // the # format will show the class of the value, and therefore the type of the value!
			}
		})
	}
}
