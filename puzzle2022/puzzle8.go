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
	NorthScore  int
	SouthScore  int
	WestScore   int
	EastScore   int
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
	//puzzleInput = utils.LoadFile("puzzle2022/puzzletext/test8.txt")
	puzzleInput = utils.LoadFile("puzzle2022/puzzletext/puzzle8.txt")

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

func CheckWest(treeExamined *Tree) {
	score := 0
	if treeExamined.X == 0 {
		treeExamined.WestScore = score
	}
	for x := treeExamined.X - 1; x >= 0; x-- {
		//fmt.Println(x)
		treeCompared := FindTree(x, treeExamined.Y)
		//fmt.Println("Orginial:", treeExamined.Height, "New tree:", treeCompared.Height)
		if treeCompared.Height < treeExamined.Height {
			score += 1
		} else {
			score += 1
			break
		}
	}
	treeExamined.WestScore = score
}

func CheckEast(treeExamined *Tree) {
	score := 0
	if treeExamined.X == treeSquareHeight-1 {
		treeExamined.NorthScore = score
		return // TODO if we fail, we should reconsider i 0 is right
	}
	for x := treeExamined.X + 1; x < treeSquareHeight; x++ {
		//fmt.Println(x)
		treeCompared := FindTree(x, treeExamined.Y)
		//fmt.Println("Orginial:", treeExamined.Height, "New tree:", treeCompared.Height)
		if treeCompared.Height < treeExamined.Height {
			score += 1
		} else {
			score += 1
			break
		}
	}
	treeExamined.EastScore = score
}

func CheckNorth(treeExamined *Tree) {
	score := 0
	if treeExamined.Y == treeSquareHeight-1 {
		treeExamined.NorthScore = score
		return // TODO if we fail, we should reconsider i 0 is right
	}
	for y := treeExamined.Y + 1; y < treeSquareHeight; y++ {
		treeCompared := FindTree(treeExamined.X, y)
		if treeCompared.Height < treeExamined.Height {
			score += 1
		} else {
			score += 1
			break
		}
	}
	treeExamined.NorthScore = score
}

func CheckSouth(treeExamined *Tree) {
	score := 0
	if treeExamined.Y == 0 {
		treeExamined.SouthScore = score
		return // TODO if we fail, we should reconsider i 0 is right
	}
	for y := treeExamined.Y - 1; y >= 0; y-- {
		//fmt.Println(y)
		treeCompared := FindTree(treeExamined.X, y)
		//fmt.Println("Orginial:", treeExamined.Height, "New tree:", treeCompared.Height)
		if treeCompared.Height < treeExamined.Height {
			score += 1
		} else {
			score += 1
			break
		}
	}
	treeExamined.SouthScore = score
}

func SolvePuzzle8Part2() {
	// First we Loop over each tree
	for row := range treeSquareHeight {
		for x := range treeSuareWidth {
			// Then we evaluate the score of each tree
			y := treeSquareHeight - (row + 1)
			tree := FindTree(x, y)
			CheckNorth(tree)
			CheckSouth(tree)
			CheckWest(tree)
			CheckEast(tree)
			tree.ScenicScore = tree.NorthScore * tree.WestScore * tree.EastScore * tree.SouthScore
			//fmt.Printf("(X:%d,Y:%d),N: %d, W: %d, E: %d, S: %d\n", tree.X, tree.Y, tree.NorthScore, tree.WestScore, tree.EastScore, tree.SouthScore)
			fmt.Printf("%d", tree.ScenicScore)
		}
		println()
	}
	highestScore := 0
	for _, tree := range treeList {
		if tree.ScenicScore > highestScore {
			highestScore = tree.ScenicScore
		}
	}
	fmt.Println("Highest scenic score:", highestScore)
}

func SolvePuzzle8() {
	SolvePuzzle8Part1()
	SolvePuzzle8Part2()
}
