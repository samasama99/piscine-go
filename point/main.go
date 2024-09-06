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
		z01.PrintRune('0')
	}
	Map := "0123456789"
	tmp := n
	res := ""
	for tmp != 0 {
		switch tmp % 10 {
		case 0:
			res = Map[0:1] + res
		case 1:
			res = Map[1:2] + res
		case 2:
			res = Map[2:3] + res
		case 3:
			res = Map[3:4] + res
		case 4:
			res = Map[4:5] + res
		case 5:
			res = Map[5:6] + res
		case 6:
			res = Map[6:7] + res
		case 7:
			res = Map[7:8] + res
		case 8:
			res = Map[8:9] + res
		}
		tmp = tmp / 10
	}

	return res
}
