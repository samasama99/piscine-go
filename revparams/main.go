package main

import (
	"github.com/01-edu/z01"
	"os"
)

func PrintStrLn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	for i := 0; i < len(args); i++ {
		PrintStrLn(string(args[len(args)-i-1]))
	}
}
