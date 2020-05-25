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
			direction: Horizontal,
			shipType:  shipAircraftCarrier,
		},
		{
			x:         7,
			y:         1,
			width:     4,
			direction: Vertical,
			shipType:  shipBattleship,
		},
		{
			x:         2,
			y:         4,
			width:     3,
			direction: Horizontal,
			shipType:  shipSubmarine,
		},
		{
			x:         2,
			y:         6,
			width:     3,
			direction: Vertical,
			shipType:  shipCruiser,
		},
		{
			x:         7,
			y:         7,
			width:     2,
			direction: Horizontal,
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
