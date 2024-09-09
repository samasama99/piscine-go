package main

import (
	"github.com/01-edu/z01"
	"os"
)

var NumbersAsRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	op := args[1]
	n1 := Atoi(args[0])
	n2 := Atoi(args[2])
	switch op {
	case "+":
		sum, overflow := add(n1, n2)
		if overflow {
			return
		}
		printInt(sum)
	case "-":
		sum, overflow := subtract(n1, n2)
		if overflow {
			return
		}
		printInt(sum)
	case "*":
		sum, overflow := multiply(n1, n2)
		if overflow {
			return
		}
		printInt(sum)
	case "/":
		if n2 == 0 {
			printStr("No division by 0")
			return
		}
		sum, overflow := divide(n1, n2)
		if overflow {
			return
		}
		printInt(sum)
	case "%":
		if n2 == 0 {
			printStr("No modulo by 0")
			return
		}
		sum, overflow := modulo(n1, n2)
		if overflow {
			return
		}
		printInt(sum)
	}

}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s == "-9223372036854775808" {
		return -9223372036854775808
	}

	if s[0] == '-' {
		return BasicAtoi2(s[1:]) * -1
	} else if s[0] == '+' {
		return BasicAtoi2(s[1:])
	}
	return BasicAtoi2(s)
}

func isValidNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func BasicAtoi2(s string) int {
	if !isValidNumber(s) {
		return 0
	}
	return BasicAtoi(s)
}

func pow(num, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res *= num
	}
	return res
}

func BasicAtoi(s string) int {
	res := 0
	length := len(s)
	for i, value := range s {
		res += (int(value) - 48) * pow(10, length-i-1)
	}
	return res
}

func add(a, b int) (sum int, overflow bool) {
	sum = a + b
	bigSum := int64(a) + int64(b)
	if int64(sum) != bigSum {
		overflow = true
		return
	}
	return
}

func multiply(a, b int) (sum int, overflow bool) {
	sum = a * b
	bigSum := int64(a) * int64(b)
	if int64(sum) != bigSum {
		overflow = true
		return
	}
	return
}

func subtract(a, b int) (sum int, overflow bool) {
	sum = a - b
	bigSum := int64(a) - int64(b)
	if int64(sum) != bigSum {
		overflow = true
		return
	}
	return
}

func divide(a, b int) (sum int, overflow bool) {
	sum = a / b
	bigSum := int64(a) / int64(b)
	if int64(sum) != bigSum {
		overflow = true
		return
	}
	return
}

func modulo(a, b int) (sum int, overflow bool) {
	sum = a % b
	bigSum := int64(a) % int64(b)
	if int64(sum) != bigSum {
		overflow = true
		return
	}
	return
}

func printInt(nbr int) {
	if nbr == 0 {
		z01.PrintRune(48)
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

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
