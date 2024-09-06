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
	switch nbr % 10 {
	case 1:
		z01.PrintRune(49)
	case 2:
		z01.PrintRune(50)
	case 3:
		z01.PrintRune(51)
	case 4:
		z01.PrintRune(52)
	case 5:
		z01.PrintRune(53)
	case 6:
		z01.PrintRune(54)
	case 7:
		z01.PrintRune(55)
	case 8:
		z01.PrintRune(56)
	}
}
