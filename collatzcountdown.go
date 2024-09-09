package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	return CollatzCountdownRecursive(start, 0)
}

func CollatzCountdownRecursive(start int, count int) int {
	if start == 1 {
		return count
	}
	if start%2 == 0 {
		return CollatzCountdownRecursive(start/2, count+1)
	}
	if start%2 != 0 {
		return CollatzCountdownRecursive(start*3+1, count+1)
	}
	return -1
}
