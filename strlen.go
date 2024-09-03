package piscine

func StrLen(s string) int {
  	res := 0
	for i, _ := range []rune(s) {
		res = i + 1
	}
    return res
}