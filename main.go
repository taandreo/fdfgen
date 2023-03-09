package main

import "fmt"
import "figletlib"

var filename = "fonts/banner.flf"

func printletter(font Font, letter rune) {
	for _, line := range font.chars[letter] {
		fmt.Println(string(line))
	}
}

func main() {
	font, err := ReadFont(filename)
	if err != nil {
		fmt.Println(err)
	}
	printletter(*font, 't')
	printletter(*font, 'a')
	printletter(*font, 'i')
}
