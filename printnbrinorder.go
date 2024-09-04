package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	histogram := make(map[int]int)
	tmp := n
	for tmp != 0 {
		histogram[tmp%10]++
		tmp = tmp / 10
	}

	for key, v := range histogram {
		for i := 0; i < v; i++ {
			z01.PrintRune(rune(key + 48))
		}
	}
}
