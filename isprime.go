package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	index := 2
	for index*index <= nb {
		if nb%index == 0 {
			return false
		}
		index++
	}
	return true
}
