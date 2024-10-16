package others

import "fmt"

func LearningFunctions() int {
	return 5
}

func Greet(name string) string {
	return "Hello, " + name + "!"
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func Run() {
	defer second()
	first()
}

func TestingPanicAndRecover() {
	// panic("Panic")
	str := recover()
	fmt.Println(str)
}
