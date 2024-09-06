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
		res = string(rune(tmp%10+48)) + res
		tmp = tmp / 10
	}

	return res
}
