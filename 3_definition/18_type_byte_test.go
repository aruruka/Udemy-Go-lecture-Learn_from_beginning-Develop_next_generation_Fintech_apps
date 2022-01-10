package three_definition

import (
	"log"
	"testing"
)

func TestByteType(t *testing.T) {
	b := []byte{72, 73}
	log.Println(b)
	log.Println(string(b))

	c := []byte("HI")
	log.Println(c)
	log.Println(string(c))
}
