package three_definition

import (
	"log"
	"testing"
)

func TestSliceType(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 6}
	log.Println(n)
	log.Println(n[2])
	log.Println(n[2:4])
	log.Println(n[:2])
	log.Println(n[2:])
	log.Println(n[:])

	n[2] = 100
	log.Println(n)

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	log.Println(board)
	board = append(board, []int{100})
	log.Println(board)
}
