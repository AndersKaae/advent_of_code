package puzzle2025

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AndersKaae/advent_of_code/utils"
)

type Product struct {
	source             string
	firstIdInt         int
	secondIdInt        int
	possibleSeparation []string
}

func CreatStruct() []Product {
	puzzleInput := utils.LoadFile("puzzle2025/puzzletext/puzzle2sample.txt")
	var puzzleStr string
	puzzleStr = puzzleInput[0]
	listOfIds := strings.Split(puzzleStr, ",")

	productList := []Product{}

	for i := range listOfIds {
		Ids := strings.Split(listOfIds[i], "-")
		firstIdInt, err := utils.ConvertStringToInt(Ids[0])
		if err != nil {
			panic("Unable to convert to int")
		}
		secondInt, err := utils.ConvertStringToInt(Ids[1])
		if err != nil {
			panic("Unable to convert to int")
		}
		product := Product{listOfIds[i], firstIdInt, secondInt, []string{}}
		productList = append(productList, product)
	}
	return productList
}

func SolvePuzzle2() {
	puzzle := "A"
	productList := CreatStruct()
	for i := range productList {
		fmt.Println(productList[i])
		for n := productList[i].firstIdInt; n <= productList[i].secondIdInt; n++ {
			nStr := strconv.Itoa(n)

			var validDivisions []int

			if puzzle == "A" {
				validDivisions = []int{2}
				if len(nStr)%2 != 0 {
					continue
				}
			} else {
				validDivisions = getValidDivisions(strconv.Itoa(n))
			}

			for z := range validDivisions {
				productList[i].possibleSeparation = splitString(z, nStr)
			}
		}
	}
}

func splitString(n int, s string) []string {
	var result []string

	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}

	return result
}

func getValidDivisions(s string) []int {
	validDivision := []int{}
	lengthS := len(s)
	for i := 1; i < lengthS; i += 1 {
		if lengthS%i == 0 {
			validDivision = append(validDivision, i)
		}
	}
	return validDivision
}

func IsInvalidId(id int) bool {
	idStr := strconv.Itoa(id)
	idLenght := len(idStr)
	if idLenght%2 == 0 {
		idx := idLenght / 2
		leftStr := idStr[:idx]
		rightStr := idStr[idx:]
		if leftStr == rightStr {
			return true
		}
	}
	return false
}
