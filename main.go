package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/taandreo/fdfgen/figletlib"
)

var msg string
var fdf string
var bgColor string
var fgColor string
var height int64

func main() {
	flag.StringVar(&msg, "msg", "", "Message to be printed on screen: E.g -msg Hello, World")
	flag.StringVar(&fdf, "map", "", "Name of the map file to be saved: E.g. -map example.fdf")
	flag.StringVar(&bgColor, "bg", "", "The background color in Hexadecimal: E.g. -bg 0xFFFFFF")
	flag.StringVar(&fgColor, "fg", "", "The foreground color in Hexadecimal: E.g. -fg 0xFFFFFF")
	flag.Int64Var(&height, "height", 3, "The height of the map(default value is 3): E.g. -height 5")
	flag.Parse()

	if msg == "" {
		_, err := fmt.Fprintln(os.Stderr, "Error: No message passed to the program")
		if err != nil {
			return
		}
		os.Exit(1)
	}
	if fdf == "" {
		_, err := fmt.Fprintln(os.Stderr, "Error: You must type the name of the map to be saved")
		if err != nil {
			fmt.Println("Error using FPrintln")
			os.Exit(1)
		}
		os.Exit(1)
	}

	font, err := figletlib.GetFontByName("figletlib/fonts", "banner.flf")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	settings := font.Settings()

	// figletlib.SPrintMsg(msg, font, 80, settings, "center")
	text := figletlib.SprintMsg(msg, font, 80, settings, "letf")
	f, err := os.Create(fdf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Close the file when done
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Close error", err)
			os.Exit(1)
		}
	}(f)

	// Write a string to the file
	_, err = io.WriteString(f, textTofdf(text, bgColor, fgColor, height))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func textTofdf(text string, bgColor string, fgColor string, height int64) string {
	fdfStr := ""
	for _, line := range strings.Split(text, "\n") {
		fdfLine := ""
		for _, l := range line {
			if l == ' ' {
				if bgColor != "" {
					fdfLine += fmt.Sprintf("0,%s 0,%s 0,%s", bgColor, bgColor, bgColor)
				} else {
					fdfLine += "0 0 0"
				}
			} else if l == '#' {
				if fgColor != "" {
					fdfLine += fmt.Sprintf("%d,%s %d,%s %d,%s", height, fgColor, height, fgColor, height, fgColor)
				} else {
					fdfLine += fmt.Sprintf("%d %d %d", height, height, height)
				}
			}
			fdfLine += " "
		}
		fdfLine += "\n"
		fdfStr += fdfLine + fdfLine + fdfLine
	}
	fdfStr = strings.Trim(fdfStr, "\n")
	return fdfStr
}
