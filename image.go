package main

import (
	"image"
	"image/color"
	"image/draw"
	"os"
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

type gameImage struct {
	height     int
	width      int
	tileHeight int
	tileWidth  int
	img        *image.RGBA
}

// GameBoardImage stores the current information for the game image
type GameBoardImage struct {
	fullImage   *image.RGBA
	playerImage gameImage
	enemyImage  gameImage
}

// NewGameImage will create a battleship gameboard with a background. The GameBoardImage returned represents
// both the player image, and the enemy image.
func NewGameImage(h, w int) (GameBoardImage, error) {
	totalW, totalH := w*2+80, h+90
	gbi := GameBoardImage{
		fullImage: image.NewRGBA(image.Rect(0, 0, totalW, totalH)),
		playerImage: gameImage{
			height:     h,
			width:      w,
			tileHeight: h / 10,
			tileWidth:  w / 10,
		},
		enemyImage: gameImage{
			height:     h,
			width:      w,
			tileHeight: h / 10,
			tileWidth:  w / 10,
		},
	}

	f, err := os.Open("game_template.png")
	if err != nil {
		return GameBoardImage{}, err
	}

	gameTemplateImg, _, err := image.Decode(f)
	if err != nil {
		return GameBoardImage{}, err
	}

	draw.Draw(gbi.fullImage, image.Rect(0, 0, totalW, totalH), gameTemplateImg, image.Pt(0, 0), draw.Over)

	gbi.playerImage.img = gbi.fullImage.SubImage(image.Rect(40, 90, w+80, totalH)).(*image.RGBA)
	gbi.playerImage.drawBackground()

	gbi.enemyImage.img = gbi.fullImage.SubImage(image.Rect(w+80, 90, totalW, totalH)).(*image.RGBA)
	gbi.enemyImage.drawBackground()

	return gbi, nil
}

func (gi gameImage) drawBackground() {
	for x := 0; x < gi.width; x++ {
		for y := 0; y < gi.height; y++ {
			if y%gi.tileHeight == 0 || x%gi.tileWidth == 0 {
				c := color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				}
				gi.img.Set(gi.img.Rect.Min.X+x, gi.img.Rect.Min.Y+y, c)
			} else {
				c := color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				}
				gi.img.Set(gi.img.Rect.Min.X+x, gi.img.Rect.Min.Y+y, c)
			}
		}
	}
}

// PlaceAircraftCarrier draws an aircraft carrier on the game board. Size 5
func (gi gameImage) PlaceAircraftCarrier(x, y int, direction string) {
	gi.drawShip(x, y, 5, direction)
}

// PlaceBattleship draws a battleship on the game board. Size 4
func (gi gameImage) PlaceBattleship(x, y int, direction string) {
	gi.drawShip(x, y, 4, direction)
}

// PlaceSubmarine draws a submarine on the game board. Size 3
func (gi gameImage) PlaceSubmarine(x, y int, direction string) {
	gi.drawShip(x, y, 3, direction)
}

// PlaceCruiser draws a cruiser on the game board. Size 3
func (gi gameImage) PlaceCruiser(x, y int, direction string) {
	gi.drawShip(x, y, 3, direction)
}

// PlaceDestroyer draws a destroyer on the game board. Size 2
func (gi gameImage) PlaceDestroyer(x, y int, direction string) {
	gi.drawShip(x, y, 2, direction)
}

func (gi gameImage) drawShip(x, y, width int, direction string) {
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
func (gi gameImage) DrawHit(x, y int) {
	gi.drawVolley(x, y, hit)
}

// DrawMiss draws a miss mark on the game image
func (gi gameImage) DrawMiss(x, y int) {
	gi.drawVolley(x, y, miss)
}

func (gi gameImage) drawVolley(x, y int, volleyType string) {
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

// isPointOnX is used when drawing the X image on the game board to indicate a volley.
// It will determine the the current pixel is anywhere on the X and need to be filled in.
func (gi gameImage) isPointOnX(x, y, xWidth, yWidth int) bool {
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
