package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/contains-duplicate-iii/editorial/

// Solution 1 - brute force - two loops - time limit exceeded
func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	tc := struct {
		i struct {
			nums      []int
			indexDiff int
			valueDiff int
		}
		o bool
	}{
		i: struct {
			nums      []int
			indexDiff int
			valueDiff int
		}{nums: []int{8, 7, 15, 1, 6, 1, 9, 15}, indexDiff: 2, valueDiff: 3},
		o: true,
	}

	ans := containsNearbyAlmostDuplicate(tc.i.nums, tc.i.indexDiff, tc.i.valueDiff)
	assert.Equal(t, tc.o, ans)
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	hm := map[int]int{}
	for i, num := range nums {
		for j := 0; j <= valueDiff; j++ {
			if j < num {
				num = num - j
			} else {
				num = j - num
			}
			if val, ok := hm[num]; ok && i-val <= indexDiff {
				return true
			}
		}
		hm[num] = i
	}
	return false
}

// Solution 2 - bucketing approach

func Test_containsNearbyAlmostDuplicate2(t *testing.T) {
	tc := struct {
		i struct {
			nums      []int
			indexDiff int
			valueDiff int
		}
		o bool
	}{
		i: struct {
			nums      []int
			indexDiff int
			valueDiff int
		}{nums: []int{-3, 3}, indexDiff: 2, valueDiff: 4},
		o: false,
	}

	ans := containsNearbyAlmostDuplicate2(tc.i.nums, tc.i.indexDiff, tc.i.valueDiff)
	assert.Equal(t, tc.o, ans)
}

func containsNearbyAlmostDuplicate2(nums []int, indexDiff int, valueDiff int) bool {
	if valueDiff < 0 {
		return false
	}
	// bucket size
	bs := valueDiff + 1
	if bs == 0 {
		return false
	}
	bucket := map[int]int{}
	for i, num := range nums {
		// handle -ve numbers

		// bucket id
		id := num / bs

		// offset the id for negative numbers
		// TC {-3, 3, 6}, 2, 3
		if num < 0 {
			id--
		}

		// check in the same bucket
		if j, ok := bucket[id]; ok && i-j <= indexDiff {
			return true
		}

		// check in the previous bucket
		if j, ok := bucket[id-1]; ok && i-j <= indexDiff && abs(nums[i]-nums[j]) <= valueDiff {
			return true
		}

		// check in the next bucket
		if j, ok := bucket[id+1]; ok && i-j <= indexDiff && abs(nums[i]-nums[j]) <= valueDiff {
			return true
		}

		// if not found in any of the buckets, add it's index to the current bucket
		bucket[id] = i
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Solution 3 - Using BST
