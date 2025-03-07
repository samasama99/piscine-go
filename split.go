package piscine

func Split(s, sep string) []string {
	var result []string
	var current string

	sliceLen := len(sep)
	length := len(s)
	for i := 0; i < length; i++ {
		if (length-i)-sliceLen == 0 && s[i:] != sep {
			current += s[i:]
			break
		} else {
			slice := s[i : i+sliceLen]
			if slice == sep && len(current) > 0 {
				result = append(result, current)
				current = ""
				i += sliceLen - 1
			} else if slice != sep {
				current += string(s[i])
			}
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}
