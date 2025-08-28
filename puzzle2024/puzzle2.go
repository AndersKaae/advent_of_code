package puzzle2024

import (
	"fmt"
	"strings"

	"github.com/AndersKaae/advent_of_code/utils"
)

type Row struct {
	content []int
	safe    []bool
}

func FormatPuzzleInput(puzzleInput []string) []Row {
	structList := []Row{}
	for _, row := range puzzleInput {
		splitRow := strings.Split(row, " ")
		intRowList := []int{}
		for _, numberStr := range splitRow {
			numberInt := utils.ConvertStringToInt(numberStr)
			intRowList = append(intRowList, numberInt)
		}
		row := Row{content: intRowList, safe: false}
		structList = append(structList, row)
	}
	return structList
}

func SolvePuzzle2() {
	puzzleIntput := utils.LoadFile("puzzle2024/puzzletext/puzzle2sample.txt")
	structList := FormatPuzzleInput(puzzleIntput)
	for _, row := range structList {
		fmt.Println(row)
	}

}
