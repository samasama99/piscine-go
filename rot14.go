package piscine

func Rot14(s string) string {
	bytes := make([]rune, len(s))
	copy(bytes, []rune(s))
	for i, b := range bytes {
		if _isAlpha(b) {
			if _isUpper(b) {
				if b+14 > 'Z' {
					bytes[i] = 'A' + (b + 14 - 'Z')
				} else {
					bytes[i] = b + 14
				}
			}
			if _isLower(b) {
				if b+14 > 'z' {
					bytes[i] = 'a' + (b + 14 - 'z')
				} else {
					bytes[i] = b + 14
				}
			}
		}
	}
	return string(bytes)
}

func _isAlpha(c rune) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}
func _isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
func _isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
