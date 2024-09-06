package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)
	printStr("x = ")
	PrintNbr(points.x)
	printStr(", y = ")
	PrintNbr(points.y)
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune(rune(48))
	}

	tmp := n
	for tmp != 0 {
		z01.PrintRune(rune(tmp%10 + 48))
		tmp = tmp / 10
	}
}
