package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	freq := make(map[string]int)
	for _, s := range Split(str, " ") {
		freq[s]++
	}
	return freq
}
