package piscine

func StringToIntSlice(str string) []int {
	runes := []rune(str)
	IntSlice := make([]int, len(runes))
	for i, r := range runes {
		IntSlice[i] = int(r)
	}
	return IntSlice
}
