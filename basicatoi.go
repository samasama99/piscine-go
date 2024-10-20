package piscine

func pow(num, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res *= num
	}
	return res
}

func BasicAtoi(s string) int {
	res := 0
	length := len(s)
	for i, value := range s {
		res += (int(value) - 48) * pow(10, length-i-1)
	}
	return res
}
