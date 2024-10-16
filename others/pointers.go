package others

import "fmt"

// functions make copies of args and dont change the original values

func changeToZero(x int) {
	x = 0
}
func actuallyChangeToZero(y *int) {
	*y = 0
}

func Prog() {
	x := 5
	changeToZero(x)
	fmt.Println(x)
	//  x is still 5
	actuallyChangeToZero(&x)
	fmt.Println(x)

	// new can also be used to create a pointer
	// new is used to allocate space for a given data type
	z := new(int)

	fmt.Println(z)

}
