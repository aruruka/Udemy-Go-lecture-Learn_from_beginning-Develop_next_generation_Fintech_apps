package three_definition

import "fmt"

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

const big = 9223372036854775807 + 1

func Constant() float64 {
	fmt.Println(Pi, Username, Password)
	return (big - 1)
}
