package main

type shipDirection int

const (
	// horizontal is used to lay a ship horizontally
	horizontal shipDirection = iota
	// vertical is used to lay a ship vertically
	vertical
)

type volleyType int

const (
	// Hit is used to classify a volley as a hit
	hit volleyType = iota
	// Miss is used to classify a volley as a miss
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
