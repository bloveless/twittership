package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ship struct {
	x         int
	y         int
	width     int
	hits      int
	direction shipDirection
	shipType  shipType
}

type volley struct {
	x          int
	y          int
	volleyType volleyType
}

// Game contains all the information necessary to play a game of battleship
type Game struct {
	shipPositionRegex   *regexp.Regexp
	volleyPositionRegex *regexp.Regexp
	playerShips         []ship
	playerVolleys       []volley
	enemyShips          []ship
	enemyVolleys        []volley
	shipWidths          map[shipType]int
}

// NewGame creates a new game with default shipPositionRegex and shipWidths
func NewGame() Game {
	g := Game{
		shipPositionRegex:   regexp.MustCompile(`([A-J])([0-9]{1,2})([H,V])`),
		volleyPositionRegex: regexp.MustCompile(`([A-J])([0-9]{1,2})`),
		shipWidths: map[shipType]int{
			shipAircraftCarrier: 5,
			shipBattleship:      4,
			shipSubmarine:       3,
			shipCruiser:         3,
			shipDestroyer:       2,
		},
	}

	return g
}

// LoadPlayerShips loads the player ships from a string of positions.
// The positions will be in the following format:
// 	 pos[0] Aircraft Carrier Position + Direction
//   pos[1] Battleship Position + Direction
//   pos[2] Submarine Position + Direction
//   pos[3] Cruiser Position + Direction
//   pos[4] Destroyer Position + Direction
// I.E. A1H;B8V;E3H;G3V;H8H
func (g *Game) LoadPlayerShips(positions string) error {
	var err error

	g.playerShips, err = g.getShipsFromPositions(positions)
	if err != nil {
		return fmt.Errorf("setting player positions: %w", err)
	}

	return nil
}

// LoadEnemyShips updates the enemy ships within the game from a string of positions.
// The positions will be in the following format:
// 	 pos[0] Aircraft Carrier Position + Direction
//   pos[1] Battleship Position + Direction
//   pos[2] Submarine Position + Direction
//   pos[3] Cruiser Position + Direction
//   pos[4] Destroyer Position + Direction
// I.E. A1H;B8V;E3H;G3V;H8H
func (g *Game) LoadEnemyShips(positions string) error {
	var err error

	g.enemyShips, err = g.getShipsFromPositions(positions)
	if err != nil {
		return fmt.Errorf("settings enemy positions: %w", err)
	}

	return nil
}

func (g Game) getShipsFromPositions(positions string) ([]ship, error) {
	pos := strings.Split(positions, ";")
	var ships []ship
	var err error

	// Deal with the Aircraft Carrier
	ships, err = g.addShip(ships, pos[0], shipAircraftCarrier)
	if err != nil {
		return []ship{}, err
	}

	// Deal with the Battleship
	ships, err = g.addShip(ships, pos[1], shipBattleship)
	if err != nil {
		return []ship{}, err
	}

	// Deal with the Submarine
	ships, err = g.addShip(ships, pos[2], shipSubmarine)
	if err != nil {
		return []ship{}, err
	}

	// Deal with the Cruiser
	ships, err = g.addShip(ships, pos[3], shipCruiser)
	if err != nil {
		return []ship{}, err
	}

	// Deal with Destroyer
	ships, err = g.addShip(ships, pos[4], shipDestroyer)
	if err != nil {
		return []ship{}, err
	}

	return ships, nil
}

func (g Game) addShip(ships []ship, position string, shipType shipType) ([]ship, error) {
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

	for _, playerShip := range ships {
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

	return append(ships, ship{
		x:         x,
		y:         y,
		width:     g.shipWidths[shipType],
		direction: direction,
		shipType:  shipType,
	}), nil
}

func (g Game) parsePosition(position string) (int, int, shipDirection, error) {
	parts := g.shipPositionRegex.FindStringSubmatch(position)
	direction := horizontal

	if parts == nil || len(parts) != 4 {
		return 0, 0, horizontal, fmt.Errorf("unable to parse ship position string: %v", position)
	}

	// yPos can't be out of range or invalid due to the regex that is used to get parts[1]
	yPos := int(parts[1][0]) - 'A'

	xPos, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, horizontal, fmt.Errorf("unable to parse ship x position: %v", err)
	}

	xPos--

	if xPos < 0 || xPos > 9 {
		return 0, 0, horizontal, fmt.Errorf("unable to parse ship x position: out of range")
	}

	switch parts[3] {
	case "H":
		direction = horizontal
	case "V":
		direction = vertical
	}

	return xPos, yPos, direction, nil
}

func (g *Game) LoadPlayerVolleys(positions string) error {
	var err error

	if len(g.enemyShips) == 0 {
		return fmt.Errorf("cannot place player volleys before placing enemy ships")
	}

	g.playerVolleys, g.enemyShips, err = g.getVolleysFromPositions(g.enemyShips, positions)
	if err != nil {
		return fmt.Errorf("setting player volleys: %w", err)
	}

	return nil
}

func (g *Game) LoadEnemyVolleys(positions string) error {
	var err error

	if len(g.playerShips) == 0 {
		return fmt.Errorf("cannot place enemy volleys before placing player ships")
	}

	g.enemyVolleys, g.playerShips, err = g.getVolleysFromPositions(g.playerShips, positions)
	if err != nil {
		return fmt.Errorf("setting enemy volleys: %w", err)
	}

	return nil
}

func (g Game) getVolleysFromPositions(ships []ship, positions string) ([]volley, []ship, error) {
	pos := strings.Split(positions, ";")
	var volleys []volley

	for _, volleyPos := range pos {
		parts := g.volleyPositionRegex.FindStringSubmatch(volleyPos)

		if parts == nil || len(parts) != 3 {
			return []volley{}, []ship{}, fmt.Errorf("unable to parse volley position string: %s", positions)
		}

		// yPos can't be out of range or invalid due to the regex that is used to get parts[1]
		yPos := int(parts[1][0]) - 'A'

		xPos, err := strconv.Atoi(parts[2])
		if err != nil {
			return []volley{}, []ship{}, fmt.Errorf("unable to parse volley x position: %v", err)
		}

		xPos--

		if xPos < 0 || xPos > 9 {
			return []volley{}, []ship{}, fmt.Errorf("unable to parse volley x position: value out of range")
		}

		vType := miss
		for i, currentShip := range ships {
			if currentShip.direction == horizontal {
				if xPos >= currentShip.x && xPos <= currentShip.x+g.shipWidths[currentShip.shipType] && yPos == currentShip.y {
					ships[i].hits++
					vType = hit
					break
				}
			}

			if currentShip.direction == vertical {
				if yPos >= currentShip.y && yPos <= currentShip.y+g.shipWidths[currentShip.shipType] && xPos == currentShip.x {
					ships[i].hits++
					vType = hit
					break
				}
			}
		}

		volleys = append(volleys, volley{
			x:          xPos,
			y:          yPos,
			volleyType: vType,
		})
	}

	return volleys, ships, nil
}
