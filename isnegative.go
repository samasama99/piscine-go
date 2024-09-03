package main

import "github.com/01-edu/z01"

func isnegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}
