package piscine

func StrLen(s string) int {
  	res := 0
	for i, _ := range s {
		res = i + 1
	}
    return res
}