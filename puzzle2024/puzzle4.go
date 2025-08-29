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
	panic("Character not found!")

}

func DrawPuzzle4(characterList []Character) {
	for y := height - 1; y >= 0; y-- {
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

func CheckHorisontal(idx int, cl []Character) ([]Character, int) {
	wordCount := 0

	// TODO: Make sure that the offset to avoid out of bounds errors are correct

	if idx+2 < len(cl) && cl[idx+1].Content == "M" && cl[idx+2].Content == "A" && cl[idx+3].Content == "S" {
		cl[idx+0].ParOfWord = true
		cl[idx+1].ParOfWord = true
		cl[idx+2].ParOfWord = true
		cl[idx+3].ParOfWord = true
		wordCount++

	}
	if idx > 2 && cl[idx-1].Content == "M" && cl[idx-2].Content == "A" && cl[idx-3].Content == "S" {
		cl[idx-0].ParOfWord = true
		cl[idx-1].ParOfWord = true
		cl[idx-2].ParOfWord = true
		cl[idx-3].ParOfWord = true
		wordCount++
	}
	return cl, wordCount
}

func CheckVertical(idx int, cl []Character) ([]Character, int) {
	wordCount := 0
	if cl[idx-(width)].Content == "M" && cl[idx-(width*2)].Content == "A" && cl[idx-(width*3)].Content == "S" {
		cl[idx].ParOfWord = true
		cl[idx-(width*1)].ParOfWord = true
		cl[idx-(width*2)].ParOfWord = true
		cl[idx-(width*3)].ParOfWord = true
		wordCount++
	}
	return cl, wordCount
}

func DetectXmas(characterList []Character) ([]Character, int) {
	totalWordCount := 0
	for idx := range characterList {
		if characterList[idx].Content == "X" {
			wordCount := 0
			characterList, wordCount = CheckHorisontal(idx, characterList)
			totalWordCount += wordCount
			characterList, wordCount = CheckVertical(idx, characterList)
		}
	}
	return characterList, totalWordCount
}

func SolvePuzzle4() {
	totalWordCount := 0
	var input = utils.LoadFile("puzzle2024/puzzletext/puzzle4sample.txt")
	height = len(input)
	width = len(input[0])
	characterList := CreateCharacterStructs(input)
	characterList, totalWordCount = DetectXmas(characterList)
	DrawPuzzle4(characterList)
	fmt.Println(totalWordCount)
}
