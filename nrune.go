package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	for i, value := range runes {
		if i == n-1 {
			return value
		}
	}
	return 0
}
