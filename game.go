package main

import (
	"fmt"
	"regexp"
	"strings"
)

type shipType int

const (
	shipAircraftCarrier shipType = iota
	shipBattleship
	shipSubmarine
	shipCruiser
	shipDestroyer
)

func (s shipType) String() string {
	return [...]string{"Aircraft Carrier", "Battleship", "Submarine", "Cruiser", "Destroyer"}[s]
}

type ship struct {
	x         int
	y         int
	width     int
	direction string
	shipType  shipType
}

// Game contains all the information necessary to play a game of battleship
type Game struct {
	positionRegex *regexp.Regexp
	playerShips   []ship
}

// NewGame creates a new game from a string of positions. The positions will be in the following format
// pos[0] Aircraft Carrier Position + Direction
// pos[1] Battleship Position + Direction
// pos[2] Submarine Position + Direction
// pos[3] Cruiser Position + Direction
// pos[4] Destroyer Position + Direction
// I.E. A1H;B8V;E3H;G3V;H8H
func NewGame(positions string) (Game, error) {
	g := Game{
		positionRegex: regexp.MustCompile(`([A-J])([0-9]{1,2})([H,V])`),
	}

	pos := strings.Split(positions, ";")

	// Deal with the Aircraft Carrier
	x, y, direction, err := g.parsePosition(pos[0])
	if err != nil {
		return Game{}, err
	}

	g.playerShips = append(g.playerShips, ship{
		x:         x,
		y:         y,
		width:     5,
		direction: direction,
		shipType:  shipAircraftCarrier,
	})

	// Deal with the Battleship
	x, y, direction, err = g.parsePosition(pos[1])
	if err != nil {
		return Game{}, err
	}

	g.playerShips = append(g.playerShips, ship{
		x:         x,
		y:         y,
		width:     4,
		direction: direction,
		shipType:  shipBattleship,
	})

	// Deal with the Submarine
	x, y, direction, err = g.parsePosition(pos[2])
	if err != nil {
		return Game{}, err
	}

	g.playerShips = append(g.playerShips, ship{
		x:         x,
		y:         y,
		width:     3,
		direction: direction,
		shipType:  shipSubmarine,
	})

	// Deal with the Cruiser
	x, y, direction, err = g.parsePosition(pos[3])
	if err != nil {
		return Game{}, err
	}

	g.playerShips = append(g.playerShips, ship{
		x:         x,
		y:         y,
		width:     3,
		direction: direction,
		shipType:  shipCruiser,
	})

	// Deal with Destroyer
	x, y, direction, err = g.parsePosition(pos[4])
	if err != nil {
		return Game{}, err
	}

	g.playerShips = append(g.playerShips, ship{
		x:         x,
		y:         y,
		width:     2,
		direction: direction,
		shipType:  shipDestroyer,
	})

	return g, nil
}

func (g Game) parsePosition(position string) (int, int, string, error) {
	parts := g.positionRegex.FindStringSubmatch(position)
	yPos := 0
	xPos := 0
	direction := Horizontal

	// We do not need defaults in the switches below when checking parts[1] (the letter column) or parts[3]
	// (the direction) because the regular expression only allows valid inputs which means that this check
	// will catch any errors in input for parts[1] (letter column) and parts[3] (direction).
	if parts == nil || len(parts) != 4 {
		return 0, 0, "", fmt.Errorf("unable to parse position string: %v", position)
	}

	switch parts[1] {
	case "A":
		yPos = 0
	case "B":
		yPos = 1
	case "C":
		yPos = 2
	case "D":
		yPos = 3
	case "E":
		yPos = 4
	case "F":
		yPos = 5
	case "G":
		yPos = 6
	case "H":
		yPos = 7
	case "I":
		yPos = 8
	case "J":
		yPos = 9
	}

	switch parts[2] {
	case "1":
		xPos = 0
	case "2":
		xPos = 1
	case "3":
		xPos = 2
	case "4":
		xPos = 3
	case "5":
		xPos = 4
	case "6":
		xPos = 5
	case "7":
		xPos = 6
	case "8":
		xPos = 7
	case "9":
		xPos = 8
	case "10":
		xPos = 9
	default:
		return 0, 0, "", fmt.Errorf("unable to parse position string: %v", position)
	}

	switch parts[3] {
	case "H":
		direction = Horizontal
	case "V":
		direction = Vertical
	}

	return xPos, yPos, direction, nil
}
