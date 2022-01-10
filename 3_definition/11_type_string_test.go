package three_definition

import (
	"log"
	"strings"
	"testing"
)

func TestStringType(t *testing.T) {
	log.Println("Hello World")
	log.Println("Hello " + "World")
	log.Println(string("Hello World"[0]))
	var s string = "Hello World"

	s = strings.Replace(s, "H", "X", 1)
	log.Println(s)
	log.Println(`Test
						Test
Test`)

	log.Println("\"")
	log.Println(`"`)
}
