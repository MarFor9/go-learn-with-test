package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}

	got := Perimeter(rectangle)
	expectedArea := 40.0

	if got != expectedArea {
		t.Errorf("expectedArea: %.2f but got: %.2f", expectedArea, got)
	}
}

func TestArea(t *testing.T) {

	type testConfig struct {
		name         string
		shape        Shape
		expectedArea float64
	}

	for _, tc := range []testConfig{
		{
			name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, expectedArea: 100.0,
		},
		{
			name: "Circles", shape: Circles{Radius: 10}, expectedArea: 314.1592653589793,
		},
		{
			name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.expectedArea {
				t.Errorf("%#v got %g expectedArea %g", tc, got, tc.expectedArea)
			}
		})
	}
}
