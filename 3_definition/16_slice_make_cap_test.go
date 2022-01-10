package three_definition

import (
	"log"
	"testing"
)

func TestSlickMakeCap(t *testing.T) {
	n := make([]int, 3, 5)
	log.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
	n = append(n, 0, 0)
	log.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5)
	log.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)

	a := make([]int, 3)
	log.Printf("len=%d cap=%d value=%v", len(a), cap(a), a)

	b := make([]int, 0)
	var c []int
	log.Printf("len=%d cap=%d value=%v", len(b), cap(b), b)
	log.Printf("len=%d cap=%d value=%v", len(c), cap(c), c)

	// c = make([]int, 5)
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		log.Println(c)
	}
	log.Println(c)
}
