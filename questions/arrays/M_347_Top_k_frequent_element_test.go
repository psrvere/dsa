package arrays

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution 1 - hashmap + array of struct

func Test_topKFrequent(t *testing.T) {
	tc := struct {
		i struct {
			nums []int
			k    int
		}
		o []int
	}{
		i: struct {
			nums []int
			k    int
		}{nums: []int{1, 1, 1, 2, 2, 3}, k: 2},
		o: []int{1, 2},
	}

	ans := topKFrequent(tc.i.nums, tc.i.k)
	assert.Equal(t, true, reflect.DeepEqual(ans, tc.o))
}

func topKFrequent(nums []int, k int) []int {
	hm := map[int]int{}

	for _, num := range nums {
		hm[num]++
	}

	pa := make(PairArray, len(hm))

	i := 0
	for key, value := range hm {
		pa[i] = Pair{key, value}
		i++
	}

	paR := sort.Reverse(pa)
	sort.Sort(paR)

	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, pa[i].Key)
	}

	return ans
}

type Pair struct {
	Key   int
	Value int
}

type PairArray []Pair

func (p PairArray) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairArray) Len() int {
	return len(p)
}

func (p PairArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
