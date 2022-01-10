package three_definition

import (
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"apple": 100, "banana": 200}
	log.Println(m)
	log.Println(m["apple"])
	m["banana"] = 300
	log.Println(m)
	m["new"] = 500
	log.Println(m)

	log.Println(m["nothing"])

	v, ok := m["apple"]
	log.Println(v, ok)

	v, ok2 := m["nothing"]
	log.Println(v, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	log.Println(m2)

	/* var m3 map[string]int
	m3["pc"] = 5000
	log.Println(m3) */

	var m3 map[string]int
	if m3 == nil {
		log.Println("Nil")
	}

	var s []int
	if s == nil {
		log.Println("Nil")
	}

}
