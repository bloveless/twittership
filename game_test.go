package main

import (
	"reflect"
	"testing"
)

func TestGameIsAbleToParseAValidPositionString(t *testing.T) {
	g, err := NewGame("A1H;B8V;E3H;G3V;H8H")
	if err != nil {
		t.Errorf("creating a new game: %v", err)
	}

	expectedShips := []ship{
		{
			x:         0,
			y:         0,
			width:     5,
			direction: horizontal,
			shipType:  shipAircraftCarrier,
		},
		{
			x:         7,
			y:         1,
			width:     4,
			direction: vertical,
			shipType:  shipBattleship,
		},
		{
			x:         2,
			y:         4,
			width:     3,
			direction: horizontal,
			shipType:  shipSubmarine,
		},
		{
			x:         2,
			y:         6,
			width:     3,
			direction: vertical,
			shipType:  shipCruiser,
		},
		{
			x:         7,
			y:         7,
			width:     2,
			direction: horizontal,
			shipType:  shipDestroyer,
		},
	}

	if !reflect.DeepEqual(g.playerShips, expectedShips) {
		t.Fatalf("Player ships did not match expected ships.\nPlayer: %v\nExpected: %v\n", g.playerShips, expectedShips)
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

func TestNewGameWillErrorIfThePositionCannotBeParsed(t *testing.T) {
	t.Parallel()

	for _, invalidPositionString := range invalidPositionStrings {
		t.Run(invalidPositionString.message, func(t *testing.T) {
			_, err := NewGame(invalidPositionString.position)
			if err == nil {
				t.Fatalf("Position string \"%s\" should have failed because of \"%s\"", invalidPositionString.position, invalidPositionString.message)
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

func TestNewGameWillFailPositionsThatCannotBePlacedOnTheBoard(t *testing.T) {
	t.Parallel()

	for _, positionOffTheBoard := range positionsOffTheBoard {
		t.Run(positionOffTheBoard.message, func(t *testing.T) {
			_, err := NewGame(positionOffTheBoard.position)
			if err == nil {
				t.Fatalf("Position string \"%s\" should have failed because \"%s\"", positionOffTheBoard.position, positionOffTheBoard.message)
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

func TestNewGameShouldNotAllowShipsToOverlap(t *testing.T) {
	t.Parallel()

	for _, overlappingPosition := range overlappingPositions {
		t.Run(overlappingPosition.message, func(t *testing.T) {
			_, err := NewGame(overlappingPosition.position)
			if err == nil {
				t.Fatalf("Position string \"%s\" should have failed because \"%s\"", overlappingPosition.position, overlappingPosition.message)
			}
		})
	}
}
