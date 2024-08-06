package main

import (
	"fmt"
	"github.com/ksrof/gocolors"
)

func main() {
	fmt.Println(ColorizeRed("Съеште еще этих мягких французских булок да выпейте чаю"))
	fmt.Println(ColorizeGreen("Съеште еще этих мягких французских булок да выпейте чаю"))
	fmt.Println(ColorizeBlue("Съеште еще этих мягких французских булок да выпейте чаю"))
	fmt.Println(ColorizeYellow("Съеште еще этих мягких французских булок да выпейте чаю"))
	fmt.Println(ColorizeMagenta("Съеште еще этих мягких французских булок да выпейте чаю"))
	fmt.Println(ColorizeCyan("Съеште еще этих мягких французских булок да выпейте чаю"))
	fmt.Println(ColorizeWhite("Съеште еще этих мягких французских булок да выпейте чаю"))
	fmt.Println(ColorizeCustom("Съеште еще этих мягких французских булок да выпейте чаю", 100, 200, 50))

}

func ColorizeRed(a string) string {
	return gocolors.Red(a, "")
}
func ColorizeGreen(a string) string {
	return gocolors.Green(a, "")
}

func ColorizeBlue(a string) string {
	return gocolors.Blue(a, "")
}
func ColorizeYellow(a string) string {
	return gocolors.Yellow(a, "")
}
func ColorizeMagenta(a string) string {
	return gocolors.Magenta(a, "")
}
func ColorizeCyan(a string) string {
	return gocolors.Cyan(a, "")
}
func ColorizeWhite(a string) string {
	return gocolors.White(a, "")
}

func ColorizeCustom(a string, r, g, b uint8) string {
	color := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	return color + a + "\033[0m"
}
