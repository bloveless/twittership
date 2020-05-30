package tests

import (
	"fmt"
	"testing"
	"twittership"
)

var validPositionStrings = []struct {
	name      string
	positions string
	expected  [2][10][10]int
}{
	{
		name:      "basic positions",
		positions: "A1H;B8V;E3H;G3V;H8H",
		expected: [2][10][10]int{
			{
				{0, 0, 0, 0, 0, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, 2, 2, 2, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, 3, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, 3, -1, -1, -1, -1, 4, 4, -1},
				{-1, -1, 3, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			},
			{
				{0, 0, 0, 0, 0, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, 2, 2, 2, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, 3, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, 3, -1, -1, -1, -1, 4, 4, -1},
				{-1, -1, 3, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			},
		},
	},
}

func TestGameIsAbleToLoadShipsFromAValidPositionString(t *testing.T) {
	t.Parallel()

	for _, validPositionString := range validPositionStrings {
		t.Run(validPositionString.name, func(t *testing.T) {
			g := twittership.NewGame()
			err := g.LoadPlayerShips(validPositionString.positions)
			if err != nil {
				t.Errorf("setting player ships: %v", err)
			}

			err = g.LoadEnemyShips(validPositionString.positions)
			if err != nil {
				t.Errorf("setting enemy ships: %v", err)
			}

			equal, output := compareBoardOutput(g.GetShipMap(), validPositionString.expected)
			if !equal {
				t.Fatalf("Ships did not match expected\n%s", output)
			}
		})
	}
}

var invalidPositionStrings = []struct {
	position string
	message  string
}{
	{
		"a1H;B8V;E3H;G3V;H8H",
		"invalid letter in aircraft carrier",
	},
	{
		"A1H;z8V;E3H;G3V;H8H",
		"invalid letter in battleship",
	},
	{
		"A1H;B8V;Y3H;G3V;H8H",
		"invalid letter in submarine",
	},
	{
		"A1H;B8V;E3H;K3V;H8H",
		"invalid letter in cruiser",
	},
	{
		"A1H;B8V;E3H;G3V;Z8H",
		"invalid letter in destroyer",
	},
	{
		"A0H;B8V;E3H;G3V;H8H",
		"invalid number in aircraft carrier",
	},
	{
		"A1H;B11V;E3H;G3V;H8H",
		"invalid number in battleship",
	},
	{
		"A1H;B8V;E0H;G3V;H8H",
		"invalid number in submarine",
	},
	{
		"A1H;B8V;E3H;G13V;H8H",
		"invalid number in cruiser",
	},
	{
		"A1H;B8V;E3H;G3V;H0H",
		"invalid number in destroyer",
	},
	{
		"A1Z;B8V;E3H;G3V;H8H",
		"invalid direction in aircraft carrier",
	},
	{
		"A1H;B8Z;E3H;G3V;H8H",
		"invalid direction in battleship",
	},
	{
		"A1H;B8V;E3Z;G3V;H8H",
		"invalid direction in submarine",
	},
	{
		"A1H;B8V;E3H;G3Z;H8H",
		"invalid direction in cruiser",
	},
	{
		"A1H;B8V;E3H;G3V;H8Z",
		"invalid direction in destroyer",
	},
}

func TestNewGameWillNotLoadPositionsThatCannotBeParsed(t *testing.T) {
	t.Parallel()

	for _, invalidPositionString := range invalidPositionStrings {
		t.Run(invalidPositionString.message, func(t *testing.T) {
			g := twittership.NewGame()
			err := g.LoadPlayerShips(invalidPositionString.position)
			if err == nil {
				t.Fatalf("Player position string \"%s\" should have failed because of \"%s\"", invalidPositionString.position, invalidPositionString.message)
			}

			err = g.LoadEnemyShips(invalidPositionString.position)
			if err == nil {
				t.Fatalf("Enemy position string \"%s\" should have failed because of \"%s\"", invalidPositionString.position, invalidPositionString.message)
			}
		})
	}
}

var positionsOffTheBoard = []struct {
	position string
	message  string
}{
	{
		"G1V;B8V;E3H;G3V;H8H",
		"the aircraft carrier position extends off the board by one square",
	},
	{
		"H1V;B8V;E3H;G3V;H8H",
		"the aircraft carrier position extends off the board by two squares",
	},
	{
		"I1V;B8V;E3H;G3V;H8H",
		"the aircraft carrier position extends off the board by three squares",
	},
	{
		"J1V;B8V;E3H;G3V;H8H",
		"the aircraft carrier position extends off the board by four squares",
	},
	{
		"A1H;B8V;E3H;G3V;J10H",
		"the destroyer position extends off the board by one square",
	},
}

func TestNewGameWillNotLoadPositionsThatCannotBePlacedOnTheBoard(t *testing.T) {
	t.Parallel()

	for _, positionOffTheBoard := range positionsOffTheBoard {
		t.Run(positionOffTheBoard.message, func(t *testing.T) {
			g := twittership.NewGame()
			err := g.LoadPlayerShips(positionOffTheBoard.position)
			if err == nil {
				t.Fatalf("Player position string \"%s\" should have failed because \"%s\"", positionOffTheBoard.position, positionOffTheBoard.message)
			}

			err = g.LoadEnemyShips(positionOffTheBoard.position)
			if err == nil {
				t.Fatalf("Enemy position string \"%s\" should have failed because \"%s\"", positionOffTheBoard.position, positionOffTheBoard.message)
			}
		})
	}
}

var overlappingPositions = []struct {
	position string
	message  string
}{
	{
		"A1V;B1H;E3H;G3V;H8H",
		"battleship is overlapping aircraft carrier",
	},
	{
		"A1H;A2H;E3H;G3V;H8H",
		"battleship is overlapping aircraft carrier",
	},
}

func TestNewGameWillNotLoadShipsThatOverlap(t *testing.T) {
	t.Parallel()

	for _, overlappingPosition := range overlappingPositions {
		t.Run(overlappingPosition.message, func(t *testing.T) {
			g := twittership.NewGame()
			err := g.LoadPlayerShips(overlappingPosition.position)
			if err == nil {
				t.Fatalf("Player position string \"%s\" should have failed because \"%s\"", overlappingPosition.position, overlappingPosition.message)
			}

			err = g.LoadEnemyShips(overlappingPosition.position)
			if err == nil {
				t.Fatalf("Enemy position string \"%s\" should have failed because \"%s\"", overlappingPosition.position, overlappingPosition.message)
			}
		})
	}
}

var validVolleyStrings = []struct {
	name            string
	shipPositions   string
	volleyPositions string
	expectedVolleys [2][10][10]int
}{
	{
		name:            "normal volleys",
		volleyPositions: "A1;B1;C8",
		shipPositions:   "A1H;B8V;E3H;G3V;H8H",
		expectedVolleys: [2][10][10]int{
			{
				{1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{0, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			},
			{
				{1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{0, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, 1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			},
		},
	},
}

func TestGameLoadsVolleysAsEitherHitsOrMisses(t *testing.T) {
	t.Parallel()

	for _, validVolleyString := range validVolleyStrings {
		t.Run(validVolleyString.name, func(t *testing.T) {
			g := twittership.NewGame()
			err := g.LoadPlayerShips(validVolleyString.shipPositions)
			if err != nil {
				t.Fatalf("setting player ships: %v", err)
			}

			err = g.LoadEnemyShips(validVolleyString.shipPositions)
			if err != nil {
				t.Fatalf("setting enemy ships: %v", err)
			}

			err = g.LoadPlayerVolleys(validVolleyString.volleyPositions)
			if err != nil {
				t.Fatalf("setting player volleys: %v", err)
			}

			err = g.LoadEnemyVolleys(validVolleyString.volleyPositions)
			if err != nil {
				t.Fatalf("setting enemy volleys: %v", err)
			}

			equal, output := compareBoardOutput(g.GetVolleyMap(), validVolleyString.expectedVolleys)
			if !equal {
				t.Fatalf("Volleys did not match expected volleys.\n%s", output)
			}
		})
	}
}

func TestCannotLoadPlayerVolleysBeforeLoadingEnemyShips(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadPlayerVolleys("A1;B1")
	if err == nil {
		t.Fatalf("Should have failed when trying to place player volleys before enemy ships")
	}
}

func TestCannotLoadEnemyVolleysBeforeLoadingPlayerShips(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadEnemyVolleys("A1;B1")
	if err == nil {
		t.Fatalf("Should have failed when trying to place enemy volleys before player ships")
	}
}

/*
func TestPlayerShipsGetUpdatedWithHitCountWhenAVolleyHitsThem(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadPlayerShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		t.Fatalf("updating player ships: %v", err)
	}

	err = g.LoadEnemyVolleys("A1;B1;C8")
	if err != nil {
		t.Fatalf("updating enemy volleys: %v", err)
	}

	if g.playerShips[0].hits != 1 {
		t.Fatalf("Player ship[0] should have received a hit")
	}

	if g.playerShips[1].hits != 1 {
		t.Fatalf("Player ship[1] should have received a hit")
	}
}

func TestEnemyShipsGetUpdatedWithHitCountWhenAVolleyHitsThem(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadEnemyShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		t.Fatalf("updating player ships: %v", err)
	}

	err = g.LoadPlayerVolleys("A1;B1;C8")
	if err != nil {
		t.Fatalf("updating enemy volleys: %v", err)
	}

	if g.enemyShips[0].hits != 1 {
		t.Fatalf("Enemy ship[0] should have received a hit")
	}

	if g.enemyShips[1].hits != 1 {
		t.Fatalf("Enemy ship[1] should have received a hit")
	}
}
*/

func TestGameRespondsWithHitWhenAVolleyHitsAnEnemyShip(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadEnemyShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		t.Fatalf("updating enemy ships: %v", err)
	}

	err = g.LoadPlayerVolleys("A1;A4")
	if err != nil {
		t.Fatalf("updating player volleys: %v", err)
	}

	response, err := g.PlayerVolley("A5")
	if err != nil {
		t.Fatalf("player volley: %v", err)
	}

	if response != "Hit" {
		t.Fatalf("expected response to be \"Hit\" but it was %s", response)
	}
}

func TestGameRespondsWithYouSunkMyWhenAVolleySinksAnEnemyShip(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadEnemyShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		t.Fatalf("load enemy ships: %v", err)
	}

	err = g.LoadPlayerVolleys("A1;A2;A3;A4")
	if err != nil {
		t.Fatalf("load player volleys: %v", err)
	}

	response, err := g.PlayerVolley("A5")
	if err != nil {
		t.Fatalf("player volley: %v", err)
	}

	if response != "You sunk my Aircraft Carrier" {
		t.Fatalf("expected response to be \"You sunk my Aircraft Carrier\" but it was %s", response)
	}
}

func TestGameRespondsWithHitWhenAVolleyHitsAPlayerShip(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadPlayerShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		t.Fatalf("load player ships: %v", err)
	}

	err = g.LoadEnemyVolleys("A1;A4")
	if err != nil {
		t.Fatalf("load enemy volleys: %v", err)
	}

	response, err := g.EnemyVolley("A5")
	if err != nil {
		t.Fatalf("enemy volley: %v", err)
	}

	if response != "Hit" {
		t.Fatalf("expected response to be \"Hit\" but it was %s", response)
	}
}

func TestGameRespondsWithYouSunkMyWhenAVolleySinksAPlayerShip(t *testing.T) {
	g := twittership.NewGame()
	err := g.LoadPlayerShips("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		t.Fatalf("updating player ships: %v", err)
	}

	err = g.LoadEnemyVolleys("A1;A2;A3;A4")
	if err != nil {
		t.Fatalf("updating enemy volleys: %v", err)
	}

	response, err := g.EnemyVolley("A5")
	if err != nil {
		t.Fatalf("enemy volley: %v", err)
	}

	if response != "You sunk my Aircraft Carrier" {
		t.Fatalf("expected response to be \"You sunk my Aircraft Carrier\" but it was %s", response)
	}
}

func redBg(i int) string {
	return fmt.Sprintf("\x1b[41m% 2d\x1b[0m", i)
}

func getDigit(i int, i2 int) (bool, string) {
	if i != i2 {
		return false, redBg(i)
	}

	return true, fmt.Sprintf("% 2d", i)
}

func compareBoardOutput(actual [2][10][10]int, expected [2][10][10]int) (bool, string) {
	output := "Player: \n"
	output += "Actual:\t\t\t\t\t\t\t\tExpected:\n"
	boardsEqual := true

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			equal, digit := getDigit(actual[0][y][x], expected[0][y][x])

			if !equal {
				boardsEqual = false
			}

			output += fmt.Sprintf("%s ", digit)
		}

		output += "\t\t"

		for x := 0; x < 10; x++ {
			equal, digit := getDigit(expected[0][y][x], actual[0][y][x])

			if !equal {
				boardsEqual = false
			}

			output += fmt.Sprintf("%s ", digit)
		}

		output += "\n"
	}

	output += "\n"

	output += "Enemy: \n"
	output += "Actual:\t\t\t\t\t\t\t\tExpected:\n"
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			equal, digit := getDigit(actual[1][y][x], expected[1][y][x])

			if !equal {
				boardsEqual = false
			}

			output += fmt.Sprintf("%s ", digit)
		}

		output += "\t\t"

		for x := 0; x < 10; x++ {
			equal, digit := getDigit(expected[1][y][x], actual[1][y][x])

			if !equal {
				boardsEqual = false
			}

			output += fmt.Sprintf("%s ", digit)
		}

		output += "\n"
	}

	return boardsEqual, output
}
