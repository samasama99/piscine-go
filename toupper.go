package piscine

func ToUpper(s string) string {
	length := len(s)
	newString := make([]rune, length)

	for i, c := range s {
		if isAlpha(c) && c > 'Z' {
			newString[i] = c - 32
		} else {
			newString[i] = c
		}
	}

	return string(newString)
}
