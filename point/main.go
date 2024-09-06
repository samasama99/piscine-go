package main

import "github.com/01-edu/z01"

var NUMBERS_AS_RUNES = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

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
	if points.x == 0 {
		z01.PrintRune(48)
	} else {
		printIntRecursive(points.x)
	}
	printStr(", y = ")
	if points.y == 0 {
		z01.PrintRune(48)
	} else {
		printIntRecursive(points.y)
	}
}

func printIntRecursive(nbr int) {
	if nbr == 0 {
		return
	}
	printIntRecursive(nbr / 10)
	z01.PrintRune(NUMBERS_AS_RUNES[nbr%10])
}
