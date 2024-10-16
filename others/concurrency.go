package others

import (
	"fmt"
	// "math/rand"
	// "time"
)

// go routines

func f(n int) {
	for i := 0; i < 9; i++ {
		fmt.Println(n, " : ", i)
		// amt := time.Duration(rand.Intn(250))
		// time.Sleep(amt * time.Millisecond)
	}
}

func ExecuteConc() {
	// for i := 0; i < 10; i++ {
	go f(0)
	// f(0)
	// }
	// var input string
	// fmt.Scanln(&input)
}
