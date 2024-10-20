package piscine

func Sqrt(nb int) int {
	index := 1
	for index*index < nb {
		index++
	}
	if index*index == nb {
		return index
	}
	return 0
}
