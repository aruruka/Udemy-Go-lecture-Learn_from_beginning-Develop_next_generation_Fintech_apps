package three_definition

import (
	"log"
	"testing"
)

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

const big = 9223372036854775807 + 1

func TestConstant(t *testing.T) {
	// t.Log(Pi, Username, Password)
	// t.Log(big - 1)
	log.Println(Pi, Username, Password)
	log.Println(big - 1)
}
