package puzzle2024

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
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

func isSkipSafe(a int, b int) bool {
	if a == b {
		return false
	}
	diff := a - b
	if diff < 0 {
		diff = diff * -1
	}
	if diff > 3 {
		return false
	}
	return true
}

func processReport(report []int, tolerance int) bool {
	safe := true
	direction := ""
	for j := 0; j < len(report)-1; j++ {
		a := report[j]
		b := report[j+1]
		safe = isSkipSafe(a, b)
		if safe == false {
			return false
		}
		if direction == "" {
			if a > b {
				direction = "down"
			} else {
				direction = "up"
			}
		}
		if direction == "down" && a < b {
			return false
		}
		if direction == "up" && a > b {
			return false
		}
	}
	return true
}

func puzzle3(reportList [][]int, tolerance int) {

	numberSaveReports := 0

	for i := 0; i < len(reportList); i++ {
		report := reportList[i]
		fmt.Println(report)
		result := processReport(report, tolerance)
		if result == true {
			numberSaveReports++
		}
	}
	fmt.Println("Number of safe reports", numberSaveReports)
}

func SolvePuzzle2() {
	content := utils.LoadFile("puzzle2024/puzzle2.txt")

	// Convert the file data to a list of list of ints
	reportList := createReportList(content)

	puzzle3(reportList, 0)
}
