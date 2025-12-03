package puzzle2021

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strings"
)

type Movement struct {
	direction    string
	movementSize int
}

type Coordinates struct {
	horizontalPos int
	verticalPos   int
	aim           int
}

func importData(filename string) []Movement {
	movementList := []Movement{}
	var input = utils.LoadFile(filename)
	for i := range input {
		splitString := strings.Split(input[i], " ")
		intMovemment, err := utils.ConvertStringToInt(splitString[1])
		if err != nil {
			panic("Cannot convert to int")
		}
		movement := Movement{direction: splitString[0], movementSize: intMovemment}
		movementList = append(movementList, movement)
	}
	return movementList
}

func SolvePuzzle2a(movementList []Movement) {
	coordinates := Coordinates{horizontalPos: 0, verticalPos: 0}
	for i := range movementList {
		if movementList[i].direction == "forward" {
			coordinates.verticalPos += movementList[i].movementSize
		} else if movementList[i].direction == "up" {
			coordinates.horizontalPos -= movementList[i].movementSize
		} else if movementList[i].direction == "down" {
			coordinates.horizontalPos += movementList[i].movementSize
		} else {
			panic("Unknow direction")
		}
	}
	fmt.Println(coordinates)
	product := coordinates.horizontalPos * coordinates.verticalPos
	fmt.Println(product)
}

func SolvePuzzle2b(movementList []Movement) {
	coordinates := Coordinates{horizontalPos: 0, verticalPos: 0, aim: 0}
	for i := range movementList {
		if movementList[i].direction == "forward" {
			coordinates.verticalPos += movementList[i].movementSize
			if i != 0 && movementList[i-1].direction == "up" {
				coordinates.horizontalPos -= movementList[i].movementSize * coordinates.aim
			} else if i != 0 && movementList[i-1].direction == "down" {
				coordinates.horizontalPos += movementList[i].movementSize * coordinates.aim
			}
		} else if movementList[i].direction == "up" {
			coordinates.horizontalPos -= movementList[i].movementSize
			coordinates.aim -= movementList[i].movementSize
		} else if movementList[i].direction == "down" {
			coordinates.horizontalPos += movementList[i].movementSize
			coordinates.aim += movementList[i].movementSize
		} else {
			panic("Unknow direction")
		}
		fmt.Println(movementList[i])
		fmt.Println(coordinates)
	}
	fmt.Println("hPos:" + string(coordinates.horizontalPos))
	product := coordinates.horizontalPos * coordinates.verticalPos
	fmt.Println(product)
}

func SolvePuzzle2() {
	movementList := importData("puzzle2021/puzzletext/puzzle2sample.txt")
	//SolvePuzzle2a(movementList)
	SolvePuzzle2b(movementList)
}
