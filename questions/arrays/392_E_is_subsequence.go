package arrays

import "fmt"

/*
Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109,
and you want to check one by one to see if t has its subsequence.
In this scenario, how would you change your code?
*/

// Approach 1 - Two Pointers
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(t) && j < len(s) {
		fmt.Println(i, j)
		if s[j] == t[i] {
			i++
			j++
		} else {
			i++
		}
	}
	fmt.Println("final: ", i, j)
	if j == len(s) {
		return true
	}
	return false
}

// Approach 2 - divide and conquer (recursion) + greedy
func isSubsequence2(s string, t string) bool {
	return checkSubsequence(s, t, 0, 0)
}

func checkSubsequence(s string, t string, si int, ti int) bool {
	// base case 1
	if si == len(s) {
		return true
	}

	// base case 2
	if ti == len(t) {
		return false
	}

	if s[si] == t[ti] {
		si++
	}
	ti++

	return checkSubsequence(s, t, si, ti)
}
