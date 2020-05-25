package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	// WriteFile("game_1.png", fullImage)
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
