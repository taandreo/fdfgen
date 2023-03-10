package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/taandreo/fdfgen/figletlib"
)

var msg string

func main() {
	flag.StringVar(&msg, "msg", "", "Message to be printed on screen")
	flag.Parse()
	font, err := figletlib.GetFontByName("figletlib/fonts/", "banner.flf")
	settings := font.Settings()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// figletlib.SPrintMsg(msg, font, 80, settings, "center")
	text := figletlib.SprintMsg(msg, font, 80, settings, "letf")
	textTofdf(text)
	// fmt.Println(text)
}

func textTofdf(text string) {
	fdfstr := ""
	for _, line := range strings.Split(text, "\n") {
		fdfline := ""
		for _, l := range line {
			if l == ' ' {
				fdfline += "0 0"
			} else if l == '#' {
				fdfline += "10 10"
			}
			fdfline += " "
		}
		fdfline += "\n"
		fdfstr += fdfline + fdfline
	}
	fmt.Println(fdfstr)
}
