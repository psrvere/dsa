package arrays

// Approach 1
func maxProductDifference(nums []int) int {
	// find highest and second highest number
	// find lowest and second lowest number
	highest, index := getHighestNumber(nums)
	nums = append(nums[:index], nums[index+1:]...)

	secondHighest, index := getHighestNumber(nums)
	nums = append(nums[:index], nums[index+1:]...)

	lowest, index := getLowestNumber(nums)
	nums = append(nums[:index], nums[index+1:]...)

	secondLowest, index := getLowestNumber(nums)

	return (highest * secondHighest) - (lowest * secondLowest)
}

func getHighestNumber(nums []int) (int, int) {
	highest := 0
	index := -1
	for i, num := range nums {
		if highest < num {
			highest = num
			index = i
		}
	}
	return highest, index
}

func getLowestNumber(nums []int) (int, int) {
	lowest := 10001
	index := -1
	for i, num := range nums {
		if lowest > num {
			lowest = num
			index = i
		}
	}
	return lowest, index
}

// Approach 2 - shorter logic

func maxProductDifference2(nums []int) int {
	// find highest and second highest number
	// find lowest and second lowest number
	highest, secondHighest := 0, 0
	lowest, secondLowest := 10001, 10001
	for _, num := range nums {
		if highest < num {
			secondHighest = highest
			highest = num
		} else {
			secondHighest = max1(secondHighest, num)
		}

		if lowest > num {
			secondLowest = lowest
			lowest = num
		} else {
			secondLowest = min1(secondLowest, num)
		}
	}
	return (highest * secondHighest) - (lowest * secondLowest)
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}
