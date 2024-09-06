package piscine

func MakeRange(min, max int) []int {
	length := max - min
	if length < 1 {
		return []int(nil)
	}
	arr := make([]int, length)
	index := 0
	for current := min; current < max; current++ {
		arr[index] = current
		index++
	}
	return arr
}
