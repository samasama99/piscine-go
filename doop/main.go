package main

import (
	"os"
)

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

var NumbersAsRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func main() {
	args := os.Args[1:]
	if len(args) != 3 || !isValidNumber(args[0]) || !isValidNumber(args[2]) {
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
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		sum, overflow := divide(n1, n2)
		if overflow {
			return
		}
		printInt(sum)
	case "%":
		if n2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
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
	if s[0] == '-' {
		s = s[1:]
	}
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
	if (b > 0 && a > MaxInt-b) || (b < 0 && a < MinInt-b) {
		overflow = true
		return
	}
	return
}

func multiply(a, b int) (sum int, overflow bool) {
	if a != 0 && (b > MaxInt/a || b < MinInt/a) {
		overflow = true
		return
	}
	sum = a * b
	return
}

func subtract(a, b int) (sum int, overflow bool) {
	sum = a - b
	if (b > 0 && a < MinInt+b) || (b < 0 && a > MaxInt+b) {
		overflow = true
		return
	}
	return
}

func divide(a, b int) (sum int, overflow bool) {
	if a == MinInt && b == -1 {
		overflow = true
		return
	}
	sum = a / b
	return
}

func modulo(a, b int) (sum int, overflow bool) {
	sum = a % b
	return
}

func printInt(nbr int) {
	if nbr < 0 {
		os.Stdout.WriteString(string('-'))
		printIntRecursive(-nbr)
	} else if nbr == 0 {
		os.Stdout.WriteString(string('0'))
	} else {
		printIntRecursive(nbr)
	}
	os.Stdout.WriteString(string('\n'))
}

func printIntRecursive(nbr int) {
	if nbr == 0 {
		return
	}
	printIntRecursive(nbr / 10)
	os.Stdout.WriteString(string(NumbersAsRunes[nbr%10]))
}
