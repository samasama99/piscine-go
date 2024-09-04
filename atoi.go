package piscine

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
