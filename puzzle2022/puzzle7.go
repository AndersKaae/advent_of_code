package puzzle2022

import (
	"fmt"

	"github.com/AndersKaae/advent_of_code/utils"
)

type Node struct {
	Name     string
	IsDir    bool
	Size     int
	Children []*Node
}

func SolvePuzzle7() {

	input := utils.LoadFile("puzzle2022/puzzle7.txt")
	for _, row := range input {
		fmt.Println(row)

	}
}
