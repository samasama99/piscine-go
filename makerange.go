package piscine

func MakeRange(min, max int) []int {
	length := max - min
	if length < 1 {
		return []int{}
	}
	arr := make([]int, length)
	index := 0
	for i := min; i < max; i++ {
		arr[index] = i
		index++
	}
	return arr
}
