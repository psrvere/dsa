package binarysearch

import "sort"

// Approach 1 - O(nlon(n))
func sortedSquares1(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

// Approach 2 - binary search - did not work
func sortedSquares2(nums []int) []int {
	lo := 0
	hi := len(nums) - 1

	// O(log(k)) 0 =< k <= n - k is number of -ve number
	for lo <= hi {
		if nums[lo] >= 0 {
			break
		}

		if -nums[lo] >= nums[hi] {
			nums = insert(nums, lo, hi)
			// lo will remain same as new number has come at that location
			hi--
		}

		if -nums[lo] < nums[hi] {
			hi--
		}
	}

	// O(n)
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}

	return nums
}

func insert(nums []int, lo, hi int) []int {
	numLo := nums[lo]
	nums = append(nums[:lo], nums[lo+1:]...) // remove number at lo
	nums = append(nums[:hi], append([]int{numLo}, nums[hi:]...)...)
	// at that number to just after high
	return nums
}
