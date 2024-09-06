package piscine

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	var current string
	for _, c := range s {
		if isWhiteSpace(c) && len(current) > 0 {
			result = append(result, current)
			current = ""
		} else if isWhiteSpace(c) {
			current += string(c)
		}
	}
	return result
}
