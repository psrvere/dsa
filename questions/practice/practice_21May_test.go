package practice

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 217 Contains duplicate
// 1) brute force: two loops -> TC - O(n2), SC - O(1)
// 2) use hash map -> TC - O(n) SC - O(n)
// 3) sort the array and check adjacent elements TC - (nlogn), SC - O(1)

func containsDuplicate_21May(nums []int) bool {
	hm := map[int]bool{}
	for _, num := range nums {
		if hm[num] {
			return true
		}
		hm[num] = true
	}
	return false
}

func containsDuplicate_21May_2(nums []int) bool {
	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// 242 Contains duplicates nearby
// 1) brute force: two loops - TC O(n2), SC O(1)
// 2) use hash map - TC O(n), SC O(n)
// 3) maintain a slice of k length

// hash map
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	hm := map[int]int{}
	for i := range nums {
		if val, ok := hm[nums[i]]; ok && i-val <= k {
			return true
		}
		hm[nums[i]] = i
	}
	return false
}

func Test_containsNearbyDuplicate_2(t *testing.T) {
	tc := struct {
		i struct {
			nums []int
			k    int
		}
		o bool
	}{i: struct {
		nums []int
		k    int
	}{nums: []int{1, 2, 3, 1}, k: 3}, o: true}

	ans := containsNearbyDuplicate_2(tc.i.nums, tc.i.k)
	assert.Equal(t, tc.o, ans)
}

// slice of k length
func containsNearbyDuplicate_2(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	arr := []int{}
	for _, num := range nums {
		if slices.Index(arr, num) != -1 {
			return true
		}
		if len(arr) == k {
			arr = arr[1:]
		}
		arr = append(arr, num)
	}
	return false
}

// 220 contains nearby approximate duplicates

// func containsNearbyAlmostDuplicate_21May(nums []int, indexDiff int, valueDiff int) bool {

// }
