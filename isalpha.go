package piscine

func isNumeric(c rune) bool {
	return c >= '0' && c <= '9'
}

func alphaNumericCount(s string) int {
	runes := []rune(s)
	count := 0
	for _, r := range runes {
		if isAlpha(r) || isNumeric(r) {
			count++
		}
	}
	return count
}

func IsAlpha(s string) bool {
	return alphaNumericCount(s) == len(s)
}
