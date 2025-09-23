package puzzle2023

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strconv"
	"strings"
	"unicode"
)

func getFirstNumber(input []string, wordNumbers []string) []string {
	var output []string
	for j := 0; j < len(input); j++ {
		fmt.Println("input", input[j])
		runes := []rune(input[j])
		var number string
		var partialNumber string
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				number = string(runes[i])
				break
			} else {
				partialNumber += string(runes[i])
				for k := 0; k < len(wordNumbers); k++ {
					if strings.Contains(partialNumber, wordNumbers[k]) {
						number = strconv.Itoa(k)
						partialNumber = ""
						break
					}
				}
			}
			if number != "" {
				break
			}
		}
		output = append(output, number)
	}
	return output
}

func reverseEachItemInSlice(input []string) []string {
	var output []string
	for i := 0; i < len(input); i++ {
		runes := []rune(input[i])
		var reversed string
		for j := len(runes) - 1; j >= 0; j-- {
			reversed += string(runes[j])
		}
		output = append(output, reversed)
	}
	return output
}

func puzzle1(input []string) {
	//input = []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	sum := 0

	for i := 0; i < len(input); i++ {
		// Convert the string to a slice of runes
		runes := []rune(input[i])

		var digit1, digit2 string
		for j := 0; j < len(runes); j++ {
			if unicode.IsDigit(runes[j]) {
				digit1 = string(runes[j])
				break
			}
		}

		for j := len(runes) - 1; j >= 0; j-- {
			if unicode.IsDigit(runes[j]) {
				digit2 = string(runes[j])
				break
			}
		}

		result, err := utils.ConvertStringToInt(digit1 + digit2)
		if err != nil {
			fmt.Println("Error converting to int:", err)
		}
		sum += resulat

	}
	fmt.Println("The sum is", sum)
}

func puzzle2(input []string) {
	wordNumbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	//input = []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	first := getFirstNumber(input, wordNumbers)
	fmt.Println(first)
	reversedInput := reverseEachItemInSlice(input)
	reversedNumbers := reverseEachItemInSlice(wordNumbers)
	last := getFirstNumber(reversedInput, reversedNumbers)
	fmt.Println(last)

	sum := 0

	for i := 0; i < len(first); i++ {
		result := first[i] + last[i]
		fmt.Println("The result is", result)
		subSum, err := utils.ConvertStringToInt(result)
		if err != nil {
			fmt.Println("Error converting to int:", err)
		}
		sum += subSum
	}

	fmt.Println("The sum is", sum)
}

func SolvePuzzle1() {
	// Read the file
	input := []string{}

	input = utils.LoadFile("puzzle2023/puzzle1.txt")
	puzzle1(input)
	puzzle2(input)

}
