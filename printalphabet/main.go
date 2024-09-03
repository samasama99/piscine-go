package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i < 'z'+1; i++ {
		err := z01.PrintRune(i)
		if err != nil {
			return
		}
	}
}
