package piscine

func TrimAtoi(s string) int {
	newString := make([]rune, len(s))
	index := 0
	for _, c := range s {
		if isNumeric(c) {
			newString[index] = c
			index++
		}
	}
	newString = newString[:index]
	return Atoi(string(newString))
}
