package readfile

import (
	"fmt"
	"os"
	"regexp"
)

func ReadFile(asciiArtFile string) []string {
	asciiArt, err := os.ReadFile(asciiArtFile) // reading the bannerfile
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	if len(asciiArt) == 0 {
		fmt.Println("Empty banner file")
		os.Exit(1)
	}
	// the variable asciiArtFile stores the data read from bytes to string

	asciiArtString := string(asciiArt)

	// Split asciiArtString into lines.
	re := regexp.MustCompile(`\r?\n`)
	bannerFileContents := re.Split(asciiArtString, -1)

	return bannerFileContents
}
