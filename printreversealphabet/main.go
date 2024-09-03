package main

import "github.com/01-edu/z01"

func main() {
	for i := 'z'; i > 'a'-1; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
