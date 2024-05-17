package array_questions

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf1(t *testing.T) {
	testcase := struct {
		i []int
		o []int
	}{
		i: []int{1, 2, 3, 4},
		o: []int{24, 12, 8, 6},
	}

	output := productExceptSelf1(testcase.i)
	assert.Equal(t, true, reflect.DeepEqual(output, testcase.o))
}

func productExceptSelf1(nums []int) []int {
	ans := make([]int, len(nums))
	i := 0
	j := 0
	prod := 1
	for {
		if j == i {
			if i == len(nums)-1 {
				ans[i] = prod
				break
			}
			j++
			continue
		}
		prod = prod * nums[j]
		if j == len(nums)-1 {
			ans[i] = prod
			i++
			j = 0
			prod = 1
			continue
		}
		j++
	}
	return ans
}

// Solution 2 - user prefix and suffix products

func TestProductExceptSelf2(t *testing.T) {
	testcase := struct {
		i []int
		o []int
	}{
		i: []int{1, 2, 3, 4},
		o: []int{24, 12, 8, 6},
	}

	output := productExceptSelf2(testcase.i)
	assert.Equal(t, true, reflect.DeepEqual(output, testcase.o))
}

func productExceptSelf2(nums []int) []int {
	ans := make([]int, len(nums))

	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0] = 1
	suffix[len(nums)-1] = 1

	// prefix product array
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	// suffix product array
	for i := len(suffix) - 2; i > -1; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < len(ans); i++ {
		ans[i] = prefix[i] * suffix[i]
	}

	return ans
}

// Solution 2a without creating prefix and suffix arrays

func TestProductExceptSelf2a(t *testing.T) {
	testcase := struct {
		i []int
		o []int
	}{
		i: []int{1, 2, 3, 4},
		o: []int{24, 12, 8, 6},
	}

	output := productExceptSelf2a(testcase.i)
	assert.Equal(t, true, reflect.DeepEqual(output, testcase.o))
}

func productExceptSelf2a(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	ans[len(nums)-1] = 1

	// prefix product array
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	// suffix product array
	suffix := 1
	for i := len(ans) - 2; i > -1; i-- {
		suffix = suffix * nums[i+1]
		ans[i] = ans[i] * suffix
	}

	return ans
}
