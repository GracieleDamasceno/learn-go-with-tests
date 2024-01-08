package structsmethodsinterfaces

import "testing"

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

	for _, testElement := range areaTests {
		t.Run(testElement.name, func(t *testing.T) {
			got := testElement.shape.Area()
			if got != testElement.hasArea {
				t.Errorf("%#v got %g want %g", testElement.shape, got, testElement.hasArea)
			}
		})

	}

}
