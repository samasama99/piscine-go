package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}
	_a := a[0]
	remain := a[1:]
	for _, _b := range remain {
		if f(_a, _b) < 0 {
			return false
		}
		_a = _b
	}
	return true
}
