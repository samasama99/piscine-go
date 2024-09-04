package piscine

func ToLower(s string) string {
	length := len(s)
	newString := make([]rune, length)

	for i, c := range s {
		if isAlpha(c) && c < 'a' {
			newString[i] = c + 32
		} else {
			newString[i] = c
		}
	}

	return string(newString)
}
