package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strconv"
	"strings"
)

type ElfPair struct {
	Raw          string
	Elf1         []int
	Elf2         []int
	Is1Contained bool
	Is2Contained bool
}

func GetNumberBetween(fromTo string) []int {
	result := strings.Split(fromTo, "-")
	from, err := strconv.Atoi(result[0])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(result[1])
	if err != nil {
		panic(err)
	}
	listOfNumbers := []int{}
	for from <= to {
		listOfNumbers = append(listOfNumbers, from)
		from++
	}
	return listOfNumbers

}

func IsContained(toTest []int, container []int) bool {
	for _, number := range toTest {
		if utils.Contains(container, number) == false {
			return false
		}
	}
	return true
}

func IsPartiallyContained(toTest []int, container []int) bool {
	for _, number := range toTest {
		if utils.Contains(container, number) == true {
			return true
		}
	}
	return false
}

func CreateStruct() []ElfPair {
	var input = utils.LoadFile("puzzle2022/puzzle4.txt")
	elfPairs := []ElfPair{}
	for _, pair := range input {
		elfPair := ElfPair{Raw: pair}
		splitPair := strings.Split(pair, ",")
		elfPair.Elf1 = GetNumberBetween(splitPair[0])
		elfPair.Elf2 = GetNumberBetween(splitPair[1])
		elfPairs = append(elfPairs, elfPair)
	}
	return elfPairs
}

func SolvePuzzle4part1() {
	containedPairs := 0
	elfPairs := CreateStruct()
	for _, pair := range elfPairs {
		pair.Is1Contained = IsContained(pair.Elf1, pair.Elf2)
		pair.Is2Contained = IsContained(pair.Elf2, pair.Elf1)
		if pair.Is1Contained == true {
			containedPairs += 1
			continue
		}
		if pair.Is2Contained == true {
			containedPairs += 1
			continue
		}
	}
	fmt.Print("Part1: ", containedPairs)
	fmt.Println()
}

func SolvePuzzle4part2() {
	containedPairs := 0
	elfPairs := CreateStruct()
	for _, pair := range elfPairs {
		pair.Is1Contained = IsPartiallyContained(pair.Elf1, pair.Elf2)
		pair.Is2Contained = IsPartiallyContained(pair.Elf2, pair.Elf1)
		if pair.Is1Contained == true {
			containedPairs += 1
			continue
		}
		if pair.Is2Contained == true {
			containedPairs += 1
			continue
		}
	}

	fmt.Print("Part1: ", containedPairs)
	fmt.Println()
}

func SolvePuzzle4() {
	SolvePuzzle4part1()
	SolvePuzzle4part2()
}
