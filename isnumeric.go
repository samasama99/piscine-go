package piscine

func numericCount(s string) int {
	runes := []rune(s)
	count := 0
	for _, r := range runes {
		if isNumeric(r) {
			count++
		}
	}
	return count
}

func IsNumeric(s string) bool {
	return numericCount(s) == len(s)
}
