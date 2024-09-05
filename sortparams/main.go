package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	strings := os.Args[1:]
	chars := make([]byte, len(strings))
	for i, s := range strings {
		chars[i] = s[0]
	}
	printInOrder(chars)
}

func printInOrder(chars []byte) {
	if len(chars) == 0 {
		return
	}

	histogram := make([]int, 128)
	for _, c := range chars {
		histogram[c]++
	}

	for key, v := range histogram {
		for i := 0; i < v; i++ {
			z01.PrintRune(rune(key))
			z01.PrintRune('\n')
		}
	}
}
