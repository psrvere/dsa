package arrays

// Boyer-Moore Voting Algorithm
// can do second pass to count and ensure candidate is a majority element or not
func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}

//[Divide&Conqueer]
