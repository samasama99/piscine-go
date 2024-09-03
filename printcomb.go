package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var f, s, t rune = 0, 1, 2
	for {
		z01.PrintRune(f + 48)
		z01.PrintRune(s + 48)
		z01.PrintRune(t + 48)
		if f != 7 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
      
        if s == 8 && f != 7 {
            f++
            s = f + 1
            t = s + 1
        } else if t == 9 && s != 8 {
			s++
			t = s + 1
		} else if t != 9 {
          t++
        } else if s != 8 {
          s++
        } else if f != 7 {
          f++
        } else {
          break
        }
	}
  z01.PrintRune('\n')
}