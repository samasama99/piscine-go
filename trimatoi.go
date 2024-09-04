package piscine

func TrimAtoi(s string) int {
	newString := make([]byte, len(s))
	index := 0
	for i, c := range s {
		if isNumeric(c) {
			newString[index] = c
			index++
		}
	}
	newString = newString[:index]
	return Atoi(string(newString))
}
