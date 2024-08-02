package functions

import (
	"fmt"
	"os"
	"strings"
)

func Text(text string) string {
	text = strings.ReplaceAll(text, "\n", "\\n")

	if text == "" {
		os.Exit(0)
	} else if text == "\\n" {
		fmt.Println()
		os.Exit(0)
	}

	result := ""
	for _, char := range text {
		if char > '~' {
			fmt.Println("Error: Non-ASCII character found! Cannot display the graphic representation")
			os.Exit(0)
		} else if char >= ' ' {
			result += string(char)
		}
	}

	return result
}

func ArtFile(banner string) string {
	switch banner {
	case "standard":
		return "standard.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	case "shadow":
		return "shadow.txt"
	case "lean":
		return "lean.txt"
	default:
		fmt.Printf("Invalid banner: %q\n", banner)
	}

	return ""
}
