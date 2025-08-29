package puzzle2024

import (
	"fmt"
	"strings"

	"github.com/AndersKaae/advent_of_code/utils"
)

type Row struct {
	content       []int
	directionUp   []bool
	directionDown []bool
}

func FormatPuzzleInput(puzzleInput []string) []Row {
	structList := []Row{}
	for _, row := range puzzleInput {
		splitRow := strings.Split(row, " ")
		newRow := Row{}
		for _, numberStr := range splitRow {
			numberInt, err := utils.ConvertStringToInt(numberStr)
			if err != nil {
				panic("Unexpected error converting string to int")
			}
			newRow = Row{content: append(newRow.content, numberInt)}
		}
		structList = append(structList, newRow)
	}
	return structList
}

func DrawResults(structList []Row) {
	fmt.Println("Drawing results:")
	fmt.Println("UP    / DOWN")
	for _, item := range structList {
		for idx, number := range item.content {
			if item.directionUp[idx] == true {
				fmt.Printf("\033[32m%d\033[0m", number) // green
			} else {
				fmt.Printf("\033[31m%d\033[0m", number) // red
			}
		}
		fmt.Print(" / ")
		for idx, number := range item.content {
			if item.directionDown[idx] == true {
				fmt.Printf("\033[32m%d\033[0m", number) // green
			} else {
				fmt.Printf("\033[31m%d\033[0m", number) // red
			}
		}
		fmt.Println()
	}
}

func GetDirections(row Row) Row {
	for idx := range row.content {
		if idx == 0 {
			row.directionUp = append(row.directionUp, true)
			row.directionDown = append(row.directionDown, true)
			continue
		}
		prev := row.content[idx-1]
		curr := row.content[idx]
		diffUp := curr - prev
		if (prev < curr) && diffUp > 0 && diffUp < 4 {
			row.directionUp = append(row.directionUp, true)
		} else {
			row.directionUp = append(row.directionUp, false)
		}
		diffDown := prev - curr
		if (prev > curr) && diffDown > 0 && diffDown < 4 {
			row.directionDown = append(row.directionDown, true)
		} else {
			row.directionDown = append(row.directionDown, false)
		}
	}
	return row
}

func RowSafe(row Row) bool {
	tmpSaveUp := true
	for _, safe := range row.directionUp {
		if safe == false {
			tmpSaveUp = false
		}
	}
	if tmpSaveUp == true {
		return true
	}

	tmpSaveDown := true
	for _, safe := range row.directionDown {
		if safe == false {
			tmpSaveDown = false
		}
	}
	if tmpSaveDown == true {
		return true
	}
	return false
}

func removeAtIndex(s []int, i int) []int {
	out := make([]int, 0, len(s)-1)
	out = append(out, s[:i]...)
	out = append(out, s[i+1:]...)
	return out
}

func SolvePuzzle2() {
	retry := true
	puzzleIntput := utils.LoadFile("puzzle2024/puzzletext/puzzle2.txt")
	structList := FormatPuzzleInput(puzzleIntput)
	safeCount := 0

	for i := range structList {
		structList[i] = GetDirections(structList[i])
		fmt.Println(structList[i])
		if RowSafe(structList[i]) == true {
			safeCount++
		} else if retry == true {
			for n := len(structList[i].content) - 1; n >= 0; n-- {
				rowMissingOne := Row{content: removeAtIndex(structList[i].content, n)}
				rowMissingOne = GetDirections(rowMissingOne)
				if RowSafe(rowMissingOne) == true {
					structList[i] = rowMissingOne
					safeCount++
					break
				}
			}
		}
	}

	DrawResults(structList)
	fmt.Printf("There are %d safe rows\n", safeCount)
}
