package piscine

func StringToIntSlice(str string) []int {
	IntSlice := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		IntSlice[i] = int(str[i])
	}
	return IntSlice
}
