package puzzle2023

import (
	"fmt"
	"github.com/AndersKaae/advent_of_code/utils"
	"slices"
	"strings"
)

type Games struct {
	GameID int
	Blue   int
	Red    int
	Green  int
}

func populateGames(input []string) []Games {
	var games []Games
	for i := 0; i < len(input); i++ {
		// Split the input line on ":" and get the second part
		parts := strings.Split(input[i], ":")
		if len(parts) < 2 {
			continue // Skip if the format is incorrect
		}
		result := parts[1]

		// Split the result into rolls
		listOfRolls := strings.Split(result, "; ")
		for j := 0; j < len(listOfRolls); j++ {
			game := Games{GameID: i + 1, Blue: 0, Red: 0, Green: 0}
			// 6 red, 1 blue, 3 green
			// Split the roll into the player and the score
			rolls := strings.Split(listOfRolls[j], ",")
			// [6 red, 1 blue, 3 green]
			for k := 0; k < len(rolls); k++ {
				rolls[k] = strings.TrimSpace(rolls[k])
				units := strings.Split(rolls[k], " ")

				score := utils.ConvertStringToInt(units[0])
				color := units[1]

				if color == "red" {
					if score > game.Red {
						game.Red = score
					}
				} else if color == "blue" {
					if score > game.Blue {
						game.Blue = score
					}
				} else if color == "green" {
					if score > game.Green {
						game.Green = score
					}
				}
			}

			// Add the game to the list of games
			games = append(games, game)
		}
	}
	return games
}

func puzzle3(input []string) {

	limit := Games{GameID: 0, Blue: 14, Red: 12, Green: 13}

	illegalIds := []int{}
	legalIds := []int{}

	games := populateGames(input)
	for i := 0; i < len(games); i++ {
		if games[i].Blue > limit.Blue || games[i].Red > limit.Red || games[i].Green > limit.Green {
			illegalIds = append(illegalIds, games[i].GameID)
		}
	}

	for i := 0; i < len(games); i++ {
		if slices.Contains(illegalIds, games[i].GameID) {
		} else {
			legalIds = append(legalIds, games[i].GameID)
		}
	}
	legalIds = slices.Compact(legalIds)
	fmt.Println("Legal games:", legalIds)
	sumLegalIds := 0
	for i := 0; i < len(legalIds); i++ {
		sumLegalIds += legalIds[i]
	}
	fmt.Println("Sum of legal games:", sumLegalIds)
}

func puzzle4(input []string) {
	games := populateGames(input)
	gameId := 1

	game := Games{GameID: gameId, Blue: 0, Red: 0, Green: 0}

	maxGames := []Games{}

	for i := 0; i < len(games); i++ {
		if games[i].GameID == gameId {
			if games[i].Blue > game.Blue {
				game.Blue = games[i].Blue
			}
			if games[i].Red > game.Red {
				game.Red = games[i].Red
			}
			if games[i].Green > game.Green {
				game.Green = games[i].Green
			}
		} else {
			maxGames = append(maxGames, game)
			gameId++
			game = Games{GameID: gameId, Blue: 0, Red: 0, Green: 0}
			i--
		}
	}
	maxGames = append(maxGames, game)
	fmt.Println(maxGames)

	power := 0
	for i := 0; i < len(maxGames); i++ {
		power += maxGames[i].Blue * maxGames[i].Red * maxGames[i].Green
	}
	fmt.Println("Power:", power)
}

func SolvePuzzle2() {
	input := utils.LoadFile("puzzle2023/puzzle2.txt")
	//puzzle3(input)
	puzzle4(input)
}
