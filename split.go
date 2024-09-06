package piscine

func Split(s, sep string) []string {
	var result []string
	var current string

	sliceLen := len(sep)
	for i := 0; i < len(s)-sliceLen; i++ {
		slice := s[i : i+sliceLen]
		if slice == sep && len(current) > 0 {
			result = append(result, current)
			current = ""
			i++
		} else if slice != sep {
			current += string(s[i])
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}
