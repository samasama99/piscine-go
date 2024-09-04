package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	fact := 1

	for nb != 0 {
		tmp := fact
		fact *= nb
		if tmp > fact {
			return 0
		}
		nb--
	}

	return fact
}
