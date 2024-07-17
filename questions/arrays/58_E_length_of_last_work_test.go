package arrays

import (
	"fmt"
	"strings"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	x := "fly  abc"
	lengthOfLastWord(x)
}

func lengthOfLastWord(s string) int {
	y := strings.Split(s, " ")
	fmt.Println(y)
	return 1
}

// Approach 1 - did not work
// if there are multiple spaces between words, it is stored in the array

func lengthOfLastWord1(s string) int {
	arr := strings.Split(s, " ")
	return len(arr[len(arr)-1])
}

// Approach 2 - iterate from backwards
func lengthOfLastWord2(s string) int {
	count := 0
	found := false
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " && !found {
			continue
		} else if string(s[i]) == " " && found {
			break
		}
		found = true
		count++
	}
	return count
}

// Approach 3 - trim trailing spaces

// Approach 4a - practice session
func lengthOfLastWord4a(s string) int {
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == " " || arr[i] == "" {
			continue
		}
		return len(arr[i])
	}
	return 0
}

// Approach 4b using strings.Fields
func lengthOfLastWord4b(s string) int {
	arr := strings.Fields(s)
	return len(arr[len(arr)-1])
}
