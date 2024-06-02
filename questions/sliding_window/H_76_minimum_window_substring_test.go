package slidingwindow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-window-substring/

// first attempt - did not work for all testcases

func Test_minWindow1(t *testing.T) {
	tc := struct {
		i struct {
			s string
			t string
		}
		o string
	}{
		i: struct {
			s string
			t string
		}{s: "aaaaaaaaaaaabbbbbcdd", t: "abcdd"},
		o: "abbbbbcdd",
	}

	ans := minWindow1(tc.i.s, tc.i.t)
	assert.Equal(t, tc.o, ans)
}

// Testcase1: bba ab ba -
// Testcase2:

func minWindow1(s string, t string) string {
	// add charcters of t in a frequency map
	// run window on s
	// minimum length of window is length of string t
	// start from this length and search for a valid substring
	// if not found, then increase maxLength by 1 i.e. len(t)+1 and search again
	// len(t) <= maxLength <= len(s)

	if len(t) > len(s) {
		return ""
	}

	charFreq := map[byte]int{}
	totalFreq := 0
	for i := range t {
		charFreq[t[i]]++
		totalFreq++
	}

	for maxLength := len(t); maxLength <= len(s); maxLength++ {
		fmt.Println("maxLength: ", maxLength)
		cf := map[byte]int{}
		copyMap(charFreq, cf)
		start := 0
		for end := 0; end < len(s); end++ {
			// if character found in charFreq, reduce count by 1
			if val, ok := cf[s[end]]; ok && val != 0 {
				cf[s[end]]--
			}

			// window is invalid as dictated by maxLength
			if end-start+1 > maxLength {
				if val, ok := cf[s[start]]; ok {
					oc := originalCount(charFreq, s[start])
					if oc-val > 0 {
						cf[s[start]]++
					}
				}
				start++
			}

			// check if all elements of t have been found
			if isMapZero(cf) {
				return s[start : end+1]
			}
		}
	}

	return ""
}

// copies a into b
func copyMap(a map[byte]int, b map[byte]int) {
	// since the unique alphabets are not going to change
	// there is no need to clear the map
	for key, val := range a {
		b[key] = val
	}
}

func isMapZero(cf map[byte]int) bool {
	for _, val := range cf {
		if val != 0 {
			return false
		}
	}
	return true
}

func originalCount(charFreq map[byte]int, char byte) int {
	if val, ok := charFreq[char]; ok {
		return val
	}
	return 0
}

// Approach 1 - Sliding window

func minWindow2(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	if len(s) < len(t) {
		return ""
	}

	fm := map[byte]int{} // frequency map of characters in fm
	for i := range t {
		fm[t[i]]++
	}

	unique := len(fm)              // unique character in t
	start := 0                     // start of window
	found := 0                     // to track count of unique character in t present in current window
	windowCounts := map[byte]int{} // count of characters in the window
	ans := []int{-1, 0, 0}         // window legth, left, right
	for end := 0; end < len(s); end++ {
		// add character at end to the window
		c := s[end]
		windowCounts[c]++

		if _, ok := fm[c]; ok && windowCounts[c] == fm[c] {
			found++
		}

		// contract the window
		for start <= end && found == unique {
			c := s[start]
			// save the smallest window till now
			if ans[0] == -1 || end-start+1 < ans[0] {
				ans[0] = end - start + 1
				ans[1] = start
				ans[2] = end
			}

			// move the window to right
			if _, ok := fm[c]; ok && windowCounts[c] < fm[c] {
				found--
			}
			start++
		}
	}
	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}
