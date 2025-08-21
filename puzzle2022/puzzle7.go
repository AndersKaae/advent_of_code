package puzzle2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AndersKaae/advent_of_code/utils"
)

var (
	puzzleInput = utils.LoadFile("puzzle2022/test7.txt")
	//puzzleInput = utils.LoadFile("puzzle2022/puzzle7.txt")
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

func goToRoot(current *Node) *Node {
	for current.Name != "root" {
		current = ChangeDir(current, "$ CD ..")
	}
	return current
}

func CreateNode(parent *Node, row string) {
	rowParted := strings.Split(row, " ")
	if rowParted[0] == "dir" {
		newFolder := &Node{Name: rowParted[1], IsDir: true, Size: 0, Parent: parent}
		fmt.Println("Created node", newFolder)
		parent.Children = append(parent.Children, newFolder)
	} else {
		intSize, err := strconv.Atoi(rowParted[0])
		if err != nil {
			panic("Error converting size string to int")
		}
		newFile := &Node{Name: rowParted[1], IsDir: false, Size: intSize, Parent: parent}
		fmt.Println("Created node", newFile)
		parent.Children = append(parent.Children, newFile)
	}
}

func ListFiles(parent *Node, idx int) {
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

func ChangeDir(currentNode *Node, row string) *Node {
	folderName := strings.Split(row, " ")[2]
	if folderName == ".." {
		fmt.Println("Change Dir up from", currentNode.Name)
		if currentNode.Name == "root" {
			panic("Already at root, cannot go up")
		}
		return currentNode.Parent
	} else {
		fmt.Println("Change Dir to", folderName)
		for _, child := range currentNode.Children {
			fmt.Println(child)
			if child.Name == folderName {
				return child
			}
		}
		panic("Child not found!")
	}
}

func PrintSimpleTree(root *Node) {
	if root == nil {
		fmt.Println("<empty>")
		return
	}
	printSimple(root, 0)
}

func printSimple(n *Node, depth int) {
	indent := strings.Repeat(" ", depth)

	// Root line as "/"
	if n.Parent == nil {
		fmt.Printf("%s- / (dir, size=%d)\n", indent, n.Size)
	} else if n.IsDir {
		fmt.Printf("%s- %s (dir, size=%d)\n", indent, n.Name, n.Size)
		if n.Size < 100001 {
			fmt.Println(n.Size)
		}
	} else {
		fmt.Printf("%s- %s (file, size=%d)\n", indent, n.Name, n.Size)
	}

	for _, c := range n.Children {
		printSimple(c, depth+1)
	}
}

func ComputeDirSizes(n *Node) int {
	if n == nil {
		return 0
	}
	if !n.IsDir {
		return n.Size
	}
	total := 0
	for _, c := range n.Children {
		total += ComputeDirSizes(c)
	}
	n.Size = total
	return total
}

func SolvePuzzle7() {
	var currentNode *Node
	for idx, row := range puzzleInput {
		if !isCommandLine(row) {
			continue
		}
		fmt.Println(row)
		if row == "$ cd /" {
			currentNode = &Node{
				Name:     "root",
				IsDir:    true,
				Size:     0,
				Parent:   nil,
				Children: []*Node{}}
			fmt.Println("Created node:", currentNode)
		} else if row == "$ ls" {
			ListFiles(currentNode, idx)
		} else if strings.Contains(row, "$ cd") {
			currentNode = ChangeDir(currentNode, row)
		}
	}

	fmt.Println("PRINT TREE")
	currentNode = goToRoot(currentNode)
	ComputeDirSizes(currentNode)
	printSimple(currentNode, 0)
}
