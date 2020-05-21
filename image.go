package main

import (
	"image"
	"image/color"
)

const (
	// Horizontal is used to lay a ship horizontally
	Horizontal = "Horizontal"
	// Vertical is used to lay a ship vertically
	Vertical = "Vertical"
	// Hit is used to classify a volley as a hit
	hit = "hit"
	// Miss is used to classify a volley as a miss
	miss = "miss"
)

// GameImage stores the current information for the game image
type GameImage struct {
	height     int
	width      int
	tileHeight int
	tileWidth  int
	img        *image.RGBA
}

// NewImage will create a battleship gameboard with a background
func NewImage(h, w int, img *image.RGBA) GameImage {
	gi := GameImage{
		height:     h,
		width:      w,
		tileHeight: h / 10,
		tileWidth:  w / 10,
		img:        img,
	}

	gi.drawBackground()

	return gi
}

func (gi GameImage) drawBackground() {
	for x := 0; x < gi.width; x++ {
		for y := 0; y < gi.height; y++ {
			if y%gi.tileHeight == 0 || x%gi.tileWidth == 0 {
				c := color.RGBA{
					0,
					0,
					0,
					255,
				}
				gi.img.Set(gi.img.Rect.Min.X+x, gi.img.Rect.Min.Y+y, c)
			} else {
				c := color.RGBA{
					255,
					255,
					255,
					255,
				}
				gi.img.Set(gi.img.Rect.Min.X+x, gi.img.Rect.Min.Y+y, c)
			}
		}
	}
}

// PlaceAircraftCarrier draws an aircraft carrier on the game board. Size 5
func (gi GameImage) PlaceAircraftCarrier(x, y int, direction string) {
	gi.drawShip(x, y, 5, direction)
}

// PlaceBattleship draws a battleship on the game board. Size 4
func (gi GameImage) PlaceBattleship(x, y int, direction string) {
	gi.drawShip(x, y, 4, direction)
}

// PlaceSubmarine draws a submarine on the game board. Size 3
func (gi GameImage) PlaceSubmarine(x, y int, direction string) {
	gi.drawShip(x, y, 3, direction)
}

// PlaceCruiser draws a cruiser on the game board. Size 3
func (gi GameImage) PlaceCruiser(x, y int, direction string) {
	gi.drawShip(x, y, 3, direction)
}

// PlaceDestroyer draws a destroyer on the game board. Size 2
func (gi GameImage) PlaceDestroyer(x, y int, direction string) {
	gi.drawShip(x, y, 2, direction)
}

func (gi GameImage) drawShip(x, y, width int, direction string) {
	startX := x * gi.tileWidth
	startY := y * gi.tileWidth
	endX := startX + width*gi.tileWidth
	endY := startY + 1*gi.tileHeight

	if direction == Vertical {
		endX = startX + 1*gi.tileWidth
		endY = startY + width*gi.tileHeight
	}

	for x := 0; x < endX-startX; x++ {
		for y := 0; y < endY-startY; y++ {
			if y%gi.tileHeight != 0 && x%gi.tileWidth != 0 {
				gi.img.Set(gi.img.Rect.Min.X+startX+x, gi.img.Rect.Min.Y+startY+y, color.RGBA{
					R: 49,
					G: 83,
					B: 123,
					A: 255,
				})
			}
		}
	}
}

// DrawHit draws a hit mark on the game image
func (gi GameImage) DrawHit(x, y int) {
	gi.drawVolley(x, y, hit)
}

// DrawMiss draws a miss mark on the game image
func (gi GameImage) DrawMiss(x, y int) {
	gi.drawVolley(x, y, miss)
}

func (gi GameImage) drawVolley(x, y int, volleyType string) {
	startX := x * gi.tileWidth
	startY := y * gi.tileWidth
	endX := startX + gi.tileWidth
	endY := startY + gi.tileHeight

	bgColor := color.RGBA{
		R: 200,
		G: 200,
		B: 200,
		A: 255,
	}
	if volleyType == hit {
		bgColor = color.RGBA{
			R: 255,
			G: 54,
			B: 51,
			A: 255,
		}
	}

	xWidth := endX - startX
	yWidth := endY - startY

	for x := 0; x < xWidth; x++ {
		for y := 0; y < yWidth; y++ {
			if gi.isPointOnX(x, y, xWidth, yWidth) {
				gi.img.Set(gi.img.Rect.Min.X+startX+x, gi.img.Rect.Min.Y+startY+y, color.RGBA{
					R: 100,
					G: 100,
					B: 100,
					A: 255,
				})

				continue
			}

			if y%gi.tileHeight != 0 && x%gi.tileWidth != 0 {
				gi.img.Set(gi.img.Rect.Min.X+startX+x, gi.img.Rect.Min.Y+startY+y, bgColor)
			}
		}
	}
}

func (gi GameImage) isPointOnX(x, y, xWidth, yWidth int) bool {
	thickness := 3

	// padding around the X axis on the X
	if x < 5 || x > xWidth-5 {
		return false
	}

	// padding around the Y axis on the Y
	if y < 5 || y > yWidth-5 {
		return false
	}

	// Is the point on the X
	if (x > y-thickness && x < y+thickness) || (x > yWidth-y-thickness && x < yWidth-y+thickness) {
		return true
	}

	return false
}
