package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/functions"
	"ascii-art/readfile"
)

func main() {
	if len(os.Args) > 2 { // ckecks if the there are three agguements, if not, the programe prints nothing
		return
	}
	args := os.Args[1]

	text := functions.Text(args)
	
	// asigning a variable, asciiArtFile that's going to store the value of the banner files
	asciiArtFile := "standard.txt"

	

	// the variable inputArgs splits the string(arguement) into substrings, by a separator "\n"
	inputArgs := strings.Split(text, "\\n")
	bannerFileContents := readfile.ReadFile(asciiArtFile)

	output := functions.Graphic(inputArgs, bannerFileContents)
	fmt.Print(output)
}