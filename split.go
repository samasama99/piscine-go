package piscine

func Split(s, sep string) []string {
	var result []string
	var current string

	sliceLen := len(sep)
	length := len(s)
	for i := 0; i < length; i++ {
		if (length-i)-sliceLen == 0 {
			current += s[i:]
			break
		} else {
			slice := s[i : i+sliceLen]
			if slice == sep && len(current) > 0 {
				result = append(result, current)
				current = ""
				i += sliceLen
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

// "AB" with "XO" as the separator => i == 0 ; i < 2 - 2
// Split(
//"rHxwB24lAECls|=choumi=|Esymirf4Qq0Ct|=choumi=|rt9Ij6eE2tnFg|=choumi=|MwICbu79WXdLX|=choumi=|2FCoN1yMvMN1W|=choumi=|694NHsXpTYH7R|=choumi=|Sr4cMNZn9KIZn|=choumi=|OQM2rXKEWkmVy",
//"|=choumi=|"
//)
//==
//[]string{"rHxwB24lAECls", "choumi=|Esymirf4Qq0Ct", "choumi=|rt9Ij6eE2tnFg", "choumi=|MwICbu79WXdLX", "choumi=|2FCoN1yMvMN1W", "choumi=|694NHsXpTYH7R", "choumi=|Sr4cMNZn9KIZn", "choumi=|OQM2rXKEWkmVy"}
//instead of
//[]string{"rHxwB24lAECls", "Esymirf4Qq0Ct", "rt9Ij6eE2tnFg", "MwICbu79WXdLX", "2FCoN1yMvMN1W", "694NHsXpTYH7R", "Sr4cMNZn9KIZn", "OQM2rXKEWkmVy"}
