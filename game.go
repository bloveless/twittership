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

type boardTile struct {
	shipIndex   int
	volleyIndex int
}

// Game contains all the information necessary to play a game of battleship
type Game struct {
	shipPositionRegex   *regexp.Regexp
	volleyPositionRegex *regexp.Regexp
	playerShips         []ship
	playerVolleys       []volley
	playerBoard         [10][10]boardTile
	enemyShips          []ship
	enemyVolleys        []volley
	enemyBoard          [10][10]boardTile
}

func newBoard() [10][10]boardTile {
	board := [10][10]boardTile{}
	for i := range board {
		boardRow := [10]boardTile{}
		for j := 0; j < 10; j++ {
			boardRow[j].shipIndex = -1
			boardRow[j].volleyIndex = -1
		}

		board[i] = boardRow
	}

	return board
}

// NewGame creates a new game with default regex's
func NewGame() Game {
	g := Game{
		shipPositionRegex:   regexp.MustCompile(`([A-J])([0-9]{1,2})([H,V])`),
		volleyPositionRegex: regexp.MustCompile(`([A-J])([0-9]{1,2})`),
		playerBoard:         newBoard(),
		enemyBoard:          newBoard(),
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

	g.playerBoard, g.playerShips, err = g.getShipsFromPositions(g.playerBoard, positions)
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

	g.enemyBoard, g.enemyShips, err = g.getShipsFromPositions(g.enemyBoard, positions)
	if err != nil {
		return fmt.Errorf("settings enemy positions: %w", err)
	}

	return nil
}

func (g Game) getShipsFromPositions(board [10][10]boardTile, positions string) ([10][10]boardTile, []ship, error) {
	pos := strings.Split(positions, ";")
	var ships []ship
	var err error

	// Deal with the Aircraft Carrier
	board, ships, err = g.addShip(board, ships, pos[0], shipAircraftCarrier)
	if err != nil {
		return [10][10]boardTile{}, []ship{}, err
	}

	// Deal with the Battleship
	board, ships, err = g.addShip(board, ships, pos[1], shipBattleship)
	if err != nil {
		return [10][10]boardTile{}, []ship{}, err
	}

	// Deal with the Submarine
	board, ships, err = g.addShip(board, ships, pos[2], shipSubmarine)
	if err != nil {
		return [10][10]boardTile{}, []ship{}, err
	}

	// Deal with the Cruiser
	board, ships, err = g.addShip(board, ships, pos[3], shipCruiser)
	if err != nil {
		return [10][10]boardTile{}, []ship{}, err
	}

	// Deal with Destroyer
	board, ships, err = g.addShip(board, ships, pos[4], shipDestroyer)
	if err != nil {
		return [10][10]boardTile{}, []ship{}, err
	}

	return board, ships, nil
}

func (g Game) addShip(board [10][10]boardTile, ships []ship, position string, shipType shipType) ([10][10]boardTile, []ship, error) {
	x, y, direction, err := g.parsePosition(position)
	if err != nil {
		return [10][10]boardTile{}, []ship{}, err
	}

	if direction == horizontal && x+getShipWidth(shipType) > 9 {
		return [10][10]boardTile{}, []ship{}, fmt.Errorf("unable to place ship as ship extends off the board")
	}

	if direction == vertical && y+getShipWidth(shipType) > 9 {
		return [10][10]boardTile{}, []ship{}, fmt.Errorf("unable to place ship as ship extends off the board")
	}

	if direction == vertical {
		for i := 0; i < getShipWidth(shipType); i++ {
			if board[y+i][x].shipIndex != -1 {
				return [10][10]boardTile{}, []ship{}, fmt.Errorf("unable to place ship as ship overlaps another ship")
			}

			board[y+i][x].shipIndex = int(shipType)
		}
	}

	if direction == horizontal {
		for i := 0; i < getShipWidth(shipType); i++ {
			if board[y][x+i].shipIndex != -1 {
				return [10][10]boardTile{}, []ship{}, fmt.Errorf("unable to place ship as ship overlaps another ship")
			}

			board[y][x+i].shipIndex = int(shipType)
		}
	}

	return board, append(ships, ship{
		x:         x,
		y:         y,
		width:     getShipWidth(shipType),
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

	_, g.enemyBoard, g.playerVolleys, g.enemyShips, err = g.updateVolleysFromPositions(g.enemyBoard, g.playerVolleys, g.enemyShips, positions)
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

	_, g.playerBoard, g.enemyVolleys, g.playerShips, err = g.updateVolleysFromPositions(g.playerBoard, g.enemyVolleys, g.playerShips, positions)
	if err != nil {
		return fmt.Errorf("setting enemy volleys: %w", err)
	}

	return nil
}

func (g *Game) PlayerVolley(position string) (string, error) {
	var err error
	var response string

	response, g.enemyBoard, g.playerVolleys, g.enemyShips, err = g.updateVolleysFromPositions(g.enemyBoard, g.playerVolleys, g.enemyShips, position)
	if err != nil {
		return "", fmt.Errorf("update player volleys from positions: %w", err)
	}

	return response, nil
}

func (g *Game) EnemyVolley(position string) (string, error) {
	var err error
	var response string

	response, g.playerBoard, g.enemyVolleys, g.playerShips, err = g.updateVolleysFromPositions(g.playerBoard, g.enemyVolleys, g.playerShips, position)
	if err != nil {
		return "", fmt.Errorf("update enemy volleys from positions: %w", err)
	}

	return response, nil
}

func (g Game) updateVolleysFromPositions(board [10][10]boardTile, volleys []volley, ships []ship, positions string) (string, [10][10]boardTile, []volley, []ship, error) {
	pos := strings.Split(positions, ";")
	response := "Miss"

	for _, volleyPos := range pos {
		parts := g.volleyPositionRegex.FindStringSubmatch(volleyPos)

		if parts == nil || len(parts) != 3 {
			return "", [10][10]boardTile{}, []volley{}, []ship{}, fmt.Errorf("unable to parse volley position string: %s", positions)
		}

		// yPos can't be out of range or invalid due to the regex that is used to get parts[1]
		yPos := int(parts[1][0]) - 'A'

		xPos, err := strconv.Atoi(parts[2])
		if err != nil {
			return "", [10][10]boardTile{}, []volley{}, []ship{}, fmt.Errorf("unable to parse volley x position: %v", err)
		}

		xPos--

		if xPos < 0 || xPos > 9 {
			return "", [10][10]boardTile{}, []volley{}, []ship{}, fmt.Errorf("unable to parse volley x position: value out of range")
		}

		vType := miss
		for i, currentShip := range ships {
			if currentShip.direction == horizontal {
				if xPos >= currentShip.x && xPos <= currentShip.x+getShipWidth(currentShip.shipType) && yPos == currentShip.y {
					ships[i].hits++
					if ships[i].hits == ships[i].width {
						response = fmt.Sprintf("You sunk my %s", shipType(i))
					} else {
						response = "Hit"
					}

					board[yPos][xPos].volleyIndex = len(volleys)
					vType = hit
					break
				}
			}

			if currentShip.direction == vertical {
				if yPos >= currentShip.y && yPos <= currentShip.y+getShipWidth(currentShip.shipType) && xPos == currentShip.x {
					ships[i].hits++
					if ships[i].hits == ships[i].width {
						response = fmt.Sprintf("You sunk my %s", shipType(i))
					} else {
						response = "Hit"
					}

					board[yPos][xPos].volleyIndex = len(volleys)
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

	return response, board, volleys, ships, nil
}
