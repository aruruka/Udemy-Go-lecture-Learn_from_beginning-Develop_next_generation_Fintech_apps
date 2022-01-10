package three_definition

import (
	"log"
	"testing"
)

func TestNumberType(t *testing.T) {
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)

	log.Println(u8, i8, f32, c64)
	log.Printf("type=%T value=%v\n", u8, u8)

	x := 1 + 1
	log.Println(x)

	log.Println(1+1, 2+2)
	log.Println("1 + 1 =", 1+1)
	log.Println("10 - 1 =", 10-1)
	log.Println("10 / 2 =", 10/2)
	log.Println("10 / 3 = ", 10/3)
	log.Println("10.0 / 3 =", 10.0/3)
	log.Println("10 / 3.0 =", 10/3.0)
	log.Println("10 % 2 =", 10%2)
	log.Println("10 % 3 =", 10%3)

	x = 0
	log.Println(x)
	x++
	log.Println(x)
	x--
	log.Println(x)

	log.Println(1 << 0)
	log.Println(1 << 1)
	log.Println(1 << 2)
	log.Println(1 << 3)
}
