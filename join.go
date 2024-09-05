package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	lastElem := strs[len(strs)-1]
	res := ""
	for _, e := range strs {
		res = res + e + sep
	}
	return res + lastElem
}
