package practice

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 217 - easy - contains duplicate

// Approach 1 - use hashmap
// TC - O(n), SC - O(k) 1=< k <= n k = unique numbers in the array
func containsDuplicate_706_1(nums []int) bool {
	// use a hahamap to store frequncy
	// if frequency is more than 1 for any number, dupliate is found
	frequency := map[int]bool{}
	for _, num := range nums {
		if _, ok := frequency[num]; ok {
			return true
		}
		frequency[num] = true
	}
	return false
}

// Approach 2 - sorting
// TC - O(nlonn), SC - O(nlogn)
func containsDuplicate_706_2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func Test_containDuplicate_706_3(t *testing.T) {
	testCase := struct {
		i []int
		o bool
	}{
		i: []int{2, 14, 18, 22, 22},
		o: true,
	}

	ans := containsDuplicate_706_3(testCase.i)
	assert.Equal(t, testCase.o, ans)
}

// Approach 3 - two pointers
// Works but time limit exceeded error
// TC - O(n2), SC - O(1)
// Usually, if an algorithm is O(n^2), it can handle n up to around 10^4. It gets Time Limit Exceeded when n≥10^5.
func containsDuplicate_706_3(nums []int) bool {
	i, j := 0, 1
	for i < len(nums)-1 && j < len(nums) {
		if nums[i] == nums[j] {
			return true
		}

		if j == len(nums)-1 {
			i++
			j = i + 1
			continue
		}

		j++
	}
	return false
}

// 219 Contains Duplicate II - easy

// Approach 1 - use two pointers or two for loops
// TC - O(n2), SC - O(1)

func containsNearbyDuplicate_706_1(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	i, j := 0, 1
	for i < len(nums) {
		if nums[i] == nums[j] && i-j <= k {
			return true
		}

		// if j ==
	}
	return false
}

// Approach 2 - use sliding window of k length
// TC - O(n*min(n, k)), SC - O(k)
func containsNearbyDuplicate_706_2(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	arr := []int{}
	for i := range nums {
		for j := range arr {
			if arr[j] == nums[i] {
				return true
			}
		}
		if len(arr) == k {
			arr = arr[1:]
		}
		arr = append(arr, nums[i])
	}
	return false
}

// Approach 3 - Binary Search Tree
// func containsNearbyDuplicate_706_3(nums []int, k int) bool {
// 	//
// }

// 220 - contain duplicates III - Hard

// Approach 1 - bucketing
// TC - O(n), SC - O(min(n, k))
func containAlmostNearbyDuplicates(nums []int, indexDiff int, valueDiff int) bool {
	bucket := map[int]int{}

	// bucket size
	size := valueDiff + 1

	for i := range nums {
		id := nums[i] / size

		if nums[i] < 0 {
			id--
		}

		// check in current bucket
		if j, ok := bucket[id]; ok && i-j <= indexDiff {
			return true
		}

		// check in previous bucket
		if j, ok := bucket[id-1]; ok && i-j <= indexDiff && isLessValueDiff(nums[i], nums[j], valueDiff) {
			return true
		}

		// check in next bucket
		if j, ok := bucket[id+1]; ok && i-j <= indexDiff && isLessValueDiff(nums[i], nums[j], valueDiff) {
			return true
		}

		// store in bucket
		bucket[id] = i
	}
	return false
}

func isLessValueDiff(x int, y int, valueDiff int) bool {
	if x < y {
		return y-x <= valueDiff
	}
	return x-y <= valueDiff
}
