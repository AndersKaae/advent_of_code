package puzzle2023

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
)

type Place struct {
	x       int
	y       int
	char    rune
	kind    string
	istouch bool
}

func ConstructObject(path string) []Place {
	numbers := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	input := utils.LoadFile(path)
	placesList := []Place{}
	for y, row := range input {
		splitRunes := []rune(row)
		for x, char := range splitRunes {
			kind := ""
			if utils.Contains(numbers, char) {
				kind = "number"
			} else if char == '.' {
				kind = "space"
			} else {
				kind = "symbol"
			}
			place := Place{x, y, char, kind, false}
			placesList = append(placesList, place)
		}
	}
	return placesList
}

func FindPlace(placeList []Place, x int, y int) int {
	for idx, place := range placeList {
		if place.x == x && place.y == y {
			return idx
		}
	}
	return -1
}

func Neighbor(place Place, placeList []Place) []int {
	above := FindPlace(placeList, place.x, place.y-1)
	below := FindPlace(placeList, place.x, place.y+1)
	left := FindPlace(placeList, place.x-1, place.y)
	right := FindPlace(placeList, place.x+1, place.y)
	aboveLeft := FindPlace(placeList, place.x-1, place.y-1)
	aboveRight := FindPlace(placeList, place.x+1, place.y-1)
	belowLeft := FindPlace(placeList, place.x-1, place.y+1)
	belowRight := FindPlace(placeList, place.x+1, place.y+1)
	return []int{above, below, left, right, aboveLeft, aboveRight, belowLeft, belowRight}
}

func PrintColored(place Place) {
	if place.kind == "number" && place.istouch == false {
		fmt.Printf("\033[31m%c\033[0m", place.char) // red
	} else if place.kind == "number" && place.istouch == true {
		fmt.Printf("\033[32m%c\033[0m", place.char) // green
	} else {
		fmt.Printf("%c", place.char)
	}
}

func DrawResult(placeList []Place) {
	fmt.Println("")
	currentLine := 0
	for _, place := range placeList {
		PrintColored(place)
		if place.y != currentLine {
			fmt.Println("")
			currentLine = place.y
		}
	}
}

func WhoIsTouching(placeList []Place) []Place {
	for idx, place := range placeList {
		neighborIdxs := Neighbor(place, placeList)
		for _, neighbor := range neighborIdxs {
			if neighbor == -1 {
				continue
			}
			if placeList[neighbor].kind == "symbol" {
				placeList[idx].istouch = true
			}
		}
	}
	return placeList
}

func MarkRestOfNumber(placeList []Place) ([]Place, int) {
	sumOfNumbers := 0
	listOfIdx := []int{}
	for idx, place := range placeList {
		if place.kind == "number" {
			listOfIdx = append(listOfIdx, idx)
		} else if place.kind != "letter" {
			if len(listOfIdx) > 0 {
				// We are determining if any of the numbers next to each other are touching a symbol
				isTouching := false
				for _, idx := range listOfIdx {
					if placeList[idx].istouch == true {
						isTouching = true
					}
				}
				concatString := ""
				if isTouching == true {
					for _, idx := range listOfIdx {
						placeList[idx].istouch = true
						concatString += string(placeList[idx].char)
					}
					integer, err := utils.ConvertStringToInt(concatString)
					if err != nil {
						fmt.Println("Not a valid number")
					}
					sumOfNumbers += integer
				}
				listOfIdx = []int{}
			}
		}

	}

	return placeList, sumOfNumbers
}

func SolvePuzzle3() {
	placeList := ConstructObject("puzzle2023/puzzletext/puzzle3a.txt")
	placeList = WhoIsTouching(placeList)
	placeList, sumOfNumbers := MarkRestOfNumber(placeList)
	fmt.Println(sumOfNumbers)
	DrawResult(placeList)
}
