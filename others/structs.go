package others

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

var r1 Rectangle
var r2 = new(Rectangle)

func CalculateArea() {
	r10 := new(Rectangle)
	i10 := new(int)
	fmt.Println(r10, i10)
	pointerR1 := &Rectangle{3.4, 4.4}
	fmt.Print(pointerR1)
	r4 := new(Rectangle)
	r5 := Rectangle{length: 20, width: 30}
	fmt.Println(r4.length, r5.width)
}

func Area(r Rectangle) float64 {
	return r.length * r.width
}
