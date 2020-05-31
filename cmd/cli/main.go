package main

import (
	"fmt"
	"log"
	"twittership"
)

func main() {
	g := twittership.NewGame()
	err := g.LoadPlayerShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		log.Fatalf("Unable to load player ships: %v", err)
	}

	err = g.LoadEnemyShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		log.Fatalf("Unable to load enemy ships: %v", err)
	}

	err = g.LoadPlayerVolleys("A1;B1;C8")
	if err != nil {
		log.Fatalf("Unable to load player volleys: %s", err)
	}

	err = g.LoadEnemyVolleys("A1;B1;C8")
	if err != nil {
		log.Fatalf("Unable to load enemy volleys: %s", err)
	}

	fmt.Println(twittership.GetGameTextFromGame(g))
}
