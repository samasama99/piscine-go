package piscine

func IsLower(s string) bool {
	if AlphaCount(s) != len(s) {
		return false
	}
	for _, value := range s {
		if value <= 'a' || value >= 'z' {
			return false
		}
	}
	return true
}
