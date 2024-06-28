package arrays

func maxNumberOfBalloons(text string) int {
	// add frequency of chars in a hashmap
	freq := map[byte]int{}

	for i := 0; i < len(text); i++ {
		if text[i] == 'b' || text[i] == 'a' || text[i] == 'l' || text[i] == 'o' || text[i] == 'n' {
			freq[text[i]]++
		}
	}

	count := 0
	for {
		freq['b']--
		freq['a']--
		freq['l'] = freq['l'] - 2
		freq['o'] = freq['o'] - 2
		freq['n']--
		if freq['b'] >= 0 && freq['a'] >= 0 && freq['l'] >= 0 && freq['o'] >= 0 && freq['n'] >= 0 {
			count++
		} else {
			break
		}
	}
	return count
}

// Approach 2
func maxNumberOfBalloons2(text string) int {
	pattern := "balloon"
	return maxNumberOfPatterns(text, pattern)
}

func maxNumberOfPatterns(text string, pattern string) int {
	freqInText := map[int]int{}
	freqInPattern := map[int]int{}

	for i := range text {
		freqInText[int(text[i]-'a')]++
	}

	for i := range pattern {
		freqInPattern[int(pattern[i]-'a')]++
	}

	count := int(^uint(0) >> 1)
	for i := 0; i < 26; i++ {
		if freqInPattern[i] > 0 {
			count = min(count, freqInText[i]/freqInPattern[i])
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
