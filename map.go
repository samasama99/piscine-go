package piscine

func Map(f func(int) bool, a []int) []bool {
	into := make([]bool, len(a))
	for i, currentValue := range a {
		into[i] = f(currentValue)
	}
	return into
}
