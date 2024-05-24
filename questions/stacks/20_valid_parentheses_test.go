package stacks

func isValid(s string) bool {
	if s == "" {
		return false
	}
	if len(s) % 2 == 1 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if rune(s[i]) == 40 && rune[s[i-1]] == 41 {

		}
	}
}

