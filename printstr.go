package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	i := -1

	for i != len(s) {
		z01.PrintRune(s[i])
		i++
	}
}
