package piscine

func JumpOver(str string) string {
	if len(str) <= 2 {
		return "\n"
	}
	return str[2:3] + JumpOver(str[3:])
}
