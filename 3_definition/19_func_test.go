package three_definition

import (
	"fmt"
	"log"
	"testing"
)

/* func add(x int, y int) {
	log.Println("add function")
	log.Println(x + y)
} */

/* func add(x int, y int) int {
	return x + y
} */

/* func add(x int, y int) (int, int) {
	return x + y, x - y
} */

func cal(price, item int) (result int) {
	result = price * item // 📓Note: this is not a declaration, because `result` is declared as the returned variable.
	return
}

func convert(price int) float64 {
	return float64(price)
}

func TestFunc(t *testing.T) {
	// add(10, 20)

	/* r := add(100, 200)
	log.Println(r) */

	/* r1, r2 := add(100, 200)
	log.Println(r1, r2) */

	r3 := cal(100, 2)
	log.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

}
