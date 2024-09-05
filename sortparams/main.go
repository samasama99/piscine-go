package main

import (
	"os"

	"github.com/01-edu/z01"
)

type Transformer[T any, U any] func(T) U

func transform[T any, U any](collection []T, transformer Transformer[T, U]) []U {
	newCollection := make([]U, len(collection))
	for i, value := range collection {
		newCollection[i] = transformer(value)
	}
	return newCollection
}

func main() {
	strings := os.Args[1:]
	chars := transform(strings, func(v string) byte {
		return v[0]
	})
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
			z01.PrintRune(rune('\n'))
		}
	}
}
