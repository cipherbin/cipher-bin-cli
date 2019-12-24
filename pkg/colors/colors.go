package colors

import "fmt"

// color is a type alias for int
type color int

// Supported Colors use the iota pattern to assign color (int) values to
// the availble color variables
const (
	Black color = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)

// Println prints the text in the designated color. It's second parameter
// is of type color which unexported... This makes it so that the caller has
// to use one of the exported constant color variables (which  are of type
// color). Keeping the color type unexported makes it so that a caller
// can't create a new Color (if it were exported) and assign it to an
// arbitrary int. The default is to just fmt.Println() so it's not exactly
// defensive coding, however it is attempting to guide the caller into using
// the supported colors. Haven't figured out if this is bad practice due to
// the second param asking for an unexported type, or if this is an elegant
// enum type solution?
func Println(text string, c color) {
	switch c {
	case Black:
		fmt.Printf("\033[30m%s\033[0m\n", text)
	case Red:
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case Green:
		fmt.Printf("\033[32m%s\033[0m\n", text)
	case Yellow:
		fmt.Printf("\033[33m%s\033[0m\n", text)
	case Blue:
		fmt.Printf("\033[34m%s\033[0m\n", text)
	case Purple:
		fmt.Printf("\033[35m%s\033[0m\n", text)
	case Cyan:
		fmt.Printf("\033[36m%s\033[0m\n", text)
	case White:
		fmt.Printf("\033[37m%s\033[0m\n", text)
	default:
		fmt.Println(text)
	}
}
