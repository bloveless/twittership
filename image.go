package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	_ "image/png"
	"os"
)

type userImage struct {
	height     int
	width      int
	tileHeight int
	tileWidth  int
	img        *image.RGBA
}

// GameImage stores the current information for the game image
type GameImage struct {
	fullImage   *image.RGBA
	playerImage userImage
	enemyImage  userImage
}

// NewGameImageFromGame will create a new game image from a game. There is no validation when converting
// a game image to a game because the validation is assume to have happened when creating the game.
func NewGameImageFromGame(g Game, h, w int) (GameImage, error) {
	gi, err := newGameImage(h, w)
	if err != nil {
		return GameImage{}, fmt.Errorf("unable to create new game image: %w", err)
	}

	for _, playerShip := range g.playerShips {
		switch playerShip.shipType {
		case shipAircraftCarrier:
			gi.playerImage.PlaceAircraftCarrier(playerShip.x, playerShip.y, playerShip.direction)
		case shipBattleship:
			gi.playerImage.PlaceBattleship(playerShip.x, playerShip.y, playerShip.direction)
		case shipSubmarine:
			gi.playerImage.PlaceSubmarine(playerShip.x, playerShip.y, playerShip.direction)
		case shipCruiser:
			gi.playerImage.PlaceCruiser(playerShip.x, playerShip.y, playerShip.direction)
		case shipDestroyer:
			gi.playerImage.PlaceDestroyer(playerShip.x, playerShip.y, playerShip.direction)
		}
	}

	return gi, nil
}

// NewGameImage will create a battleship gameboard with a background. The GameImage returned represents
// both the player image, and the enemy image.
func newGameImage(h, w int) (GameImage, error) {
	totalW, totalH := w*2+80, h+90
	gi := GameImage{
		fullImage: image.NewRGBA(image.Rect(0, 0, totalW, totalH)),
		playerImage: userImage{
			height:     h,
			width:      w,
			tileHeight: h / 10,
			tileWidth:  w / 10,
		},
		enemyImage: userImage{
			height:     h,
			width:      w,
			tileHeight: h / 10,
			tileWidth:  w / 10,
		},
	}

	f, err := os.Open("game_template.png")
	if err != nil {
		return GameImage{}, fmt.Errorf("opening game_template: %w", err)
	}

	gameTemplateImg, _, err := image.Decode(f)
	if err != nil {
		return GameImage{}, fmt.Errorf("decoding game_template: %w", err)
	}

	draw.Draw(gi.fullImage, image.Rect(0, 0, totalW, totalH), gameTemplateImg, image.Pt(0, 0), draw.Over)

	gi.playerImage.img = gi.fullImage.SubImage(image.Rect(40, 90, w+80, totalH)).(*image.RGBA)
	gi.playerImage.drawBackground()

	gi.enemyImage.img = gi.fullImage.SubImage(image.Rect(w+80, 90, totalW, totalH)).(*image.RGBA)
	gi.enemyImage.drawBackground()

	return gi, nil
}

func (ui userImage) drawBackground() {
	for x := 0; x < ui.width; x++ {
		for y := 0; y < ui.height; y++ {
			if y%ui.tileHeight == 0 || x%ui.tileWidth == 0 {
				c := color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 255,
				}
				ui.img.Set(ui.img.Rect.Min.X+x, ui.img.Rect.Min.Y+y, c)
			} else {
				c := color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				}
				ui.img.Set(ui.img.Rect.Min.X+x, ui.img.Rect.Min.Y+y, c)
			}
		}
	}
}

// PlaceAircraftCarrier draws an aircraft carrier on the game board. Size 5
func (ui userImage) PlaceAircraftCarrier(x, y int, direction shipDirection) {
	ui.drawShip(x, y, 5, direction)
}

// PlaceBattleship draws a battleship on the game board. Size 4
func (ui userImage) PlaceBattleship(x, y int, direction shipDirection) {
	ui.drawShip(x, y, 4, direction)
}

// PlaceSubmarine draws a submarine on the game board. Size 3
func (ui userImage) PlaceSubmarine(x, y int, direction shipDirection) {
	ui.drawShip(x, y, 3, direction)
}

// PlaceCruiser draws a cruiser on the game board. Size 3
func (ui userImage) PlaceCruiser(x, y int, direction shipDirection) {
	ui.drawShip(x, y, 3, direction)
}

// PlaceDestroyer draws a destroyer on the game board. Size 2
func (ui userImage) PlaceDestroyer(x, y int, direction shipDirection) {
	ui.drawShip(x, y, 2, direction)
}

func (ui userImage) drawShip(x, y, width int, direction shipDirection) {
	startX := x * ui.tileWidth
	startY := y * ui.tileWidth
	endX := startX + width*ui.tileWidth
	endY := startY + 1*ui.tileHeight

	if direction == vertical {
		endX = startX + 1*ui.tileWidth
		endY = startY + width*ui.tileHeight
	}

	for x := 0; x < endX-startX; x++ {
		for y := 0; y < endY-startY; y++ {
			if y%ui.tileHeight != 0 && x%ui.tileWidth != 0 {
				ui.img.Set(ui.img.Rect.Min.X+startX+x, ui.img.Rect.Min.Y+startY+y, color.RGBA{
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
func (ui userImage) DrawHit(x, y int) {
	ui.drawVolley(x, y, hit)
}

// DrawMiss draws a miss mark on the game image
func (ui userImage) DrawMiss(x, y int) {
	ui.drawVolley(x, y, miss)
}

func (ui userImage) drawVolley(x, y int, volley volleyType) {
	startX := x * ui.tileWidth
	startY := y * ui.tileWidth
	endX := startX + ui.tileWidth
	endY := startY + ui.tileHeight

	bgColor := color.RGBA{
		R: 200,
		G: 200,
		B: 200,
		A: 255,
	}

	if volley == hit {
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
			if ui.isPointOnX(x, y, xWidth, yWidth) {
				ui.img.Set(ui.img.Rect.Min.X+startX+x, ui.img.Rect.Min.Y+startY+y, color.RGBA{
					R: 100,
					G: 100,
					B: 100,
					A: 255,
				})

				continue
			}

			if y%ui.tileHeight != 0 && x%ui.tileWidth != 0 {
				ui.img.Set(ui.img.Rect.Min.X+startX+x, ui.img.Rect.Min.Y+startY+y, bgColor)
			}
		}
	}
}

// isPointOnX is used when drawing the X image on the game board to indicate a volley.
// It will determine the the current pixel is anywhere on the X and need to be filled in.
func (ui userImage) isPointOnX(x, y, xWidth, yWidth int) bool {
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

func (gi GameImage) WriteImage(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("unable to write game board image: %w", err)
	}
	defer f.Close()

	err = png.Encode(f, gi.fullImage)
	if err != nil {
		return fmt.Errorf("unable to encode game board image to png: %w", err)
	}

	return nil
}
