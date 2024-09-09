package piscine

func Unmatch(a []int) int {
	history := make(map[int]bool)
	for _, value := range a {
		history[value] = !history[value]
	}
	for key, value := range history {
		if !value {
			return key
		}
	}
	return -1
}
