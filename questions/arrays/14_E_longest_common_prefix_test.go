package arrays

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// [Greedy], [BinarySearch], [Divide&Counquer]

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"a"}
	longestCommonPrefix(strs)
}

// Approach 1 - self

func longestCommonPrefix(strs []string) string {

	index := 0
	ans := ""
	for {
		breakOuter := false
		var char byte
		if index < len(strs[0]) {
			char = strs[0][index]
		} else {
			breakOuter = true // handles case [""]
		}

		for i := 0; i < len(strs); i++ { // i = 0 handles case ["x"]
			if index >= len(strs[i]) {
				breakOuter = true
				break
			}
			if strs[i][index] != char {
				breakOuter = true
				break
			}
			if i == len(strs)-1 {
				ans = ans + string(char)
			}
			fmt.Println(ans)
		}
		if breakOuter {
			break
		}
		index++
	}
	return ans
}

// Another shorter implementation of above
func longestCommonPrefix1b(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	for j := 0; j < len(strs[0]); j++ {
		addToPrefix := string(strs[0][j])
		breakOuter := false
		for i := 0; i < len(strs); i++ {
			if j >= len(strs[i]) || string(strs[i][j]) != addToPrefix {
				breakOuter = true
				break
			}
		}
		if breakOuter {
			break
		}
		prefix += addToPrefix
	}
	return prefix
}

// Approach 2 - horizontal scanning
// take first word as prefix
// iterate over the rest of the words in array
// for each word, check if prefix is found
// if not then keep shortening the prefix from last
// TC - O(n) - when all the words are same
// TC - O(l) - when the first word is the longest and no match found with the second word, l is the length of first word
// TC  O(max(n, l))

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := range strs {
		// loop until prefix is found
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// Approach 3 - vertical scanning - compring each character at a time
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := range strs {
			if i > len(strs[j])-1 || char != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func Test_longestCommonPrefix4(t *testing.T) {
	tc := struct {
		i []string
		o string
	}{
		i: []string{"flower", "flower", "flower", "flower"},
		o: "flower",
	}
	ans := longestCommonPrefix4(tc.i)
	assert.Equal(t, tc.o, ans)
}

// Approach 4 - Divide and Conquer
// [DEBUG] - did not work, index going out of range in lcpRight expression
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var LCP = func(str1, str2 string) string {
		minLength := len(str1)
		if len(str2) < len(str1) {
			minLength = len(str2)
		}
		for i := 0; i < minLength; i++ {
			if str1[i] != str2[i] {
				return str1[:i]
			}
		}
		return str1[:minLength]
	}

	// [NEW] - this can not be merged in single line as we will not
	// be able to call in recursively then
	var divideAndConquer func(strs []string, l, r int) string
	divideAndConquer = func(strs []string, l, r int) string {
		// base cases for recursion
		if len(strs) == 2 {
			return LCP(strs[0], strs[1])
		} else if len(strs) == 1 {
			return strs[0]
		}

		mid := (l + r) / 2
		lcpLeft := divideAndConquer(strs[l:mid], l, mid)
		lcpRight := divideAndConquer(strs[mid:r+1], mid, r)
		return LCP(lcpLeft, lcpRight)
	}

	return divideAndConquer(strs, 0, len(strs)-1)
}
