package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	f, s, t := 0, 1, 2
	for {
		z01.PrintRune(rune(f + '0'))
		z01.PrintRune(rune(s + '0'))
		z01.PrintRune(rune(t + '0'))
		if f != 7 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}

		if t == 9 && s != 8 {
			s++
			t = s + 1
		} else if s != 8 {
			t++
		} else if f != 7 {
			f++
			s = f + 1
		} else {
			break
		}
	}
}