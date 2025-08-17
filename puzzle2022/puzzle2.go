package puzzle2022

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"strconv"
)

type Game struct {
	Oponent string
	Me      string
}

func CreateGameStruct() []Game {
	var input = utils.LoadFile("puzzle2022/puzzle2.txt")
	var games []Game
	for _, g := range input {
		var game = Game{Oponent: g[0:1], Me: g[2:3]}
		games = append(games, game)
	}
	return games
}

func RockPaperScissorsSolver(game Game) string {
	switch {
	case game.Oponent == "A" && game.Me == "X": // Rock Rock
		return "draw"
	case game.Oponent == "B" && game.Me == "Y": // Paper Paper
		return "draw"
	case game.Oponent == "C" && game.Me == "Z": // Scissor Scissor
		return "draw"
	case game.Oponent == "A" && game.Me == "Y": // Rock Paper
		return "win"
	case game.Oponent == "A" && game.Me == "Z": // Rock Scissors
		return "loss"
	case game.Oponent == "B" && game.Me == "X": // Paper Rock
		return "loss"
	case game.Oponent == "B" && game.Me == "Z": // Paper Scissors
		return "win"
	case game.Oponent == "C" && game.Me == "X": // Scissor Rock
		return "win"
	case game.Oponent == "C" && game.Me == "Y": // Scissor Paper
		return "loss"
	}
	panic("We did not account for all game combinations!")
}

func RockPaperScissorsReverseSolver(game *Game) {
	// X: lose
	// Y: draw
	// Z: win

	// My Choices
	// X: Rock (1 point)
	// Y: Paper (2 points)
	// Z: Scissors (3 points)
	switch {
	case game.Oponent == "A" && game.Me == "X": // Rock lose
		game.Me = "Z"
		return
	case game.Oponent == "B" && game.Me == "Y": // Paper draw
		game.Me = "Y"
		return
	case game.Oponent == "C" && game.Me == "Z": // Scissor win
		game.Me = "X"
		return
	case game.Oponent == "A" && game.Me == "Y": // Rock draw
		game.Me = "X"
		return
	case game.Oponent == "A" && game.Me == "Z": // Rock win
		game.Me = "Y"
		return
	case game.Oponent == "B" && game.Me == "X": // Paper lose
		game.Me = "X"
		return
	case game.Oponent == "B" && game.Me == "Z": // Paper win
		game.Me = "Z"
		return
	case game.Oponent == "C" && game.Me == "X": // Scissor lose
		game.Me = "Y"
		return
	case game.Oponent == "C" && game.Me == "Y": // Scissor draw
		game.Me = "Z"
		return
	}
	panic("We did not account for all game combinations: " + game.Oponent + "; " + game.Me)
}

func ScoreGame(g Game) int {
	score := 0
	gameResult := RockPaperScissorsSolver(g)
	if gameResult == "draw" {
		score += 3
	} else if gameResult == "win" {
		score += 6
	} else if gameResult == "loss" {
		score += 0
	} else {
		panic("Invalid game result")
	}

	if g.Me == "X" {
		score += 1
	}

	if g.Me == "Y" {
		score += 2
	}

	if g.Me == "Z" {
		score += 3
	}
	return score
}

func SolvePart1() {
	games := CreateGameStruct()
	scoreTotal := 0
	for _, g := range games {
		score := ScoreGame(g)
		scoreTotal += score
	}
	fmt.Println("Part1: " + strconv.Itoa(scoreTotal))
}

func SolvePart2() {
	games := CreateGameStruct()
	scoreTotal := 0
	for g := range games {
		RockPaperScissorsReverseSolver(&games[g])
	}

	for _, g := range games {
		score := ScoreGame(g)
		scoreTotal += score
	}
	fmt.Println("Part2: " + strconv.Itoa(scoreTotal))
}

func SolvePuzzle2() {
	SolvePart1()
	SolvePart2()
}
