package two_pointers

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Soultion 1
// time limit exceeded
func Test_maxArea1(t *testing.T) {
	tc := struct {
		i []int
		o int
	}{
		i: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		o: 49,
	}

	area := maxArea1(tc.i)
	assert.Equal(t, tc.o, area)
}

func maxArea1(height []int) int {
	// use two pointers to iterate through the array
	// area = min (height[i], height[j]) * (j - i)
	// find maximum ares
	i := 0
	j := len(height) - 1
	maxArea := 0
	t1 := time.Now().Unix()
	fmt.Println("length: ", len(height))
	for i < j && i < len(height) {
		var area int
		if height[i] < height[j] {
			area = height[i] * (j - i)
		} else {
			area = height[j] * (j - i)
		}
		if maxArea < area {
			maxArea = area
		}
		// pointers increment/decrement
		if j-1 > i {
			j--
		} else {
			j = len(height) - 1
			i++
		}
	}
	t2 := time.Now().Unix()
	fmt.Println("time taken: ", t2-t1)
	return maxArea
}

// Solution 2
// move the pointer which is on the lower height - after checking hint
// Why?
// two ways area can be maximum - a) max length between bars, and b) max height of bars
// a) is coverved by start i an j from extremes. To cover b) case we change pointers as above
func Test_maxArea2(t *testing.T) {
	tc := struct {
		i []int
		o int
	}{
		// i: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		i: []int{1, 2, 3, 0, 0, 0, 1, 1, 3, 2, 1, 2, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1},
		o: 49,
	}

	area := maxArea2(tc.i)
	assert.Equal(t, tc.o, area)
}

func maxArea2(height []int) int {
	// use two pointers to iterate through the array
	// area = min (height[i], height[j]) * (j - i)
	// find maximum ares
	i := 0
	j := len(height) - 1
	maxArea := 0
	t1 := time.Now().UnixMicro()
	fmt.Println("length: ", len(height))
	for i < j && i < len(height) {
		var area int
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if maxArea < area {
			maxArea = area
		}
	}
	t2 := time.Now().UnixMicro()
	fmt.Println("time taken: ", t2-t1)
	return maxArea
}
