package puzzle2021

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
)

type Measurements struct {
	value int
	diff  int
}

func GetNumberIncreases(listOfMeasurements []Measurements) int {
	numberOfIncreases := 0
	for i := range listOfMeasurements {
		fmt.Println(listOfMeasurements[i])
		if listOfMeasurements[i].diff < 0 {
			numberOfIncreases += 1
		}
	}
	return numberOfIncreases
}

func SolvePuzzle1a(input []string) []Measurements {
	listOfMeasurements := []Measurements{}
	for i := range input {
		depth, err := utils.ConvertStringToInt(input[i])
		if err != nil {
			panic("Failed to convert string to int")

		}
		var meassurement Measurements
		if i == 0 {
			meassurement = Measurements{value: depth}
		} else {
			prevDepth := listOfMeasurements[i-1].value
			currentDiff := prevDepth - depth

			meassurement = Measurements{value: depth, diff: currentDiff}
		}
		listOfMeasurements = append(listOfMeasurements, meassurement)
	}
	numberOfIncreases := GetNumberIncreases(listOfMeasurements)
	fmt.Println(numberOfIncreases)
	return listOfMeasurements
}

func SolvePuzzle1b(oldLMeassurements []Measurements) {
	listOfMeassurements := []Measurements{}
	for i := range oldLMeassurements {
		if i > 1 {
			currentDepth := oldLMeassurements[i-2].value + oldLMeassurements[i-1].value + oldLMeassurements[i].value
			meassurement := Measurements{value: currentDepth}
			listOfMeassurements = append(listOfMeassurements, meassurement)
		}
	}
	for i := range listOfMeassurements {
		if i == 0 {
			continue
		}
		currentDiff := listOfMeassurements[i-1].value - listOfMeassurements[i].value
		listOfMeassurements[i].diff = currentDiff
	}
	numberOfIncreases := GetNumberIncreases(listOfMeassurements)
	fmt.Println(numberOfIncreases)
}

func SolvePuzzle1() {
	var input = utils.LoadFile("puzzle2021/puzzletext/puzzle1.txt")
	listOfMeassurements := SolvePuzzle1a(input)
	SolvePuzzle1b(listOfMeassurements)
}
