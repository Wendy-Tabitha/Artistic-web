package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/functions"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) == 1 {
		return
	}
	args := os.Args[1]
	
	// for _, args := range args1 {
	// 	args = strings.ReplaceAll(args, "\n", "\\n")
	// 	if args == "" {
	// 		return
	// 	}

	if args == "\\n" {
		fmt.Println()
		return
	}

	for _, chr := range args {
		if chr > 126 {
			fmt.Println("Error : Non Ascii character found!! can not display the graphic representation")
			os.Exit(1)

		} else if args == "" {
			return
		}
	}

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

	inputArgs := strings.Split(args, "\\n")
	asciiArt, err := os.ReadFile(asciiArtFile)
	
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	asciiArtString := string(asciiArt)
	if args == asciiArtFile {
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
