package puzzle2024

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"slices"
	"strings"
)

func createReportList(content []string) [][]int {
	reportList := [][]int{}
	for i := 0; i < len(content); i++ {
		numbers := strings.Split(content[i], " ")

		report := []int{}
		for j := 0; j < len(numbers); j++ {
			numberInt := utils.ConvertStringToInt(numbers[j])
			report = append(report, numberInt)
		}
		reportList = append(reportList, report)
	}
	return reportList
}

func isSkipSafe(direction string, a int, b int) int {
	diff := a - b
	if diff < 0 {
		diff = diff * -1
	}
	if diff > 3 {
		if direction == "up" {
			return 2
		}
		return 1
	}
	return 0
}

func getDirection(report []int) string {
	sum := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		sum += diff
	}
	if sum > 0 {
		return "down"
	}
	return "up"
}

func isDirectionValid(a int, b int, direction string) int {
	if direction == "up" && a > b {
		return 1
	}
	if direction == "down" && a < b {
		return 1
	}
	if a == b {
		return 2
	}
	return 0
}

func mapInvalids(pos int, invalidPos int, mappedInvalids []bool) []bool {
	if invalidPos == 1 {
		mappedInvalids[pos] = false
		mappedInvalids[pos+1] = true
	}
	if invalidPos == 2 {
		mappedInvalids[pos] = true
		mappedInvalids[pos+1] = false
	}
	return mappedInvalids
}

func removeFirstFalseElement(report []int, mappedInvalids []bool) ([]int, []bool) {
	for i := 0; i < len(mappedInvalids); i++ {
		if mappedInvalids[i] == false {
			// Remove the first false element
			mappedInvalids = append(mappedInvalids[:i], mappedInvalids[i+1:]...)
			report = append(report[:i], report[i+1:]...)
			break
		}
	}
	return report, mappedInvalids
}

func processReport(report []int) bool {
	direction := getDirection(report)
	mappedInvalids := []bool{}
	problemDampenerUsed := false
	for i := 0; i < len(report); i++ {
		mappedInvalids = append(mappedInvalids, true)
	}
	for j := 0; j < len(report)-1; j++ {
		fmt.Println(report)
		a := report[j]
		b := report[j+1]
		invalidPos := isDirectionValid(a, b, direction)
		mappedInvalids = mapInvalids(j, invalidPos, mappedInvalids)
		invalidPos = isSkipSafe(direction, a, b)
		mappedInvalids = mapInvalids(j, invalidPos, mappedInvalids)
		if problemDampenerUsed == false && slices.Contains(mappedInvalids, false) {
			problemDampenerUsed = true
			report, mappedInvalids = removeFirstFalseElement(report, mappedInvalids)
			j = -1
			continue
		}
	}

	fmt.Println(mappedInvalids)
	if slices.Contains(mappedInvalids, false) {
		return false
	}
	return true
}

func puzzle3(reportList [][]int) {
	numberSaveReports := 0
	for i := 0; i < len(reportList); i++ {
		report := reportList[i]
		result := processReport(report)
		if result == true {
			numberSaveReports++
			fmt.Println("Safe")
		} else {
			fmt.Println("Not safe")
		}
	}
	fmt.Println("Number of safe reports", numberSaveReports)
}

func SolvePuzzle2() {
	content := utils.LoadFile("puzzle2024/puzzle2sample.txt")

	// Convert the file data to a list of list of ints
	reportList := createReportList(content)

	puzzle3(reportList)
}
