package dsimplementations

import (
	"container/heap"
	"fmt"
	"testing"
)

type Heap struct {
	Key   int
	Value int
}

type MinHeap []Heap

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].Value < m[j].Value }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(h any) {
	*m = append(*m, h.(Heap))
}

func (m *MinHeap) Pop() any {
	// code 1
	// *m = (*m)[0 : len(*m)-1]

	// code2
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func Test_MinHeap(t *testing.T) {
	h := &MinHeap{Heap{Key: 1, Value: 2}, Heap{Key: 2, Value: 4}, Heap{Key: 3, Value: 1}}
	heap.Init(h)
	heap.Push(h, Heap{Key: 4, Value: 0})
	fmt.Println(h)
	// Output:
	// [{4 0} {3 1} {1 2} {2 4}]
}
