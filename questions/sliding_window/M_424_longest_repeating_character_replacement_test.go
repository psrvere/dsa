package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_characterReplacement1(t *testing.T) {
	tc := struct {
		i struct {
			s string
			k int
		}
		o int
	}{
		i: struct {
			s string
			k int
		}{s: "AABA", k: 0},
		o: 2,
	}

	ans := characterReplacement1(tc.i.s, tc.i.k)
	assert.Equal(t, tc.o, ans)
}

func characterReplacement1(s string, k int) int {
	if len(s) == 1 {
		return 1
	}
	left := 0
	right := 1
	count := k
	max := 0
	mm := 1
	firstChangeIndex := 0
	for {
		if s[right] == s[right-1] {
			mm++
			if max < mm {
				max = mm
			}
			if right+1 < len(s) {
				right++
			} else {
				return max
			}
		} else {
			if count > 0 {
				if count == k {
					firstChangeIndex = right
				}
				count--
				mm++
				s = s[:right] + string(s[right-1]) + s[right+1:]
				if max < mm {
					max = mm
				}
				if right+1 < len(s) {
					right++
				} else {
					return max
				}
			} else {
				// if k == 0 {
				// 	return
				// }
				left = firstChangeIndex
				if right+1 < len(s) {
					right = left + 1
				} else {
					return max
				}

				count = k
				mm = 0
			}
		}
	}
}

// Solution 2 - Sliding Window + Binary search

func Test_characterReplacement2(t *testing.T) {
	tc := struct {
		i struct {
			s string
			k int
		}
		o int
	}{
		i: struct {
			s string
			k int
		}{s: "AABCBCB", k: 2},
		o: 5,
	}

	ans := characterReplacement2(tc.i.s, tc.i.k)
	assert.Equal(t, tc.o, ans)
}

// func characterReplacement2(s string, k int) int {
// 	lo := 1
// 	hi := len(s) + 1
// 	for lo+1 < hi {
// 		mid := lo + (hi-lo)/2
// 		if canMakeValidSubstring2(s, mid, k) {
// 			lo = mid
// 		} else {
// 			hi = mid
// 		}
// 	}
// 	return lo
// }

// func canMakeValidSubstring2(s string, substringLength int, k int) bool {
// 	fm := map[byte]int{}
// 	maxFreq := 0
// 	start := 0

// 	for end := 0; end < len(s); end++ {
// 		fm[s[end]]++

// 		if end+1-start > substringLength {
// 			fm[s[start]]--
// 			start++
// 		}

// 		if maxFreq < fm[s[end]] {
// 			maxFreq = fm[s[end]]
// 		}

// 		if substringLength-maxFreq <= k {
// 			return true
// 		}
// 	}
// 	return false
// }

// Appraoch 1 - sliding window + binary search
// for a given length of stsing s, say l - we start with lo = 1, and hi = l+1
// here lo points to the end of longest substring after character replacement and hi points to one level above
// TC - O(nlogn)

func characterReplacement2(s string, k int) int {
	lo := 1
	hi := len(s) + 1
	for lo+1 < hi {
		mid := lo + (hi-lo)/2
		if isValidSubstring2(s, mid, k) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

func isValidSubstring2(s string, substringLength int, k int) bool {
	start := 0
	fm := map[byte]int{}
	maxFreq := 0
	for end := 0; end < len(s); end++ {
		fm[s[end]]++

		if end+1-start > substringLength {
			fm[s[start]]--
			start++
		}

		if maxFreq < fm[s[end]] {
			maxFreq = fm[s[end]]
		}

		if substringLength-maxFreq <= k {
			return true
		}
	}
	return false
}

// Approach 2 - sliding window - slow
// TC - O(mn) where m is count of unique charaters in a string, 0 =< m <= 26 and n is the length of string. Each character enters window
// once and exist window once for ever iteration on a unique character. Hence complextiy is 2O(n) ~ O(n)
// SC - O(m)

func Test_characterReplacement3(t *testing.T) {
	tc := struct {
		i struct {
			s string
			k int
		}
		o int
	}{
		i: struct {
			s string
			k int
		}{s: "AABCBCB", k: 2},
		o: 5,
	}

	ans := characterReplacement3(tc.i.s, tc.i.k)
	assert.Equal(t, tc.o, ans)
}

func characterReplacement3(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	hm := map[byte]bool{}
	for i := range s {
		hm[s[i]] = true
	}

	maxLength := 0
	for char := range hm {
		start, count := 0, 0
		for end := 0; end < len(s); end++ {
			// new character entered window
			if s[end] == char {
				count++
			}

			// if window length (end-start+1) - count of a unique char (count) > max replacement number (k), eject the first character
			if end-start+1-count > k {
				if s[start] == char {
					count--
				}
				start++
			}

			// window is valid at this point - update maxLength
			if maxLength < end-start+1 {
				maxLength = end - start + 1
			}
		}
	}
	return maxLength
}

// Approach 3 - Sliding window - fast
// TC - O(n)
// SC - O(m) 0 <= m <= 26
func characterReplacement4(s string, k int) int {
	start := 0
	maxFreq := 0
	windowSize := 0 // valid window
	freqMap := map[byte]int{}
	for end := 0; end < len(s); end++ {
		freqMap[s[end]]++

		if maxFreq < freqMap[s[end]] {
			maxFreq = freqMap[s[end]]
		}

		// check if window is invalid
		if end-start+1-maxFreq > k {
			// if invalid, eject the left most character from window
			freqMap[s[start]]--
			start++
		} else {
			windowSize = end - start + 1
		}
	}
	return windowSize
}
