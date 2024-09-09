package piscine

func ForEach(f func(int), a []int) {
	for _, currentValue := range a {
		f(currentValue)
	}
}
