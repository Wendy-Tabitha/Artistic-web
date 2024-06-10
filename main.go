package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/functions"
)

func main() {
	if len(os.Args) != 3{ // ckecks if the there are three agguements, if not, the programe prints nothing
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	args := os.Args[1]

	args = strings.ReplaceAll(args, "\n", "\\n") // replaces all the new lines in the arguement with a new line

	if args == "\\n" { // if  the arguement is "\n" the program prints a new line
		fmt.Println()
		return
	}

	if args == "" {
		return
	}
	// checking if the runes of the string in the arguement is not of an ascii decimal value of more than 126
	for _, chr := range args {
		if chr > '~' {
			fmt.Println("Error : Non Ascii character found!! can not display the graphic representation")
			return
			// checks if the arguement is an empty string , and prints nothing incase the condition is met
		}
	}
	// asigning a variable, asciiArtFile that's going to store the value of the banner files
	asciiArtFile := os.Args[2]

	// if len(os.Args) == 3 {
	// 	args2 := os.Args[2]
	// fmt.Println(args2)
	switch asciiArtFile {
	case "standard":
		asciiArtFile = "standard.txt"
	case "thinkertoy":
		asciiArtFile = "thinkertoy.txt"
	case "shadow":
		asciiArtFile = "shadow.txt"
	default:
		if !(asciiArtFile == "standard" || asciiArtFile == "thinkertoy" || asciiArtFile == "shadow") {
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			os.Exit(0)
		}
	}

	//}

	// the variable inputArgs splits the string(arguement) into substrings, by a separator "\n"
	inputArgs := strings.Split(args, "\\n")
	asciiArt, err := os.ReadFile(asciiArtFile) // reading the bannerfile
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	// the variable asciiArtFile stores the data read from bytes to string
	asciiArtString := string(asciiArt)

	var bannerFileContents []string
	if asciiArtFile == "thinkertoy.txt" {
		bannerFileContents = strings.Split(asciiArtString, "\r\n")
	} else {
		bannerFileContents = strings.Split(asciiArtString, "\n")
	}

	output := functions.Graphic(inputArgs, bannerFileContents)
	fmt.Print(output)
}
