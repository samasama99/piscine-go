package main

import (
	"github.com/01-edu/z01"
	"os"
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
			continue
		}
		printStr(string(content))
		println()
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
