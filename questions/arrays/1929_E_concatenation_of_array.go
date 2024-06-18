package arrays

// Approach 1
func getConcatenation1(nums []int) []int {
	ans := make([]int, 2*len(nums))
	for i := range nums {
		ans[i], ans[i+len(nums)] = nums[i], nums[i]
	}
	return ans
}

func getConcatenation2(nums []int) []int {
	return append(nums, nums...)
}
