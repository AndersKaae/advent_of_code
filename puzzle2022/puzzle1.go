package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"sort"
	"strconv"
)

func GetCaloriesList() []int {
	var input = utils.LoadFile("puzzle2022/puzzle1.txt")
	var elfCalories = 0
	var totalCaloriesPerElf []int
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			//fmt.Println("")
			totalCaloriesPerElf = append(totalCaloriesPerElf, elfCalories)
			elfCalories = 0
			continue
		}
		//fmt.Println(input[i])
		numberInt, err := strconv.Atoi(input[i])
		if err != nil {
			fmt.Println("Error converting:", err)
		}
		elfCalories += numberInt
	}
	sort.Ints(totalCaloriesPerElf)
	return totalCaloriesPerElf
}

func SolvePuzzle1() {
	caloriesList := GetCaloriesList()
	answer1 := caloriesList[len(caloriesList)-1]
	fmt.Println(answer1)

	answer2 := 0
	for i := 1; i < 4; i++ {
		elf := caloriesList[len(caloriesList)-i]
		answer2 += elf
	}
	fmt.Println(answer2)
}
