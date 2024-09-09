package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	_a := a[0]
	remain := a[1:]
	sorted := 0

	for _, _b := range remain {
		if sorted == 0 {
			if f(_a, _b) < 0 {
				sorted = -1
			} else if f(_a, _b) > 0 {
				sorted = 1
			}
		} else {
			if sorted != f(_a, _b) {
				return false
			}
		}
		_a = _b
	}
	return false
}
