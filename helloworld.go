package main

import (
	"fmt"
	"os"
	"using-book/others"
)

const boilingPointF = 212.0

func main() {
	// // learningVariables()
	// // takingUserInputs()
	// // forLoop()
	// // switchCase()
	// // arraysAndSlices()
	// // maps()
	// // others.
	// test2 := others.Greet("me")
	// // fmt.Println(Mandem)
	// // test := strings.Split(SayTest(), "")
	// fmt.Println(test2)
	// others.TestingPanicAndRecover()
	// others.CalculateArea()
	others.ExecuteConc()
}
func maps() {
	// var x map[string]int
	x := make(map[string]int)
	x["key"] = 20

	fmt.Println(x["key"])

	elements := make(map[string]string)

	elements["H"] = "hydrogen"
	elements["Hi"] = "Helium"
	elements["O"] = "oxygen"

	if name, ok := elements["test"]; ok {
		fmt.Println(name, ok)
	}

	students := map[string]int{
		"Frimps": 3031020,
		"paps":   239943,
	}

	students2 := map[string]map[string]string{
		"frimps": {
			"fullName": "Akwasi Ampomah",
		},
	}

	fmt.Println(students)
	fmt.Println(students2["frimps"]["fullName"])

}
func arraysAndSlices() {
	var numbers [10]int
	var names [10]string

	newNames := [4]string{"frimps", "ama", "akosuah", "test"}
	fmt.Println(numbers)
	fmt.Println(names)
	fmt.Println(newNames)

	for _, value := range numbers {
		fmt.Println(value)
	}

	// slices
	var x []float32

	y := make([]int, 3)

	xy := make([]int, 5, 10)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(cap(x))
	fmt.Println(cap(y))
	fmt.Println(len(xy), cap(xy))

	colors := [5]string{"red", "blue", "white", "black", "orange"}

	// does not include end index but thats where it ends
	// array[low(included):high(not included)]
	// array[0:], array [:]
	colorsSlice := colors[1:3]
	fmt.Println(colorsSlice)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)
}

func forLoop() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}
func switchCase() {
	i := 1
	switch i {
	case 1:
		{
			fmt.Print("It is one")
		}
	case 2:
		{
			fmt.Print("It is one")
		}
	case 3:
		{
			fmt.Print("It is one")
		}
	case 4:
		{
			fmt.Print("It is one")
		}
	default:
		{
			fmt.Println("None of the above")
		}
	}
}
func ifElseStatement() {
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
func takingUserInputs() {
	var userInput int
	fmt.Println("Enter a your year of birth: ")
	fmt.Scanf("%d", &userInput)
	output := 2024 - userInput
	fmt.Println(output)
}
func learningVariables() {
	var x string = "frimps"
	var name string
	name = "opoku"

	m := "test"
	fmt.Println(x)
	fmt.Println(name)
	fmt.Println(m)
}

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
func learningString() {
	fmt.Println(len(("Hello, world")))
	fmt.Println("hello. world"[1])
}

func otherProg() {
	var f = boilingPointF
	var c = (f - 32) * 5 / 9
	var myName string
	fmt.Printf(myName)
	var file, fileOpeningError = os.Open("samplex.txt")

	if fileOpeningError == nil {
		fmt.Print(file)
	}

	fmt.Println(fileOpeningError)

	fmt.Printf("The boiling point of water = %gF or %gC ", f, c)

	i, j := 0, 3

	fmt.Print(i, j)
	out, err := os.Create("test.txt")

	if err != nil {
		fmt.Print(out)
	}

	// pointer()
}
