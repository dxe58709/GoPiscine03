package piscine

func IsNumeric(s string) bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return false
		}
	}
	return true 
}
