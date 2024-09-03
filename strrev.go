package piscine

func StrRev(s string) string {
	sliceOfRunes := []rune(s)
	rev := make([]rune, len(s))
	length := len(sliceOfRunes)

	for i, value := range sliceOfRunes {
		rev[length-i-1] = value
	}

	return string(rev)
}
