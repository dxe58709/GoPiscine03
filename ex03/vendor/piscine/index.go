package piscine

func Substring(s string, start int, len int) string {
	return s[start : start + len]
}

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i <= len(s) - len(toFind); i++ {
		substr := Substring(s, i, len(toFind))
		if substr == toFind {
			return i
		}
	}
	return -1
}