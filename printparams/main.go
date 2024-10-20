package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStrLn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1:]
	for _, arg := range args {
		PrintStrLn(arg)
	}
}
