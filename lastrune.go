package piscine

func LastRune(s string) rune {
	sliceOfRunes := []rune(s)
	length := len(sliceOfRunes)
	return sliceOfRunes[length-1]
}
