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
		} else if !isWhiteSpace(c) {
			current += string(c)
		}
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}

// SplitWhiteSpaces("+k#BK2hPEf&mt N)i\\fdH%QS)z> /UCHU.t1qww;e [=yo'Cpm#z-kA e!wxQLIS.CzWk b;<\\]kC*x[b:q e7QXt+ihx#~hJ n6`~:`aVsDz  ")
//==
//[]string{"+k#BK2hPEf&mt", "N)i\\fdH%QS)z>", "/UCHU.t1qww;e", "[=yo'Cpm#z-kA", "e!wxQLIS.CzWk", "b;<\\]kC*x[b:q", "e7QXt+ihx#~hJ", "n6`~:`aVsDz", ""}
//instead of
//[]string{"+k#BK2hPEf&mt", "N)i\\fdH%QS)z>", "/UCHU.t1qww;e", "[=yo'Cpm#z-kA", "e!wxQLIS.CzWk", "b;<\\]kC*x[b:q", "e7QXt+ihx#~hJ", "n6`~:`aVsDz"}
