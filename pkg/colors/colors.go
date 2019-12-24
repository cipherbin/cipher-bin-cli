package colors

import "fmt"

// Println prints the text in the designated color
func Println(text, color string) {
	switch color {
	case "black":
		fmt.Printf("\033[30m%s\033[0m\n", text)
	case "red":
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case "green":
		fmt.Printf("\033[32m%s\033[0m\n", text)
	case "yellow":
		fmt.Printf("\033[33m%s\033[0m\n", text)
	case "blue":
		fmt.Printf("\033[34m%s\033[0m\n", text)
	case "purple":
		fmt.Printf("\033[35m%s\033[0m\n", text)
	case "cyan":
		fmt.Printf("\033[36m%s\033[0m\n", text)
	case "white":
		fmt.Printf("\033[37m%s\033[0m\n", text)
	default:
		fmt.Println(text)
	}
}
