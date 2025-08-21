package puzzle2022

import (
	"fmt"

	"github.com/AndersKaae/advent_of_code/utils"
)

type Tree struct {
	X           int
	Y           int
	Height      int
	Visible     bool
	ScenicScore int
}

var (
	treeList         = []Tree{}
	treeSquareHeight = 0
	treeSuareWidth   = 0
)

func PrintColored(tree Tree) {
	if tree.Visible == false {
		fmt.Printf("\033[31m%d\033[0m", tree.Height) // red
	} else {
		fmt.Printf("\033[32m%d\033[0m", tree.Height) // green
	}
}

func CreateTreeStruct() []Tree {
	puzzleInput = utils.LoadFile("puzzle2022/puzzletext/test8.txt")
	//puzzleInput = utils.LoadFile("puzzle2022/puzzletext/puzzle8.txt")

	treeSquareHeight = len(puzzleInput)
	y := treeSquareHeight - 1
	for _, row := range puzzleInput {
		runes := []rune(row)
		if treeSuareWidth == 0 {
			treeSuareWidth = len(runes)
		}
		for idx, item := range runes {
			tree := Tree{X: idx, Y: y, Height: utils.ConvertStringToInt(string(item))}
			treeList = append(treeList, tree)
		}
		y -= 1
	}
	return treeList
}

func FindTree(x int, y int) *Tree {
	for i := range treeList {
		if treeList[i].X == x && treeList[i].Y == y {
			return &treeList[i] // pointer to real element
		}
	}
	panic("Tree not found with those coordinates")
}

func PrintTrees(treeList []Tree, scenic bool) {
	for y := range treeSquareHeight {
		for x := range treeSuareWidth {
			//fmt.Printf("X: %d, Y: %d   ", x, treeSquareHeight-(y+1)) // %d for integers
			tree := FindTree(x, treeSquareHeight-(y+1))
			if scenic == false {
				PrintColored(*tree)
			} else {
				fmt.Printf("%d", tree.ScenicScore)
			}

		}
		println()
	}
}

func CalcVisibleTrees() {
	visibleTrees := 0
	for _, tree := range treeList {
		if tree.Visible {
			visibleTrees += 1
		}
	}
	fmt.Println("Total visible trees", visibleTrees)
}

func SolvePuzzle8Part1() {
	treeList := CreateTreeStruct()

	// Loop over Y axis
	for y := range treeSquareHeight {
		prevTreeHeight := 0
		for x := 0; x < treeSuareWidth; x++ {
			tree := FindTree(x, y)
			if tree.X == 0 {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				prevTreeHeight = tree.Height
			}
		}
		prevTreeHeight = 0
		for x := treeSuareWidth - 1; 0 <= x; x-- {
			tree := FindTree(x, y)
			if tree.X == treeSuareWidth-1 {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				prevTreeHeight = tree.Height
			}
		}
	}

	// Loop over X axis
	for x := range treeSuareWidth {
		prevTreeHeight := 0
		for y := 0; y < treeSquareHeight; y++ {
			tree := FindTree(x, y)
			if tree.Y == 0 {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				prevTreeHeight = tree.Height
			}
		}
		prevTreeHeight = 0
		for y := treeSquareHeight - 1; 0 <= y; y-- {
			tree := FindTree(x, y)
			if tree.Y == treeSquareHeight-1 {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				tree.Visible = true
			}
			if tree.Height > prevTreeHeight {
				prevTreeHeight = tree.Height
			}
		}
	}

	PrintTrees(treeList, false)
	CalcVisibleTrees()
}

func SolvePuzzle8Part2() {
	treeList := CreateTreeStruct()
	PrintTrees(treeList, true)
}

func SolvePuzzle8() {
	SolvePuzzle8Part1()
	SolvePuzzle8Part2()
}
