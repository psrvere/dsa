package binarysearch

// Approach 1 - linear search

func findPeakElement1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

// Approach 2 - binary search - iterative
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := l + (r-l)/2
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
