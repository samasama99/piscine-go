package piscine

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
