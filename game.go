package main

import (
	"fmt"
	"regexp"
	"strings"
)

type ship struct {
	x         int
	y         int
	width     int
	direction shipDirection
	shipType  shipType
}

// Game contains all the information necessary to play a game of battleship
type Game struct {
	positionRegex *regexp.Regexp
	playerShips   []ship
	shipWidths    map[shipType]int
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
		shipWidths: map[shipType]int{
			shipAircraftCarrier: 5,
			shipBattleship:      4,
			shipSubmarine:       3,
			shipCruiser:         3,
			shipDestroyer:       2,
		},
	}

	pos := strings.Split(positions, ";")

	var err error

	// Deal with the Aircraft Carrier
	g.playerShips, err = g.addShip(pos[0], shipAircraftCarrier)
	if err != nil {
		return Game{}, err
	}

	// Deal with the Battleship
	g.playerShips, err = g.addShip(pos[1], shipBattleship)
	if err != nil {
		return Game{}, err
	}

	// Deal with the Submarine
	g.playerShips, err = g.addShip(pos[2], shipSubmarine)
	if err != nil {
		return Game{}, err
	}

	// Deal with the Cruiser
	g.playerShips, err = g.addShip(pos[3], shipCruiser)
	if err != nil {
		return Game{}, err
	}

	// Deal with Destroyer
	g.playerShips, err = g.addShip(pos[4], shipDestroyer)
	if err != nil {
		return Game{}, err
	}

	return g, nil
}

func (g Game) addShip(position string, shipType shipType) ([]ship, error) {
	x, y, direction, err := g.parsePosition(position)
	if err != nil {
		return []ship{}, err
	}

	if direction == horizontal && x+g.shipWidths[shipType] > 9 {
		return []ship{}, fmt.Errorf("unable to place ship as ship extends off the board")
	}

	if direction == vertical && y+g.shipWidths[shipType] > 9 {
		return []ship{}, fmt.Errorf("unable to place ship as ship extends off the board")
	}

	for _, playerShip := range g.playerShips {
		if playerShip.direction == vertical {
			if y >= playerShip.y && y <= playerShip.y+playerShip.width && x == playerShip.x {
				return []ship{}, fmt.Errorf("unable to place ship as ship overlaps another ship")
			}
		}

		if playerShip.direction == horizontal {
			if x >= playerShip.x && x <= playerShip.x+playerShip.width && y == playerShip.y {
				return []ship{}, fmt.Errorf("unable to place ship as ship overlaps another ship")
			}
		}
	}

	return append(g.playerShips, ship{
		x:         x,
		y:         y,
		width:     g.shipWidths[shipType],
		direction: direction,
		shipType:  shipType,
	}), nil
}

func (g Game) parsePosition(position string) (int, int, shipDirection, error) {
	parts := g.positionRegex.FindStringSubmatch(position)
	yPos := 0
	xPos := 0
	direction := horizontal

	// We do not need defaults in the switches below when checking parts[1] (the letter column) or parts[3]
	// (the direction) because the regular expression only allows valid inputs which means that this check
	// will catch any errors in input for parts[1] (letter column) and parts[3] (direction).
	if parts == nil || len(parts) != 4 {
		return 0, 0, horizontal, fmt.Errorf("unable to parse position string: %v", position)
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
		return 0, 0, horizontal, fmt.Errorf("unable to parse position string: %v", position)
	}

	switch parts[3] {
	case "H":
		direction = horizontal
	case "V":
		direction = vertical
	}

	return xPos, yPos, direction, nil
}
