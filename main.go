package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	w, h := 401, 401
	fullImage := image.NewRGBA(image.Rect(0, 0, w*2, h+50))

	f, err := os.Open("game_header.png")
	if err != nil {
		panic(err)
	}

	gameHeaderImg, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	draw.Draw(fullImage, image.Rect(0, 0, 802, 50), gameHeaderImg, image.Pt(0, 0), draw.Over)

	yourGameBoard := fullImage.SubImage(image.Rect(0, 50, w, h+50)).(*image.RGBA)
	theirGameBoard := fullImage.SubImage(image.Rect(w, 50, w*2, h+50)).(*image.RGBA)

	gi := NewImage(w, h, yourGameBoard)
	gi.PlaceAircraftCarrier(0, 0, Horizontal)
	gi.PlaceBattleship(1, 2, Horizontal)
	gi.PlaceSubmarine(2, 4, Horizontal)
	gi.PlaceCruiser(3, 6, Horizontal)
	gi.PlaceDestroyer(4, 8, Horizontal)

	gi.DrawHit(1, 0)
	gi.DrawMiss(1, 1)

	gi2 := NewImage(w, h, theirGameBoard)

	gi2.DrawHit(0, 0)
	gi2.DrawMiss(0, 1)

	WriteFile("game.png", fullImage)
}

// WriteFile will write the current gameboard to the file system
func WriteFile(path string, img image.Image) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, img)
}
