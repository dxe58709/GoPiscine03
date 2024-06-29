package piscine

func LastRune(s string) rune {
	if s != "" {
		runes := []rune(s)
		n := len(runes)
		return []rune(s)[n - 1]
		}
	return 0
}
