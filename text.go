package twittership

import "fmt"

func blue(message string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", message)
}

func blueBg(message string) string {
	return fmt.Sprintf("\x1b[44m%s\x1b[0m", message)
}

func red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

func redBg(message string) string {
	return fmt.Sprintf("\x1b[41m%s\x1b[0m", message)
}

func grey(message string) string {
	return fmt.Sprintf("\x1b[90m%s\x1b[0m", message)
}

func greyBg(message string) string {
	return fmt.Sprintf("\x1b[99m%s\x1b[0m", message)
}

func printTile(tile boardTile) {

}

func PrintGameTextFromGame() {
	fmt.Println("|-----------------------------------------------------------------|")
	fmt.Println("|         PLAYER BOARD          | |          ENEMY BOARD          |")
	fmt.Println("|-----------------------------------------------------------------|")
	fmt.Println("|-| 1| 2| 3| 4| 5| 6| 7| 8| 9|10| |-| 1| 2| 3| 4| 5| 6| 7| 8| 9|10|")

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

		}

		for j := 0; j < 10; j++ {

		}
	}

	fmt.Println("|A|  |  |  |  |  |  |  |  |  |  | |A|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|B|  |  |  |  |  |  |  |  |  |  | |B|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|C|  |  |  |  |  |  |  |  |  |  | |C|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|D|  |  |  |  |  |  |  |  |  |  | |D|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|E|  |  |  |  |  |  |  |  |  |  | |E|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|F|  |  |  |  |  |  |  |  |  |  | |F|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|G|  |  |  |  |  |  |  |  |  |  | |G|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|H|  |  |  |  |  |  |  |  |  |  | |H|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|I|  |  |  |  |  |  |  |  |  |  | |I|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|J|  |  |  |  |  |  |  |  |  |  | |J|  |  |  |  |  |  |  |  |  |  |")
	fmt.Println("|-----------------------------------------------------------------|")

	fmt.Printf("|%s|\n", blueBg(grey("X")))
	fmt.Printf("|%s|\n", redBg(grey("X")))
}
