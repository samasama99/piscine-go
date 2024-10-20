package piscine

func AppendRange(min, max int) []int {
	var arr []int
	length := max - min
	if length < 1 {
		return arr
	}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
