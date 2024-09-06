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
	printStr(intToString(points.x))
	printStr(", y = ")
	printStr(intToString(points.y))
}

func intToString(n int) string {
	if n == 0 {
		z01.PrintRune(rune(48))
	}
	tmp := n
	res := ""
	for tmp != 0 {
		switch tmp % 10 {
		case 0:
			res = "0" + res
		case 1:
			res = "1" + res
		case 2:
			res = "2" + res
		case 3:
			res = "3" + res
		case 4:
			res = "4" + res
		case 5:
			res = "5" + res
		case 6:
			res = "6" + res
		case 7:
			res = "7" + res
		case 8:
			res = "8" + res
		case 9:
			res = "9" + res
		}
		tmp = tmp / 10
	}

	return res
}
