package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"
)

func readTestFile(t *testing.T, filename string) image.Image {
	f, err := os.Open("./test_assets/" + filename)
	if err != nil {
		t.Errorf("opening file %s: %v", filename, err)
	}

	testFile, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("decoding file: %v", err)
	}

	return testFile

}

// writeFile will write the current gameboard to the file system, mainly used for debugging the image tests
func writeFile(path string, img image.Image) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, img)
}

// TestDrawBaseImage generates an image and makes sure that it will generate the same image
// compared pixel by pixel. NOTE: If modifying the expected image make sure to save it without
// an alpha channel otherwise it will generate a 32bit PNG rather than a 24bit PNG, as is expected.
func TestDrawBaseImage(t *testing.T) {
	w, h := 401, 401
	gameBoardImage, err := NewGameImage(w, h)
	if err != nil {
		t.Errorf("creating new game image: %v", err)
	}

	gameBoardImage.playerImage.PlaceAircraftCarrier(0, 0, Horizontal)
	gameBoardImage.playerImage.PlaceBattleship(1, 2, Horizontal)
	gameBoardImage.playerImage.PlaceSubmarine(2, 4, Horizontal)
	gameBoardImage.playerImage.PlaceCruiser(3, 6, Horizontal)
	gameBoardImage.playerImage.PlaceDestroyer(4, 8, Horizontal)

	gameBoardImage.playerImage.DrawHit(1, 0)
	gameBoardImage.playerImage.DrawMiss(1, 1)

	gameBoardImage.enemyImage.DrawHit(0, 0)
	gameBoardImage.enemyImage.DrawMiss(0, 1)

	expectedImage := readTestFile(t, "game_1.png")

	// Uncomment for debugging test
	// writeFile("game.png", gameBoardImage.fullImage)

	for x := 0; x < 882; x++ {
		for y := 0; y < 491; y++ {
			if expectedImage.At(x, y) != gameBoardImage.fullImage.At(x, y) {
				t.Fatalf("Color at x: %d y: %d did not match the expected image color", x, y)
			}
		}
	}
}
