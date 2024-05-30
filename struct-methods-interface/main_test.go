package struct_methods_interface

import "testing"

func Perimeter(width float64, heigth float64) float64 {
	return 2 * (width + heigth)
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func Area(width, height float64) float64 {
	return width * height
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
