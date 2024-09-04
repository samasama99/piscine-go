package piscine

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
