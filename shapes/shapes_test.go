package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %.2f, want %.2f", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rect := Rectangle{Width: 10.0, Height: 10.0}
	// 	want := 100.0
	// 	checkArea(t, rect, want)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{Radius: 10.0}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.want {
				t.Errorf("%#v got %g, wanted %g", areaTest.shape, got, areaTest.want)
			}
		})
	}
}
