package two_pointers

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/3sum/description/

func Test_threeSum1(t *testing.T) {
	tc := struct {
		i []int
		o [][]int
	}{
		i: []int{-1, 0, 1, 2, -1, -4},
		o: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	}

	ans := threeSum1(tc.i)
	assert.Equal(t, true, reflect.DeepEqual(ans, tc.o))
}

func threeSum1(nums []int) [][]int {
	numsU := removeDuplicates(nums)
	i := 0
	j := 0
	k := 0
	ans := [][]int{}
	for i < len(numsU) {
		// skip logic
		if i == j || j == k || k == i {
			k++
			if k >= len(numsU)-1 {
				k = 0
				j++
				if j >= len(numsU)-1 {
					j = 0
					i++
				}
			}
			continue
		}
		// selection logic
		if numsU[i]+numsU[j]+numsU[k] == 0 {
			ans = append(ans, []int{numsU[i], numsU[j], numsU[k]})
		}
		// increment
		k++
		// break logic
		// if k >= len(numsU)-1 {
		// 	k = 0
		// 	j++
		// 	if j >= len(numsU)-1 {
		// 		j = 0
		// 		i++
		// 	}
		// }
	}
	return ans // remove duplicat arrays from ans ?
}

func removeDuplicates(nums []int) []int {
	m := make(map[int]bool)
	for _, num := range nums {
		if !m[num] {
			m[num] = true
		}
	}

	arr := []int{}
	for num := range m {
		arr = append(arr, num)
	}

	return arr
}

// Solution 2
// Using two sum II method

func Test_threeSum2(t *testing.T) {
	tc := struct {
		i []int
		o [][]int
	}{
		i: []int{-1, 0, 1, 2, -1, -4},
		o: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	}

	ans := threeSum2(tc.i)
	assert.Equal(t, true, reflect.DeepEqual(ans, tc.o))
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums) // internally calls slices.Sort
	// there is another sort.Sort which requires less function for sorting

	ans := [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; i++ { // no cases will be found above nums[i] > 0
		if i == 0 || nums[i-1] != nums[i] { // i == 0 to avoid out of index error & second condition to avoid running loops for duplicate numbers
			twoSum2(i, nums, &ans)
		}
	}
	return ans
}

func twoSum2(i int, nums []int, ans *[][]int) {
	lo, hi := i+1, len(nums)-1
	for lo < hi {
		sum := nums[i] + nums[lo] + nums[hi]
		if sum < 0 {
			lo++
		} else if sum > 0 {
			hi--
		} else {
			*ans = append(*ans, []int{nums[i], nums[lo], nums[hi]}) // can not append with a pointer
			lo++
			hi--
			for lo < hi && nums[lo] == nums[lo-1] { // good addition
				lo++
			}
		}
	}
}

// Solution 3
// Using Hashset in two sum function

func Test_threeSum3(t *testing.T) {
	tc := struct {
		i []int
		o [][]int
	}{
		i: []int{-2, 0, 1, 1, 2},
		o: [][]int{{-2, 0, 2}, {-2, 1, 1}},
	}
	// tc := struct {
	// 	i []int
	// 	o [][]int
	// }{
	// 	i: []int{-1, 0, 1, 2, -1, -4},
	// 	// o: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	// 	o: [][]int{{-1, 0, 1}, {-1, 2, -1}},
	// }

	ans := threeSum3(tc.i)
	// fmt.Println("ans: ", ans)
	// fmt.Println("tc.0: ", tc.o)
	// fmt.Println(reflect.DeepEqual(ans, tc.o))

	// a := [3]int{1, 2, 3}
	// b := [3]int{3, 2, 1}
	// fmt.Println(a == b)

	// assert.ElementsMatch(t, tc.o, ans)

	assert.Equal(t, true, reflect.DeepEqual(ans, tc.o))
}

func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] { // beautiful condition
			twoSum3(i, nums, &ans)
		}
	}
	return ans
}

func twoSum3(i int, nums []int, ans *[][]int) {
	hash := make(map[int]bool)
	for j := i + 1; j < len(nums); j++ {
		complement := -nums[i] - nums[j]
		if hash[complement] {
			*ans = append(*ans, []int{nums[i], nums[j], complement})
			for j+1 < len(nums) && nums[j] == nums[j+1] { // beautiful condition
				j++
			}
		}
		hash[nums[j]] = true
	}
}

// func threeSum3(nums []int) [][]int {
// 	sort.Ints(nums)
// 	res := [][]int{}
// 	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
// 		if i == 0 || nums[i-1] != nums[i] {
// 			twoSum3(nums, i, &res)
// 		}
// 	}
// 	return res
// }

// func twoSum3(nums []int, i int, res *[][]int) {
// 	seen := map[int]bool{}
// 	for j := i + 1; j < len(nums); j++ {
// 		complement := -nums[i] - nums[j]
// 		if seen[complement] {
// 			*res = append(*res, []int{nums[i], nums[j], complement})
// 			for j+1 < len(nums) && nums[j] == nums[j+1] {
// 				j++
// 			}
// 		}
// 		seen[nums[j]] = true
// 	}
// }
