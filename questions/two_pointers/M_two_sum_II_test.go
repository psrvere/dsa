package two_pointers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	tc := struct {
		i struct {
			numbers []int
			target  int
		}
		o []int
	}{
		i: struct {
			numbers []int
			target  int
		}{
			numbers: []int{2, 7, 11, 15},
			target:  9,
		},
		o: []int{1, 2},
	}

	// for follo up questions
	// fmt.Println(math.MaxInt64, math.MaxInt64/math.Pow(1, 1))
	// fmt.Println(math.Pow(2, 63) - 1)
	// fmt.Printf("%d", 1<<63-1)

	ans := twoSum(tc.i.numbers, tc.i.target)
	assert.Equal(t, true, reflect.DeepEqual(ans, tc.o))
}

func twoSum(numbers []int, target int) []int {
	lo := 0
	hi := len(numbers) - 1
	for lo < hi {
		if numbers[lo]+numbers[hi] == target {
			return []int{lo + 1, hi + 1}
		} else if numbers[lo]+numbers[hi] < target {
			lo++
		} else if numbers[lo]+numbers[hi] > target {
			hi--
		}
	}
	return []int{}
}

// follow up question - handle overflow
// two solutions - a) cast int32 to int64 or b) detect overflow (num1 > 1 << 31 - 1 - num2)and throw error
