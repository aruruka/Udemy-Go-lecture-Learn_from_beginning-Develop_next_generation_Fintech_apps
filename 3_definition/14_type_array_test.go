package three_definition

import (
	"log"
	"testing"
)

func TestArrayType(t *testing.T) {
	var a [2]int
	a[0] = 100
	a[1] = 200
	log.Println(a)

	var b [2]int = [2]int{100, 200}
	log.Printf("type: %T; value: %v", b, b)
	// append(b, 300) // cannot compile, array is immutable.

	var c []int = []int{100, 200}
	c = append(c, 300)
	log.Printf("type: %T; value: %v", c, c)
}
