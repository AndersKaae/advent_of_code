package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
)

func GetNCharacter(noChars int, idx int, letters []rune) []string {
	var listOfCharcters []string
	for i := 0; i < noChars; i++ {
		listOfCharcters = append(listOfCharcters, string(letters[idx+i]))
	}
	return listOfCharcters
}

func isDuplicates(listCharacter []string) bool {
	var ourChars []string
	for _, letter := range listCharacter {
		if utils.Contains(ourChars, letter) {
			return true
		}
		ourChars = append(ourChars, letter)
	}
	return false
}

func SolvePuzzle6Solver(noChars int) {
	input := utils.LoadFile("puzzle2022/puzzle6.txt")

	for _, row := range input {
		fmt.Println(row)
		letters := []rune(row)
		for idx := range letters {
			if idx+noChars > len(letters) { // This is to avoid attempting to get characters that would be out of bounds
				break
			}
			list4character := GetNCharacter(noChars, idx, letters)
			//fmt.Println(list4character)
			if !isDuplicates(list4character) {
				fmt.Println(idx + noChars)
				break
			}
		}
		fmt.Println()
	}
}

func SolvePuzzle6() {
	SolvePuzzle6Solver(4)
	SolvePuzzle6Solver(14)
}
