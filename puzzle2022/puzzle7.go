package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strconv"
	"strings"
)

var (
	puzzleInput = utils.LoadFile("puzzle2022/puzzle7.txt")
)

type Node struct {
	Name     string
	IsDir    bool
	Size     int
	Parent   *Node
	Children []*Node
}

func isCommandLine(row string) bool {
	rowChars := []rune(row)
	if string(rowChars[0]) == "$" {
		return true
	}
	return false
}

func CreateNode(parent Node, row string) {
	rowParted := strings.Split(row, " ")
	if rowParted[0] == "dir" {
		newFolder := &Node{Name: rowParted[1], IsDir: true, Size: 0, Parent: &parent}
		fmt.Println("Created node", newFolder)
		parent.Children = append(parent.Children, newFolder)
	} else {
		intSize, err := strconv.Atoi(rowParted[0])
		if err != nil {
			panic("Error converting size string to int")
		}
		newFile := &Node{Name: rowParted[1], IsDir: false, Size: intSize, Parent: &parent}
		fmt.Println("Created node", newFile)
		parent.Children = append(parent.Children, newFile)
	}

}

func ListFiles(parent Node, idx int) {
	fmt.Println("Listing files:")
	idx += 1 // We make sure to go to the first file
	for true {
		if idx == len(puzzleInput) {
			break
		}
		if isCommandLine(puzzleInput[idx]) || idx == len(puzzleInput) {
			break
		}
		fmt.Println(puzzleInput[idx])
		CreateNode(parent, puzzleInput[idx])
		idx += 1
	}

}

func SolvePuzzle7() {
	var currentNode Node
	for idx, row := range puzzleInput {
		if !isCommandLine(row) {
			continue
		}
		fmt.Println(row)
		if row == "$ cd /" {
			currentNode := &Node{
				Name:     "root",
				IsDir:    true,
				Size:     0,
				Parent:   nil,
				Children: []*Node{}}
			fmt.Println("Created node:", currentNode)
		} else if row == "$ ls" {
			ListFiles(currentNode, idx)
		} else if strings.Contains(row, "$ cd") {
			rowSplit := strings.Split(row, " ")
			if rowSplit[2] == ".." {
				fmt.Println("Change Dir up")
			} else {
				fmt.Println("Change Dir to", rowSplit[2])
			}
		}
	}
}
