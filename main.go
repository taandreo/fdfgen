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

func main() {
	flag.StringVar(&msg, "msg", "", "Message to be printed on screen")
	flag.StringVar(&fdf, "map", "", "Name of the map file to be saved")
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

	font, err := figletlib.GetFontByName("figletlib/fonts/", "banner.flf")
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
	_, err = io.WriteString(f, textTofdf(text))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func textTofdf(text string) string {
	fdfstr := ""
	for _, line := range strings.Split(text, "\n") {
		fdfline := ""
		for _, l := range line {
			if l == ' ' {
				fdfline += "0 0 0"
			} else if l == '#' {
				fdfline += "3 3 3"
			}
			fdfline += " "
		}
		fdfline += "\n"
		fdfstr += fdfline + fdfline + fdfline
	}
	fdfstr = strings.Trim(fdfstr, "\n")
	return fdfstr
}
