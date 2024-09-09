package piscine

func Unmatch(a []int) int {
	history := make(map[int]bool)
	for _, value := range a {
		history[value] = !history[value]
	}
	for _, value := range a {
		if history[value] {
			return value
		}
	}
	return -1
}
