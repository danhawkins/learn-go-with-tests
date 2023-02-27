package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}

		checkArea(t, rect, 72.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		checkArea(t, circle, 314.1592653589793)
	})
}
