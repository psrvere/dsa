package slidingwindow

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution 1 - brute force. TC O(n3)

func lengthOfLongestSubstring(s string) int {
	// ss := []rune{}
	// for _, char := range s {
	// 	if slices.Index(ss, char) == -1 {
	// 		ss = append(ss, char)
	// 	}
	// }
	// return len(ss)
	ss := []byte{}
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if slices.Index(ss, s[j]) == -1 {
				ss = append(ss, s[j])
			} else {
				break
			}
		}
		if max < len(ss) {
			max = len(ss)
		}
		ss = []byte{}
	}
	return max
}

// Solution 2 - sliding window

// func lengthOfLongestSubstring(s string) int {
// 	// special cases
// 	// length 1
// 	// i := 0 // i and j are left and right ends of sliding window
// 	// j := 1
// 	// for i < len(s) && j < len(s) {
// 	// 	ss := s[i:j]
// 	// 	if j + 1 < len(s) {
// 	// 		j++
// 	// 	}

// 	// }

// 	max := 0
// 	ss := []rune{}
// 	hm := map[rune]bool{}
// 	for i, char := range s {
// 		if val := hm[char]; !val {
// 			hm[char] = true
// 		}
// 		hm[char]
// 		if max < len(hm) {
// 			max = len(hm)
// 		}
// 	}
// }

func Test_lengthOfLongestSubstring2(t *testing.T) {
	tc := struct {
		i string
		o int
	}{
		i: "abcb",
		o: 3,
	}

	ans := lengthOfLongestSubstring2(tc.i)
	assert.Equal(t, tc.o, ans)
}

func lengthOfLongestSubstring2(s string) int {
	chars := make(map[byte]int)

	left := 0
	right := 0

	res := 0
	for right < len(s) {
		r := s[right]
		chars[r]++

		for chars[r] > 1 {
			l := s[left]
			chars[l]--
			left++
		}

		res = max(res, right-left+1)
		right++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
