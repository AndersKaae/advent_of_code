package puzzle2025

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strconv"
)

type Rotation struct {
	Value     string
	Direction string
	Movement  int
	Position  int
}

func CreateStruct() []Rotation {
	rotationList := []Rotation{}
	puzzleInput := utils.LoadFile("puzzle2025/puzzletext/puzzle1sample.txt")
	for i := range puzzleInput {
		movementStr := string(puzzleInput[i][1:])
		movementInt, err := utils.ConvertStringToInt(movementStr)
		if err != nil {
			panic("Cant convert to number")
		}
		rotation := Rotation{Value: puzzleInput[i], Direction: string(puzzleInput[i][0]), Movement: movementInt}
		rotationList = append(rotationList, rotation)
	}
	return rotationList
}

func RotateTheDial(dialValue int, rotation *Rotation) int {
	hitZero := 0
	for i := 0; i < rotation.Movement; i++ {
		if rotation.Direction == "L" {
			if dialValue == 0 {
				if i != 0 {
					hitZero += 1
				}
				dialValue = 99
			} else {
				dialValue -= 1
			}
		} else if rotation.Direction == "R" {
			if dialValue == 99 {
				dialValue = 0
				hitZero += 1
			} else {
				dialValue += 1
			}
		} else {
			panic("Unexpected direction")
		}
	}
	rotation.Position = dialValue
	return hitZero
}

func SolvePuzzle1() {
	dialValue := 50
	stoppedAtZeroCount := 0
	hitZeroCount := 0
	rotationList := CreateStruct()
	for i := range rotationList {
		hitZero := RotateTheDial(dialValue, &rotationList[i])
		fmt.Println(rotationList[i])
		if hitZero > 0 {
			fmt.Println("Hit Zero (" + strconv.Itoa(hitZero) + ")")
		}
		hitZeroCount += hitZero
		dialValue = rotationList[i].Position
		if dialValue == 0 {
			stoppedAtZeroCount += 1
		}
	}
	fmt.Println(stoppedAtZeroCount)
	fmt.Println(hitZeroCount)
}
