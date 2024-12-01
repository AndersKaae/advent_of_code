package puzzle

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func convertStringToInt(input string) int {
	// Remove any whitespace
	input = strings.TrimSpace(input)

	intOutput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Failed processing %v and got error: %v", input, err)
	}
	return intOutput
}

func convertFileToTwoLists(content []string, listA []int, listB []int) ([]int, []int) {
	for i := 0; i < len(content); i++ {
		// Split the string by two spaces
		splitString := strings.Split(content[i], "  ")

		// Convert the string to int
		intA := convertStringToInt(splitString[0])
		intB := convertStringToInt(splitString[1])

		listA = append(listA, intA)
		listB = append(listB, intB)
	}
	return listA, listB
}

func puzzle1(listA []int, listB []int) {
	// Sorting the list lowest to highest
	sort.Ints(listA)
	sort.Ints(listB)

	listOfDiff := []int{}

	// Looping through the list and calculating the difference
	for i := 0; i < len(listA); i++ {
		tmp := listA[i] - listB[i]
		if tmp < 0 {
			tmp = tmp * -1
		}
		listOfDiff = append(listOfDiff, tmp)
	}

	totalDiff := 0
	for i := 0; i < len(listOfDiff); i++ {
		totalDiff += listOfDiff[i]
	}
	fmt.Println("The distance is", totalDiff)
}

func countInstancesInList(list []int, value int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			count++
		}
	}
	return count
}

func puzzle2(listA []int, listB []int) {
	alreadyCounted := []int{}
	sum := 0
	for i := 0; i < len(listA); i++ {
		// Check if the value is already counted
		if slices.Contains(alreadyCounted, i) == false {
			// Count the instances of the value in listB
			countA := countInstancesInList(listB, listA[i])
			result := countA * listA[i]
			sum += result
			alreadyCounted = append(alreadyCounted, i)
		}
	}
	fmt.Println("The sum is", sum)
}

func SolvePuzzle1() {
	// Load the file
	content := utils.LoadFile()

	listA := []int{}
	listB := []int{}

	listA, listB = convertFileToTwoLists(content, listA, listB)

	//listA = []int{3, 4, 2, 1, 3, 3}
	//listB = []int{4, 3, 5, 3, 9, 3}

	fmt.Println("Puzzle 1")
	puzzle1(listA, listB)
	fmt.Println("Puzzle 2")
	puzzle2(listA, listB)

}
