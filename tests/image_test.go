package tests

import (
	"image"
	"os"
	"testing"
	"twittership"
)

func readTestFile(t *testing.T, filename string) image.Image {
	f, err := os.Open("./assets/" + filename)
	if err != nil {
		t.Errorf("opening file %s: %v", filename, err)
	}

	testFile, _, err := image.Decode(f)
	if err != nil {
		t.Errorf("decoding file: %v", err)
	}

	return testFile
}

var drawImageParams = []struct {
	name            string
	playerPositions string
	playerVolleys   string
	enemyPositions  string
	enemyVolleys    string
	expectedImage   string
}{
	{
		name:            "normal game player and enemy have same positions",
		playerPositions: "A1H;B8V;E3H;G3V;H8H",
		playerVolleys:   "A1;B1",
		enemyPositions:  "A1H;B8V;E3H;G3V;H8H",
		enemyVolleys:    "A2;B2",
		expectedImage:   "game_1.png",
	},
}

// TestGameImageCanDrawPixelPerfectImages generates an image and makes sure that it will generate the same image
// compared pixel by pixel. NOTE: If modifying the expected image make sure to save it without
// an alpha channel otherwise it will generate a 32bit PNG rather than a 24bit PNG, as is expected.
func TestGameImageCanDrawPixelPerfectImages(t *testing.T) {
	t.Parallel()

	for _, drawImageParam := range drawImageParams {
		t.Run(drawImageParam.name, func(t *testing.T) {
			w, h := 401, 401
			game := twittership.NewGame()
			err := game.LoadPlayerShips(drawImageParam.playerPositions)
			if err != nil {
				t.Fatalf("setting player ships: %v", err)
			}

			err = game.LoadEnemyShips(drawImageParam.enemyPositions)
			if err != nil {
				t.Fatalf("setting enemy ships: %v", err)
			}

			err = game.LoadPlayerVolleys(drawImageParam.playerVolleys)
			if err != nil {
				t.Fatalf("setting player volleys: %v", err)
			}

			err = game.LoadEnemyVolleys(drawImageParam.enemyVolleys)
			if err != nil {
				t.Fatalf("setting enemy volleys: %v", err)
			}

			gameImage, err := twittership.NewGameImageFromGame(game, w, h, "../game_template.png")
			if err != nil {
				t.Fatalf("creating new game image: %v", err)
			}

			expectedImage := readTestFile(t, drawImageParam.expectedImage)

			for x := 0; x < 882; x++ {
				for y := 0; y < 491; y++ {
					if expectedImage.At(x, y) != gameImage.GetFullImage().At(x, y) {
						t.Fatalf("Color at x: %d y: %d did not match the expected image color", x, y)
					}
				}
			}
		})
	}
}
