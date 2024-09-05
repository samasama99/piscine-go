package piscine

func isNumeric(c rune) bool {
	return c >= '0' && c <= '9'
}


func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isNumeric(r)
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isAlpha(c rune) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
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

func Capitalize(s string) string {
	isBeginning := true
	newString := make([]byte, len(s))

	for i, c := range s {
		transform := 0
		if isAlphaNumeric(c) {
			if isBeginning {
				isBeginning = false
				if isAlpha(c) && isLower(c) {
					transform = -1
				}
			} else if isAlpha(c) && isUpper(c) {
				transform = 1
			}
		} else {
			isBeginning = true
		}
		newString[i] = byte(c + int32(transform)*32)
	}
	return string(newString)
}
