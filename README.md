<h><b> ASCII ART </h></b>
* ASCII-art is a program written in Go that takes a string as input and outputs a graphic representation of the string using ASCII characters. This program supports numbers, letters, space, special characters, and newline characters. 

## Features ##
* Input Handling :-  Accepts strings containing numbers, letters, spaces, special characters, and newline characters.
* ASCII Representation:- Converts the input string into a graphical representation using ASCII characters.
* Customizable Templates:- Provides different graphical templates such as "shadow," "standard," and "thinkertoy" for ASCII representation.
* Command-Line Interface-: Accessible via the command line for easy input and output.
* Error Handling:- Implements robust error handling to ensure reliable performance.

## Banner Format
* Each character in the ASCII representation has a height of 8 lines. 
* Characters are separated by a newline character \n. 

The program includes several banner formats for graphical representation using ASCII characters:

* Shadow
* Standard
* Thinkertoy
* The program automatically takes the banner format of Standard, but if the user wants to using other graphical representation, the banner format desired should be the third arguement. 

## Usage
```bash
go run . "Your Input String"
```
### Running Tests
To run unit tests, navigate to the project directory and run:
```bash
go test -v
```
