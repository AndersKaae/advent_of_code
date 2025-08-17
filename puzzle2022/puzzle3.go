package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
)

type Backpack struct {
	Group    int
	One      string
	Two      string
	Common   []string
	Priority []int
}

type Group struct {
	Number    int
	Backpacks []string
	Common    []string
	Priority  []int
}

var (
	abc = [52]string{
		"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
	}
)

func GetCummulatedPriorities(backpack *Backpack) {
	for c := range backpack.Common {
		found := false
		for idx, letter := range abc {
			if letter == backpack.Common[c] {
				backpack.Priority = append(backpack.Priority, idx+1)
				found = true
				break
			}
		}
		if found == false {
			panic("No Match Found for: " + backpack.Common[c])
		}
	}
}

func FindCommonItems(backpack Backpack) []string {
	var listOfCommon []string
	backpackTwoRunes := []rune(backpack.Two)
	for _, r1 := range backpack.One {
		found := false
		for i := len(backpackTwoRunes) - 1; i >= 0; i-- {
			if r1 == backpackTwoRunes[i] && found == false { // We only evaluate until the first match
				listOfCommon = append(listOfCommon, string(r1))
				backpackTwoRunes = append(backpackTwoRunes[:i], backpackTwoRunes[i+1:]...) // Removing the matched item from the list
				found = true
			}
		}
	}
	return listOfCommon
}

func CreateBackpackStruct() []Backpack {
	var input = utils.LoadFile("puzzle2022/puzzle3.txt")
	var backpacks []Backpack
	groupNumber := 0
	for idx, bp := range input {
		if (idx)%3 == 0 { // every third line is a group
			groupNumber += 1
		}
		var half int = len(bp) / 2
		backpack := Backpack{Group: groupNumber, One: bp[0:half], Two: bp[half:]}
		if len(backpack.One)+len(backpack.Two) != len(bp) {
			panic("Something went wrong in the division")
		}
		backpacks = append(backpacks, backpack)
	}
	return backpacks
}

func CalcTotalScore(backpacks []Backpack) {
	// It seem that we only calculate each letter once, I assume in this calculation that there is never more than one type of letter identical in each backpack
	totalScore := 0
	for i := range backpacks {
		totalScore += backpacks[i].Priority[0]
	}
	fmt.Println(totalScore)
}

func SolvePuzzle3Part1() {
	backpacks := CreateBackpackStruct()
	for i := range backpacks {
		commonItems := FindCommonItems(backpacks[i])
		backpacks[i].Common = commonItems
		GetCummulatedPriorities(&backpacks[i])
		fmt.Println(backpacks[i])
	}
	CalcTotalScore(backpacks)
}

func CreateGroupStruct(backpacks []Backpack) []Group {
	var groupList []Group
	var totalElfGroups = 100
	for i := 1; i < totalElfGroups+1; i++ {
		elfGroup := Group{Number: i}
		for j := range backpacks {
			if backpacks[j].Group == i {
				backpacksConcat := backpacks[j].One + backpacks[j].Two
				elfGroup.Backpacks = append(elfGroup.Backpacks, backpacksConcat)
			}
		}
		if len(elfGroup.Backpacks) != 3 {
			panic("There are supposed to be 3 elfs in each group")
		}
		groupList = append(groupList, elfGroup)
	}
	if len(groupList) != totalElfGroups {
		panic("There are supposed to be 100 groups")
	}
	return groupList
}

func isLetterInString(letter string, stringToCheck string) bool {
	runes := []rune(stringToCheck)
	for i := 0; i < len(runes); i++ {
		if letter == string(runes[i]) {
			return true
		}
	}
	return false
}

func FindCommonBadge(groupList []Group) {
	for group := range groupList {
		for letter := range abc {
			found := true
			if isLetterInString(abc[letter], groupList[group].Backpacks[0]) == false {
				found = false
			}
			if isLetterInString(abc[letter], groupList[group].Backpacks[1]) == false {
				found = false
			}
			if isLetterInString(abc[letter], groupList[group].Backpacks[2]) == false {
				found = false
			}
			if found == true {
				groupList[group].Common = append(groupList[group].Common, abc[letter])
			}
		}
		if len(groupList[group].Common) != 1 {
			panic("No commmon badge was found or more than one badge was found!")
		}
	}
}

func GetCummulatedBadgePriorities(group *Group) {
	for c := range group.Common {
		found := false
		for idx, letter := range abc {
			if letter == group.Common[c] {
				group.Priority = append(group.Priority, idx+1)
				found = true
				break
			}
		}
		if found == false {
			panic("No Match Found for: " + group.Common[c])
		}
	}
}

func SolvePuzzle3Part2() {
	backpacks := CreateBackpackStruct()
	groupList := CreateGroupStruct(backpacks)
	FindCommonBadge(groupList)
	totalScore := 0
	for group := range groupList {
		GetCummulatedBadgePriorities(&groupList[group])
		fmt.Println(groupList[group])
		totalScore += groupList[group].Priority[0]
	}
	println(totalScore)
}

func SolvePuzzle3() {
	//SolvePuzzle3Part1()
	SolvePuzzle3Part2()
}
