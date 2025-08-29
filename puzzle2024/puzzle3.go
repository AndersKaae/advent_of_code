package puzzle2024

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strings"
)

type Puzzle struct {
	Do      bool
	Content string
}

func GetPuzzleString(file string) string {
	puzzleIntput := utils.LoadFile(file)
	return strings.Join(puzzleIntput, "\n")
}

func SolvePuzzle3partA(puzzleInput string) int {
	numberSequence := []string{}

	// Split by mul
	splitPuzzleInput := strings.Split(puzzleInput, "mul(")
	for _, split := range splitPuzzleInput {
		fmt.Println(split)
		splitRunes := []rune(split)
		for idx, glyph := range splitRunes {
			acceptableStrings := []string{",", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", ")"}
			// save everything prior to ")"
			if !utils.Contains(acceptableStrings, string(glyph)) {
				break
			}
			if string(glyph) == ")" {
				numberSequence = append(numberSequence, string(splitRunes[:idx]))
				break
			}

		}
	}
	sum := 0
	for _, pair := range numberSequence {
		fmt.Println(pair)
		splitPair := strings.Split(pair, ",")
		firstInt, err := utils.ConvertStringToInt(splitPair[0])
		if err != nil {
		}
		secondInt, err := utils.ConvertStringToInt(splitPair[1])
		if err != nil {
		}
		sum += firstInt * secondInt
	}
	return sum
}

func SolvePuzzle3partB(puzzleInput string) []Puzzle {
	fmt.Println(puzzleInput)
	splitPuzzleInput := strings.Split(puzzleInput, "don't()")
	puzzleStructList := []Puzzle{}
	for idx, dontSubstring := range splitPuzzleInput {
		if idx == 0 {
			puzzle := Puzzle{Do: true, Content: dontSubstring}
			puzzleStructList = append(puzzleStructList, puzzle)
			continue
		}
		doSubstring := strings.Split(dontSubstring, "do()")
		for idx, string := range doSubstring {
			if idx == 0 {
				puzzle := Puzzle{Do: false, Content: string}
				puzzleStructList = append(puzzleStructList, puzzle)
				continue
			}
			puzzle := Puzzle{Do: true, Content: string}
			puzzleStructList = append(puzzleStructList, puzzle)
		}
	}

	for _, puzzleStruct := range puzzleStructList {
		fmt.Println(puzzleStruct)
	}
	return puzzleStructList

}

func SolvePuzzle3() {
	file := "puzzle2024/puzzletext/puzzle3.txt"
	puzzleIntput := GetPuzzleString(file)
	sum := SolvePuzzle3partA(puzzleIntput)
	fmt.Println("Puzzle 1 result")
	fmt.Println(sum)
	fmt.Println("----------")
	sum = 0
	puzzleStructList := SolvePuzzle3partB(puzzleIntput)
	for _, puzzle := range puzzleStructList {
		if puzzle.Do == true {
			sum += SolvePuzzle3partA(puzzle.Content)
		}
	}
	fmt.Println("Puzzle 2 result")
	fmt.Println(sum)
}
