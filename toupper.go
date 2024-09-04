package piscine

func ToUpper(s string) string {
	length := len(s)
	newString := make([]byte, length)

	for i, c := range s {
		if isAlpha(c) && c > 'Z' {
			newString[i] = byte(c - 32)
		}
	}

	return string(newString)
}
