package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(rune(48))
	}

	histogram := make([]int, 10)
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
