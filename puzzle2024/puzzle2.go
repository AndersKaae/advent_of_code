package puzzle2024

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AndersKaae/advent_of_code/utils"
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

func isSkipSafe(report []int, mappedInvalids []bool) []bool {
	difflist := []int{}
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if diff < 0 {
			diff = diff * -1
		}
		difflist = append(difflist, diff)
	}
	for i := 0; i < len(difflist)-1; i++ {
		if difflist[i] > 3 {
			mappedInvalids[i] = false
		}
	}
	return mappedInvalids
}

func getDirection(report []int) string {
	up := 0
	down := 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if diff > 0 {
			down++
		} else {
			up++
		}
	}
	if down > up {
		return "down"
	}
	return "up"
}

func isDirectionValid(report []int, direction string, mappedInvalids []bool) []bool {
	listOfInvalids := []bool{}
	for i := 0; i < len(report)-1; i++ {
		if direction == "down" {
			if report[i] < report[i+1] {
				listOfInvalids = append(listOfInvalids, false)
			} else {
				listOfInvalids = append(listOfInvalids, true)
			}
		}
		if direction == "up" {
			if report[i] > report[i+1] {
				listOfInvalids = append(listOfInvalids, false)
			} else {
				listOfInvalids = append(listOfInvalids, true)
			}
		}
	}
	for i := 0; i < len(listOfInvalids); i++ {
		if listOfInvalids[i] == false {
			mappedInvalids[i] = false
		}
	}
	return mappedInvalids
}

func removeFirstFalseElement(report []int, mappedInvalids []bool) []int {
	report2 := []int{}
	deleted := false
	for i := 0; i < len(mappedInvalids); i++ {
		if mappedInvalids[i] == false && deleted == false {
			deleted = true
			continue
		}
		report2 = append(report2, report[i])
	}
	fmt.Println(report2)
	return report2
}

func createMapsinvalids(report []int) []bool {
	mappedInvalids := []bool{}
	for i := 0; i < len(report); i++ {
		mappedInvalids = append(mappedInvalids, true)
	}
	return mappedInvalids
}

func processReport(report []int) bool {
	direction := getDirection(report)
	mappedInvalids := []bool{}
	problemDampenerUsed := false

	fmt.Println(report)
	mappedInvalids = createMapsinvalids(report)
	for {
		mappedInvalids = isDirectionValid(report, direction, mappedInvalids)
		mappedInvalids = isSkipSafe(report, mappedInvalids)
		if problemDampenerUsed == false && slices.Contains(mappedInvalids, false) {
			problemDampenerUsed = true
			report = removeFirstFalseElement(report, mappedInvalids)
			mappedInvalids = createMapsinvalids(report)
			continue
		}
		break
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
		fmt.Println("-------")
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
