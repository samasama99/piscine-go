package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	limit := 5
	for i := 1; i < limit; i++ {
		for j := 1; j < limit; j++ {
			if arr[j-1] > arr[j] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		limit--
	}
	return arr[2]
}
