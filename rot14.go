package piscine

func Rot14(s string) string {
	bytes := make([]rune, len(s))
	copy(bytes, []rune(s))
	for i, b := range bytes {
		if isAlpha(b) {
			if isUpper(b) {
				if b+14 > 'Z' {
					bytes[i] = 'A' + (b + 14 - 'Z')
				} else {
					bytes[i] = b + 14
				}
			}
			if isLower(b) {
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
