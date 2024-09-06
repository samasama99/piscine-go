package piscine

import "github.com/01-edu/z01"

func PrintStrLn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	for _, w := range a {
		PrintStrLn(w)
	}
}
