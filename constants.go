package main

type shipDirection int

const (
	horizontal shipDirection = iota
	vertical
)

type volleyType int

const (
	hit volleyType = iota
	miss
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

const (
	aircraftCarrierWidth = 5
	battleshipWidth      = 4
	submarineWidth       = 3
	cruiserWidth         = 3
	destroyerWidth       = 2
)

func getShipWidth(shipType shipType) int {
	if shipType == shipAircraftCarrier {
		return aircraftCarrierWidth
	}

	if shipType == shipBattleship {
		return battleshipWidth
	}

	if shipType == shipSubmarine {
		return submarineWidth
	}

	if shipType == shipCruiser {
		return cruiserWidth
	}

	if shipType == shipDestroyer {
		return destroyerWidth
	}

	return 0
}
