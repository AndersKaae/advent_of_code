package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strings"
)

type Rope struct {
	Location Coordinates
	Visited  Coordinates
}

type Coordinates struct {
	X int
	Y int
}

type Move struct {
	Direction string
	Places    int
}

func CreateMovesStrucs(input []string) []Move {
	movesList := []Move{}
	for _, row := range input {
		rowSplit := strings.Split(row, " ")
		newMove := Move{Direction: rowSplit[0], Places: utils.ConvertStringToInt(rowSplit[1])}
		movesList = append(movesList, newMove)
	}
	return movesList
}

func CalcDestination(move Move, ropeHead Rope) Coordinates {
	destination := Coordinates{X: ropeHead.Location.X, Y: ropeHead.Location.Y}
	if move.Direction == "U" {
		destination.Y += move.Places
	} else if move.Direction == "D" {
		destination.Y -= move.Places
	} else if move.Direction == "R" {
		destination.X += move.Places
	} else if move.Direction == "L" {
		destination.X -= move.Places
	} else {
		panic("Move type  now found")
	}
	return destination
}

func isAtDestination(destination Coordinates, rope Rope) bool {
	if destination.X == rope.Location.X && destination.Y == rope.Location.Y {
		return true
	}
	return false
}

func MoveHead(move Move, rope Rope) Rope {
	if move.Direction == "U" {
		rope.Location.Y += 1
	} else if move.Direction == "D" {
		rope.Location.Y -= 1
	} else if move.Direction == "R" {
		rope.Location.X += 1
	} else if move.Direction == "L" {
		rope.Location.X -= 1
	} else {
		msg, err := fmt.Println("Invalid move:", move.Direction)
		if err != nil {
			panic("Could not format string")
		}
		panic(msg)
	}
	return rope
}

func SolvePuzzle9() {
	input := utils.LoadFile("puzzle2022/puzzletext/test9.txt")
	movesList := CreateMovesStrucs(input)
	ropeHead := Rope{
		Location: Coordinates{X: 0, Y: 0},
	}
	for _, move := range movesList {
		fmt.Println("MOVE: ", move)
		destination := CalcDestination(move, ropeHead)
		for true {
			if isAtDestination(destination, ropeHead) {
				break
			}
			ropeHead = MoveHead(move, ropeHead)
			fmt.Println("Rope", ropeHead, ", Destination", destination)
		}
	}

}
