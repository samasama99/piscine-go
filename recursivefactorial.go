package piscine

func recursiveFactorialHelper(nb, current int) int {
	if nb == 0 {
		return 1
	}
	tmp := current * nb
	if tmp < current {
		return 0
	}
	return recursiveFactorialHelper(nb-1, tmp)
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	return recursiveFactorialHelper(nb, 1)
}
