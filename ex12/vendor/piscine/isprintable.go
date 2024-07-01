package piscine

func IsPrintable(s string) bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if !(s[i] >= 32 && s[i] <= 126) {
			return false
		}
	}
	return true
}