package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	freq := make(map[string]int)
	count := 0
	for _, s := range Split(str, " ") {
		freq[s]++
		count++
	}
	for _, c := range str {
		if c == ' ' {
			count--
		}
	}
	freq[""] = count
	return freq
}
