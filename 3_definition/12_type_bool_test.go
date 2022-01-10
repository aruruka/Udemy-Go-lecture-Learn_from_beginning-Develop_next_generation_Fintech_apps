package three_definition

import (
	"log"
	"testing"
)

func TestBoolType(t *testing.T) {
	xt, f := true, false
	log.Printf("%T %v %t\n", xt, 1, xt)
	log.Printf("%T %v %t\n", f, 0, f)

	// log.Println(true && true)
	log.Println(true && false)
	// log.Println(false && false)

	// log.Println(true || true)
	log.Println(true || false)
	// log.Println(false || false)

	log.Println(!true)
	log.Println(!false)
}
