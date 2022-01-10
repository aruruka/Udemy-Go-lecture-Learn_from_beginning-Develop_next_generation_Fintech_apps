package three_definition

import (
	"log"
	"strconv"
	"testing"
)

func TestTypeConversion(t *testing.T) {
	var x int = 1
	xx := float64(x)
	log.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	log.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	log.Printf("%T %v\n", i, i)

	var s2 string = "Hello"
	i2, _ := strconv.Atoi(s2)
	log.Printf("%T %v\n", i2, i2)

	h := "Hello World"
	log.Println(string(h[0]))
}
