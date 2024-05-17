package two_pointers

import (
	"regexp"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-palindrome/description/

// Solution 1
// using regexp library
func Test_isPalindrome1(t *testing.T) {
	tc := struct {
		i string
		o bool
	}{
		// "A man, a plan, a canal: Panama"
		i: "A man, a plan, a canal: Panama",
		o: true,
	}

	ok := isPalindrome1(tc.i)
	assert.Equal(t, tc.o, ok)
}

func isPalindrome1(s string) bool {

	// convert all uppercase to lowercase
	// remove non alpha numeric letters
	// use two pointers to check if the string is a palindrome

	s = strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "")

	if s == "" {
		return true
	}

	i, j := 0, len(s)-1
	for i < j { // improvement over below if condition
		// if i >= j {
		// 	break
		// }
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// solution 2
// reverse the string and compare + use unicode library

func Test_isPalindrome2(t *testing.T) {
	tc := struct {
		i string
		o bool
	}{
		// "A man, a plan, a canal: Panama"
		i: "A man, a plan, a canal: Panama",
		o: true,
	}

	ok := isPalindrome2(tc.i)
	assert.Equal(t, tc.o, ok)
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	i := 0
	for {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			if i == len(s)-1 {
				s = s[:i]
			}
			s = s[:i] + s[i+1:]
		} else {
			i++
		}

		if i >= len(s) {
			break
		}
	}
	rs := reverseString(s)
	return s == rs
}

func reverseString(s string) string {
	rs := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		rs[i], rs[len(s)-1-i] = rs[len(s)-1-i], rs[i]
	}
	return string(rs)
}
