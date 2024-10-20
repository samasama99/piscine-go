package piscine

func Compact(ptr *[]string) int {
	newArr := make([]string, 0)
	currentArr := *ptr

	for _, s := range currentArr {
		if len(s) != 0 {
			newArr = append(newArr, s)
		}
	}
	*ptr = newArr
	return len(newArr)
}
