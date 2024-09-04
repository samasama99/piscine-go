package piscine

func IsPrime(nb int) bool {
	if nb == 0 {
		return false
	}
	if nb < 0 {
		nb = nb * -1
	}
	index := 2
	for index*index <= nb {
		if nb%index == 0 {
			return false
		}
	}
	return true
}
