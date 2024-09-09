package piscine

func StringToIntSlice(str string) []int {
	if len(str) == 0 {
		return []int(nil)
	}
	runes := []rune(str)
	IntSlice := make([]int, len(runes))
	for i, r := range runes {
		IntSlice[i] = int(r)
	}
	return IntSlice
}
