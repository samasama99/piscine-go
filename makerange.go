package piscine

func MakeRange(min, max int) []int {
	length := max - min
	arr := make([]int, length)
	index := 0
	if length < 1 {
		return arr
	}
	for i := min; i < max; i++ {
		arr[index] = i
		index++
	}
	return arr
}
