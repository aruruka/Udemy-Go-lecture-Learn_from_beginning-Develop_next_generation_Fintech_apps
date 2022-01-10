package three_definition

import (
	"fmt"
	"testing"
)

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func TestVariadicParameter(t *testing.T) {
	foo()
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)
}
