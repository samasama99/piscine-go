package piscine

func Index(s string, toFind string) int {
	length := len(toFind)
	for i := 0; i+length < len(s); i++ {
		if s[i:i+length] == toFind {
			return i
		}
	}
	return -1
}
