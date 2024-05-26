package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/functions"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) == 1 { // ckecks if the there are three agguements, if not, the programe prints nothing
		return
	}
	args := os.Args[1]
	
	args = strings.ReplaceAll(args, "\n", "\\n") // replaces all the new lines in the arguement with a new line


	if args == "\\n" { // if  the arguement is "\n" the program prints a new line
		fmt.Println()
		return
	}
// checking if the runes of the string in the arguement is not of an ascii decimal value of more than 126
	for _, chr := range args {
		if chr > 126 {
			fmt.Println("Error : Non Ascii character found!! can not display the graphic representation")
			os.Exit(1)
// checks if the arguement is an empty string , and prints nothing incase the condition is met
		} else if args == "" {
			return
		}
	}
// asigning a variable, asciiArtFile that's going to store the value of the banner files
	asciiArtFile := "standard.txt"

	if len(os.Args) == 3 {
		args2 := os.Args[2]
		switch args2 {
		case "standard":
			asciiArtFile = "standard.txt"
		case "thinkertoy":
			asciiArtFile = "thinkertoy.txt"
		case "shadow":
			asciiArtFile = "shadow.txt"
		default:
			asciiArtFile = "standard.txt"
		}
	}
// the variable inputArgs splits the string(arguement) into substrings, by a separator "\n"
	inputArgs := strings.Split(args, "\\n")
	asciiArt, err := os.ReadFile(asciiArtFile) // reading the bannerfile
	
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	// the variable asciiArtFile stores the data read from bytes to string
	asciiArtString := string(asciiArt)
	if args == asciiArtFile { // if the arguement is a bannerfile (eg: standard / thinkertoy), it should not print it. 
		fmt.Print()
		return
	}

	var bannerFileContents []string
	if asciiArtFile == "thinkertoy.txt" {
		bannerFileContents = strings.Split(asciiArtString, "\r\n")
	} else {
		bannerFileContents = strings.Split(asciiArtString, "\n")
	}

	output := functions.Graphic(inputArgs, bannerFileContents)
	fmt.Print(output)
}
