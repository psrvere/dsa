package practice

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 217 Contains Duplicates
// Approaches - a) two loops (O(n2)), b) check adjacent element after sorting array (O(nlogn)), c) use hash map O(n)

func containsDuplicate_20May(nums []int) bool {
	hm := map[int]bool{}
	for _, num := range nums {
		if hm[num] { //
			return true
		}
		hm[num] = true
	}
	return false
}

// 219 Contains Duplicates II
// Aproaches - a) maintain a slice of k length and check for duplicates in it, b) use hash map - both are O(n)
// k can be zero and can be more than n

func Test_containsNearbyDuplicate_20May(t *testing.T) {
	testcase := struct {
		i struct {
			nums []int
			k    int
		}
		o bool
	}{
		i: struct {
			nums []int
			k    int
		}{nums: []int{1, 2, 3, 1, 2, 3}, k: 2},
		o: false,
	}

	ans := containsNearbyDuplicate_20May(testcase.i.nums, testcase.i.k)
	assert.Equal(t, testcase.o, ans)
}

func containsNearbyDuplicate_20May(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	arr := []int{}
	for _, num := range nums {
		if slices.Index(arr, num) != -1 {
			return true
		}
		if len(arr) >= k {
			arr = arr[1:]
		}
		arr = append(arr, num)
	}
	return false
}

// Approach 2 - use hash map

func containsNearbyDuplicate2_20May(nums []int, k int) bool {
	hm := map[int]int{}
	for i, num := range nums {
		if val, ok := hm[num]; ok && i-val <= k {
			return true
		}
		hm[num] = i
	}
	return false
}

// 242 Valid Anagram
// Approach 1 - use hm to count charaters, increase for s and decrease for t

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := map[byte]int{}
	for i := range s {
		hm[s[i]] = hm[s[i]] + 1
		hm[t[i]] = hm[t[i]] - 1
	}
	for _, val := range hm {
		if val != 0 {
			return false
		}
	}
	return true
}
