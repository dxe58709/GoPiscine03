package piscine

func IsUpper(s string) bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if !(s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}
	return true 
}
