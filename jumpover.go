package piscine

func JumpOver(str string) string {
	if len(str) <= 2 {
		return "\n"
	}
	runes := make([]rune, 0)
	for i, c := range str {
		if i != 0 && i%2 == 0 {
			runes = append(runes, c)
		}
	}
	return string(runes)
}
