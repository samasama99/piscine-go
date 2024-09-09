package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := 99
	second := 98

	for {
		if second == 0 {
			first--
			second = first - 1
		}
		if second < 0 {
			return
		}
		printInt(first)
		z01.PrintRune(' ')
		printInt(second)
		if first != 1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		second--
	}
}

var NumbersAsRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func printInt(nbr int) {
	if nbr < 10 {
		z01.PrintRune('0')
		z01.PrintRune(NumbersAsRunes[nbr])
	} else {
		printIntRecursive(nbr)
	}
}

func printIntRecursive(nbr int) {
	if nbr == 0 {
		return
	}
	printIntRecursive(nbr / 10)
	z01.PrintRune(NumbersAsRunes[nbr%10])
}
