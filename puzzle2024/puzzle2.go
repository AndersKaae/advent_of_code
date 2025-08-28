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
			numberInt := utils.ConvertStringToInt(numberStr)
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

func CountSaveReports(structList []Row) int {
	saveReportUp := 0
	saveReportDown := 0
	for _, row := range structList {
		tmpSaveUp := true
		for _, safe := range row.directionUp {
			if safe == false {
				tmpSaveUp = false
			}
		}
		if tmpSaveUp == true {
			saveReportUp++
		}

		tmpSaveDown := true
		for _, safe := range row.directionDown {
			if safe == false {
				tmpSaveDown = false
			}
		}
		if tmpSaveDown == true {
			saveReportDown++
		}
	}
	return saveReportDown + saveReportUp
}

func SolvePuzzle2() {
	puzzleIntput := utils.LoadFile("puzzle2024/puzzletext/puzzle2sample.txt")
	structList := FormatPuzzleInput(puzzleIntput)

	for i := range structList {
		fmt.Println(structList[i])
		structList[i] = GetDirections(structList[i]) // write back
		fmt.Println(structList[i])
	}

	DrawResults(structList)
	fmt.Println(CountSaveReports(structList))
}
