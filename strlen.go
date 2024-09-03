package piscine

func StrLen(s string) int {
	sliceOfRunes := []rune(s)
	return len(sliceOfRunes)
}
