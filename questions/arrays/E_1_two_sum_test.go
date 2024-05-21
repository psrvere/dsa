package arrays

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	hm := map[int]int{}
	// for i, num := range nums {
	// 	hm[num] = i
	// }
	for i, num := range nums {
		complement := target - num
		if val, ok := hm[complement]; ok && val != i {
			return []int{i, val}
		}
		hm[num] = i
	}
	return []int{}
}

// Solution 1a - notice that i != j condition is not required because in Solution 1 the map was created beforehand
// and In below solution value at i is added to map after the if condition

func Test_TwoSum2(t *testing.T) {
	testcase := struct {
		i struct {
			nums   []int
			target int
		}
		o []int
	}{
		i: struct {
			nums   []int
			target int
		}{
			nums:   []int{3, 3},
			target: 6,
		},
		o: []int{0, 1},
	}

	ans := twoSum2(testcase.i.nums, testcase.i.target)
	assert.Equal(t, true, reflect.DeepEqual(ans, testcase.o))
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}
