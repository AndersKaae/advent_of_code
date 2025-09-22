package puzzle2024

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
)

type Character struct {
	X         int
	Y         int
	Content   string
	ParOfWord bool
}

type Direction struct {
	Direction string
	Sucess    bool
	Character Character
}

var (
	height int
	width  int
)

func CreateCharacterStructs(input []string) []Character {
	charaterList := []Character{}
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			char := Character{X: x, Y: y, Content: string(input[y][x]), ParOfWord: false}
			charaterList = append(charaterList, char)

		}
	}
	return charaterList
}

func FindIdxCharacter(x int, y int, characterList []Character) int {
	for idx, character := range characterList {
		if character.X == x && character.Y == y {
			return idx
		}
	}
	return -1
}

func DrawPuzzle4(characterList []Character) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := FindIdxCharacter(x, y, characterList)
			if characterList[idx].ParOfWord == false {
				fmt.Printf("\033[31m%s\033[0m", characterList[idx].Content) // red
			} else {
				fmt.Printf("\033[32m%s\033[0m", characterList[idx].Content) // green
			}
		}
		println()
	}
}

func CreateDirectionsStruct() []Direction {
	directions := []string{"up", "down", "left", "right", "upLeft", "upRight", "downLeft", "downRight"}
	directionStruct := []Direction{}
	for _, dir := range directions {
		directionStruct = append(directionStruct, Direction{Direction: dir, Sucess: false})
	}
	return directionStruct
}

func GetDirection(x int, y int, direction Direction, cl []Character) direction {
	idx := -1
	if direction.Direction == "up" {
		idx = FindIdxCharacter(x, y+1, cl)
	} else if direction.Direction == "down" {
		idx = FindIdxCharacter(x, y-1, cl)
	} else if direction.Direction == "left" {
		idx = FindIdxCharacter(x-1, y, cl)
	} else if direction.Direction == "right" {
		idx = FindIdxCharacter(x+1, y, cl)
	} else if direction.Direction == "upLeft" {
		idx = FindIdxCharacter(x-1, y+1, cl)
	} else if direction.Direction == "upRight" {
		idx = FindIdxCharacter(x+1, y+1, cl)
	} else if direction.Direction == "downLeft" {
		idx = FindIdxCharacter(x-1, y-1, cl)
	} else if direction.Direction == "downRight" {
		idx = FindIdxCharacter(x-1, y-1, cl)
	} else {
		panic("Unknown direction")
	}
	if cl[idx].Content == letter {

	}
	return idx
}

func DetectXmas(cl []Character) ([]Character, int) {
	totalWordCount := 0
	directionsStruct := CreateDirectionsStruct()
	lettersNeeded := []string{"M", "A", "S"}
	for idx := range cl {
		if cl[idx].Content == "X" {
			x := cl[idx].X
			y := cl[idx].Y

			for _, direction := range directionsStruct {

			}
		}
	}
	return characterList, totalWordCount
}

func SolvePuzzle4() {
	totalWordCount := 0
	var input = utils.LoadFile("puzzle2024/puzzletext/puzzle4.txt")
	height = len(input)
	width = len(input[0])
	characterList := CreateCharacterStructs(input)
	characterList, totalWordCount = DetectXmas(characterList)
	DrawPuzzle4(characterList)
	fmt.Println(totalWordCount)
}
