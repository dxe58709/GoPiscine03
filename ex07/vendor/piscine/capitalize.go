package piscine

func IsDigit(s rune) bool {
	return s >= '0' && s <= '9'
}

func IsAlpha(s rune) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
}

func ToLower(s rune) rune {
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}
	return s
}

func ToUpper(s rune) rune {
	if s >= 'a' && s <= 'z' {
		return s - 32
	}
	return s
}

func Capitalize(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n; i++ {
		if IsAlpha(runes[i]) || IsDigit(runes[i]) {
			if i == 0 || !((IsAlpha(runes[i - 1])) || IsDigit(runes[i - 1])) {
				runes[i] = ToUpper(runes[i])
			} else {
				runes[i] = ToLower(runes[i])
			}
		}
	}
	return string(runes)
}