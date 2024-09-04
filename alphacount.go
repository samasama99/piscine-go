package piscine

func isAlpha(c rune) bool {
	return c > 'A' && c < 'Z' || c > 'a' && c < 'z'
}

func AlphaCount(s string) int {
	runes := []rune(s)
	count := 0
	for _, r := range runes {
		if isAlpha(r) {
			count++
		}
	}
	return count
}
