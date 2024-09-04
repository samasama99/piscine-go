package piscine

func IsUpper(s string) bool {
	if AlphaCount(s) != len(s) {
		return false
	}
	for _, value := range s {
		if value <= 'A' || value >= 'Z' {
			return false
		}
	}
	return true
}
