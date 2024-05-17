package array_questions

import (
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Question:
// https://leetcode.com/problems/longest-consecutive-sequence/

func Test_longestConsecutive1(t *testing.T) {
	testcase := struct {
		i []int
		o int
	}{
		i: []int{100, 4, 200, 1, 3, 2},
		o: 4,
	}

	ans := longestConsecutive1(testcase.i)
	assert.Equal(t, true, ans == testcase.o)
}

func longestConsecutive1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// first, remove duplicates
	// add data to map
	m := make(map[int]int)
	for _, num := range nums {
		m[num] = m[num] + 1
	}

	// add data back to an array
	a := []int{}
	for key := range m {
		a = append(a, key)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	// count sequence
	count := make([]int, len(a))
	count[0] = 1
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] == 1 {
			count[i] = count[i-1] + 1
		} else {
			count[i] = 1
		}
	}

	// get max value
	ans := slices.Max(count)
	return ans
}

// Solution 2
// sort the array, no need to remove duplicates

func Test_longestConsecutive2(t *testing.T) {
	testcase := struct {
		i []int
		o int
	}{
		i: []int{100, 4, 200, 1, 3, 2},
		o: 4,
	}

	ans := longestConsecutive2(testcase.i)
	assert.Equal(t, true, ans == testcase.o)
}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	count := 1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 0 {
			continue
		} else if nums[i]-nums[i-1] == 1 {
			count++
		} else {
			if max < count {
				max = count
			}
			count = 1
		}
	}
	if max < count {
		max = count
	}
	return max
}

// Solution 3
// use only map to get answer
// [REPEAT]

func longestConsecutive3(nums []int) int {
	num_set := make(map[int]bool)
	for _, num := range nums {
		num_set[num] = true
	}
	longestStreak := 0
	for num := range num_set {
		if _, exists := num_set[num-1]; !exists {
			currentNum := num
			currentStreak := 1
			for exists := num_set[currentNum+1]; exists; exists = num_set[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
