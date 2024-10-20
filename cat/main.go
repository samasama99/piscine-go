package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		for {
			bytes := make([]byte, 1024)
			os.Stdin.Read(bytes)
			printStr(string(bytes))
		}
	}
	files := os.Args[1:]
	for _, filename := range files {
		content, err := os.ReadFile(filename)
		if err != nil {
			printStr("ERROR: ")
			printStr(err.Error())
			printStr("\n")
			os.Exit(1)
		}
		printStr(string(content))
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
