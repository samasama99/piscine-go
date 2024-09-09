package piscine

func Unmatch(a []int) int {
	history := make(map[int]bool)
	for _, value := range a {
		history[value] = !history[value]
	}
	minUnmatched := int(^(uint(0)) >> 1)
	for key, value := range history {
		if value {
			if minUnmatched > key {
				minUnmatched = key
			}
		}
	}
	if minUnmatched != int(^(uint(0))>>1) {
		return minUnmatched
	}
	return -1
}
