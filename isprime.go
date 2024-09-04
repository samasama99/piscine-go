package piscine

func IsPrime(nb int) bool {
	index := 1
	for index*index <= nb {
		if nb%index == 0 {
			return false
		}
	}
	return true
}
