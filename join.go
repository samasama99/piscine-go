package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	lastElem := strs[len(strs)-1]
	return BasicJoin(strs[0:]) + lastElem
}
