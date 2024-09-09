package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	limit := 5
	for i := 1; i < 5; i++ {
		for j := 1; j < limit; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		limit--
	}
	return arr[2]
}
