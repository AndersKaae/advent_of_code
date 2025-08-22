package puzzle2022

import (
	//"bufio"
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"math"
	//"os"
	"strings"
)

type Rope struct {
	Location Coordinates
	Visited  []Coordinates
}

type Coordinates struct {
	X int
	Y int
}

type Move struct {
	Direction string
	Places    int
}

type Board struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
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

func isRopeTaut(ropeHead Rope, ropeTail Rope) bool {
	dx := int(math.Abs(float64(ropeHead.Location.X - ropeTail.Location.X)))
	dy := int(math.Abs(float64(ropeHead.Location.Y - ropeTail.Location.Y)))
	if dx > 1 || dy > 1 {
		return true
	}
	return false
}

func MoveTail(ropeHead Rope, ropeTail Rope) Rope {
	if !isRopeTaut(ropeHead, ropeTail) {
		return ropeTail
	}
	ropeTail.Location = ropeHead.Visited[len(ropeHead.Visited)-1]
	return ropeTail
}

func GetUniqueTailPositions(ropeTail Rope) int {
	uniquCoords := []Coordinates{}
	for _, coords := range ropeTail.Visited {
		found := false
		for _, uCord := range uniquCoords {
			if uCord == coords {
				found = true
			}
		}
		if found == false {
			uniquCoords = append(uniquCoords, coords)
		}
	}
	return len(uniquCoords)
}

func DrawMovements(ropeHead Rope, ropeTail Rope) {
	// Finding the outer bounds
	board := Board{MinX: 0, MaxX: 0, MinY: 0, MaxY: 0}
	for _, move := range ropeHead.Visited {
		if move.X < board.MinX {
			board.MinX = move.X
		}
		if move.X > board.MaxX {
			board.MaxX = move.X
		}
		if move.Y < board.MinY {
			board.MinY = move.Y
		}
		if move.Y > board.MaxY {
			board.MaxY = move.Y
		}
	}
	fmt.Println(board)

	for idx, move := range ropeHead.Visited {
		fmt.Println(move)
		for y := board.MaxY; board.MinX <= y; y-- {
			for x := board.MinX; board.MaxX >= x; x++ {
				if move.X == x && move.Y == y {
					fmt.Printf("H")
				} else if ropeTail.Visited[idx].X == x && ropeTail.Visited[idx].Y == y {
					fmt.Printf("T")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
		//bufio.NewReader(os.Stdin).ReadBytes('\n')
		fmt.Println()
	}
}

func CreateTailStructs(number int) []Rope {
	ropeTailList := []Rope{}
	ropeTail := Rope{
		Location: Coordinates{X: 0, Y: 0},
		Visited:  []Coordinates{{X: 0, Y: 0}},
	}
	for i := 0; i < number; i++ {
		ropeTailList = append(ropeTailList, ropeTail)
	}
	return ropeTailList
}

func SolvePuzzle9() {
	var tailNumber int = 9
	input := utils.LoadFile("puzzle2022/puzzletext/test9.txt")
	movesList := CreateMovesStrucs(input)
	ropeTailList := CreateTailStructs(tailNumber)
	ropeTail := ropeTailList[0] // TODO This should be deleted
	ropeHead := Rope{
		Location: Coordinates{X: 0, Y: 0},
		Visited:  []Coordinates{{X: 0, Y: 0}},
	}
	for _, move := range movesList {
		fmt.Println("MOVE: ", move)
		destination := CalcDestination(move, ropeHead)
		for true {
			if isAtDestination(destination, ropeHead) {
				break
			}
			ropeHead = MoveHead(move, ropeHead)
			ropeTail = MoveTail(ropeHead, ropeTail)
			ropeHead.Visited = append(ropeHead.Visited, Coordinates{X: ropeHead.Location.X, Y: ropeHead.Location.Y})
			ropeTail.Visited = append(ropeTail.Visited, Coordinates{X: ropeTail.Location.X, Y: ropeTail.Location.Y})
		}
	}
	//DrawMovements(ropeHead, ropeTail)
	number := GetUniqueTailPositions(ropeTail)
	fmt.Println(number)
}
