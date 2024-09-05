package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func main() {
	PrintStr(os.Args[0][2:])
	z01.PrintRune('\n')
}
