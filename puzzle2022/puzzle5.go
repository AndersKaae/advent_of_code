package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strconv"
	"strings"
)

type Movement struct {
	Quantity int
	From     int
	To       int
}

var (
	//columnsWithLetters = []int{1, 5, 9, 13, 17, 21, 25, 29, 34}
	columnsWithLetters = []int{1, 5, 9}
)

func RemoveEmptySpaces(stack []string) []string {
	cleanStack := []string{}
	for _, item := range stack {
		if strings.TrimSpace(item) != "" {
			cleanStack = append(cleanStack, item)
		}
	}
	return cleanStack
}

func CreateMovementsStruct() []Movement {
	var input = utils.LoadFile("puzzle2022/puzzle5.txt")
	listOfMovements := []Movement{}
	for _, row := range input {
		splitRow := strings.Fields(row)
		if len(splitRow) == 0 {
			continue
		}
		if splitRow[0] == "move" {
			fmt.Println(row)
			movement := Movement{
				Quantity: utils.ConvertStringToInt(splitRow[1]),
				From:     utils.ConvertStringToInt(splitRow[3]),
				To:       utils.ConvertStringToInt(splitRow[5]),
			}
			listOfMovements = append(listOfMovements, movement)
		}
	}
	return listOfMovements
}

func CreateCrates() [][]string {
	var input = utils.LoadFile("puzzle2022/puzzle5.txt")
	listOfRows := [][]string{}
	for _, row := range input {
		fmt.Println(row)
		rowList := []string{}
		rowRunes := []rune(row)
		for _, r := range rowRunes {
			rowList = append(rowList, string(r))
		}
		if len(rowList) == 0 {
			break
		}
		listOfRows = append(listOfRows, rowList)
	}

	listOfColumns := [][]string{}
	for _, column := range columnsWithLetters {
		stack := []string{}
		for _, row := range listOfRows {
			stack = append(stack, row[column])
		}
		stackNoSpaces := RemoveEmptySpaces(stack)
		listOfColumns = append(listOfColumns, stackNoSpaces)
	}
	return listOfColumns
}

func GetColumn(columnNumber int, listOfColumns [][]string) int {
	for idx, column := range listOfColumns {
		columnStr := strconv.Itoa(columnNumber)
		if utils.Contains(column, columnStr) {
			return idx
		}
	}
	errorMsg := "We did not find column " + strconv.Itoa(columnNumber)
	panic(errorMsg)
}

func TakeCratesFromColumn(quantity int, takeIdx int, listOfColumns [][]string) ([]string, []string) {
	targetColumn := listOfColumns[takeIdx]
	takeOff := targetColumn[:quantity]
	remainder := targetColumn[quantity:]
	return takeOff, remainder
}

func SolvePuzzle5() {
	listOfColumns := CreateCrates()
	fmt.Println(listOfColumns)
	listOfMovements := CreateMovementsStruct()
	for _, movement := range listOfMovements {
		fmt.Println("Movement:", movement)
		columnTakeFrom := GetColumn(movement.From, listOfColumns)
		takeOff, remainder := TakeCratesFromColumn(movement.Quantity, columnTakeFrom, listOfColumns)
		fmt.Println("Taking", takeOff, "from", listOfColumns[columnTakeFrom])
		listOfColumns[columnTakeFrom] = remainder
		fmt.Println("Now ", listOfColumns[columnTakeFrom])
		columnDeliverTo := GetColumn(movement.To, listOfColumns)
		fmt.Println("Adding", takeOff, "to", listOfColumns[columnDeliverTo])
		// TODO:We need to reverse takeoff before adding it to the stack
		appendedColumn := append(append([]string{}, takeOff...), listOfColumns[columnDeliverTo]...)
		listOfColumns[columnDeliverTo] = appendedColumn
		fmt.Println("---", listOfColumns)
	}
}
