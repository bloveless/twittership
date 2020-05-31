package twittership

import (
	"fmt"
)

func blueBg(message string) string {
	return fmt.Sprintf("\x1b[44m%s\x1b[0m", message)
}

func redBg(message string) string {
	return fmt.Sprintf("\x1b[41m%s\x1b[0m", message)
}

func getTileString(tile boardTile, isEnemyBoard bool) string {
	// Ship with no volley
	if tile.shipIndex != -1 && tile.volleyIndex == -1 && !isEnemyBoard {
		return blueBg(" ")
	}

	// Hit volley
	if tile.shipIndex != -1 && tile.volleyIndex != -1 {
		return redBg("X")
	}

	// Missed volley
	if tile.shipIndex == -1 && tile.volleyIndex != -1 {
		return "X"
	}

	// Tile with no ship and no volley
	return " "
}

// GetGameTextFromGame will return the game as text which can be printed for a CLI
// version of twittership.
func GetGameTextFromGame(g Game) string {
	output := "|---------------------------------------------|\n"
	output += "|     PLAYER BOARD    | |     ENEMY BOARD     |\n"
	output += "|---------------------------------------------|\n"
	output += "|-|\u2488|\u2489|\u248A|\u248B|\u248C|\u248D|\u248E|\u248F|\u2490|\u2491| |-|\u2488|\u2489|\u248A|\u248B|\u248C|\u248D|\u248E|\u248F|\u2490|\u2491|\n"

	for y := 0; y < 10; y++ {
		output += fmt.Sprintf("|%s", string('A'+y))

		for x := 0; x < 10; x++ {
			output += fmt.Sprintf("|%s", getTileString(g.playerBoard[y][x], false))
		}

		output += fmt.Sprintf("| |%s", string('A'+y))

		for x := 0; x < 10; x++ {
			output += fmt.Sprintf("|%s", getTileString(g.enemyBoard[y][x], true))
		}

		output += "|\n"
	}

	return output
}
