package notes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Video - freecodecamp course - https://youtu.be/oBt53YbR9Kk?si=F2G4-EMaQGogTuJe

// Part 1 - memoization
// Part 2 - tabulation

// **************************************** //

// Problem 1 - Write a function fib(n) that takes in a number as an argument.
// The function should return the n-th number of the Fibonacci sequence.
// recursion implementation

func Benchmark_Fibonacci(b *testing.B) {
	getFibonacciMemo(50)
}

func getFib(n int) int {
	// base cases
	if n <= 2 {
		return 1
	}
	return getFib(n-1) + getFib(n-2) // not optimised
}

// optimised recursion implementation - tail recursion

// func getFib2(n int) int {

// }

// func recGetFib2(n int, sum int) int {
// 	// base case
// 	if n == 0 {
// 		return sum
// 	}
// 	return recGetFib2(n-1, sum + )
// }

// Question - what is the time complexity and space complexity of Recursion?

// Time space complexity in recursion
// For following function:
// TC: O(N) - n function calls n, n-1, .... 3, 2, 1
// SC: O(N) - storing these function calls in stack
// NOTE: For calculating space complexity in recursion, add stack space used also

func simpleRec(n int) {
	if n <= 1 {
		return
	}
	simpleRec(n - 1)
}

// Another example below
// simpleRec2(5) --> 5 -> 3 -> 1  - total n/2 calls
// simpleRec2(6) --> 6 -> 4 -> 2 -> 1  - total n/2+1 calls
// TC: O(N)
// SC: On(N)

func simpleRec2(n int) {
	if n <= 1 {
		return
	}
	simpleRec(n - 2)
}

// Next example below
// doubleRec(5) --> two doubleRec(4) --> each of these will split in two doubleRec(3)
// this will make a perfectly balanced binary tree with n levels
// TC: 2^N
// SC: N
// NOTE - For a binary tree the number of functions in stack at a given time will not exceed N
// Although we make 2^N function calls, these are not made at once
// At any point maximum of N function calls will be made

func doubleRec(n int) {
	if n <= 1 {
		return
	}
	doubleRec(n - 1)
	doubleRec(n - 1)
}

// Another example below
// levels or heigh = n/2 or n/2 + 1 ~ n/2
// TC: 2^(n/2) ~ 2^n
// SC: n/2 ~ n

func doubleRec2(n int) {
	if n <= 1 {
		return
	}
	doubleRec2(n - 2)
	doubleRec2(n - 2)
}

// So going back to the getFib function
// TC: 2^N
// SC: N

// fib(50) has time complexity of 2^50
// to convert this to power of 10 --> 10^m : m = 50 * log2 ~ 50 * 0.3 ~ 15
// Now fib(50) has time complexity of 10^15 and 50 is a modest number
// We need a better algorithm

// Memoization is one of the overarching strategies to solve dynamic programming problem
// come from memo word - reminder for myself

// Let's use Memoization to solve fib problem

func Benchmark_FibonacciMemo(b *testing.B) {
	getFibonacciMemo(50)
}

func Test_FibonacciMemo(t *testing.T) {
	ans := getFibonacciMemo(50)
	fmt.Println("ans: ", ans)
}

type fibMemo struct {
	fibValue map[int]int
}

func NewFibMemo() *fibMemo {
	return &fibMemo{
		fibValue: map[int]int{},
	}
}

func (m *fibMemo) getFib(n int) int {
	if val, ok := m.fibValue[n]; ok {
		fmt.Println("val: ", val)
		return val
	}
	// base cases
	if n <= 2 {
		return 1
	}
	m.fibValue[n] = getFib(n-1) + getFib(n-2)
	fmt.Println("size: ", len(m.fibValue))
	return m.fibValue[n]
}

func getFibonacciMemo(n int) int {
	memo := NewFibMemo()
	return memo.getFib(n)
}

// The time complexity after memoization is O(n) for above solution
// Space complexity is O(n)

// **************************************** //

// Problem 2: Grid Traveller
// Say that you are a traveller on a 2D grid. You begin in the top left corner and your goal is to
// travel to the bottom right corner. You may only move down or right.
// In how many ways can you travel to the goal on a grid with dimensions m x n?

func Test_gridTraveler(t *testing.T) {
	ans := getRoutesGridTraveler(3, 3)
	fmt.Println(ans)
}

type gridRoutes struct {
	routesMap map[int]map[int]int
}

func NewGridRoutes() *gridRoutes {
	return &gridRoutes{
		routesMap: make(map[int]map[int]int), // initialize outer map
	}
}

func (gr *gridRoutes) gridTraveler(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}

	if _, outerOk := gr.routesMap[m]; outerOk {
		if val, innerOk := gr.routesMap[m][n]; innerOk {
			return val // memo base case
		}
		// initialize inner map
		gr.routesMap[m] = make(map[int]int)
	}

	// initialize outer loop
	gr.routesMap[m] = make(map[int]int)
	gr.routesMap[m][n] = gr.gridTraveler(m-1, n) + gr.gridTraveler(m, n-1)

	// gridTraverler(a, b) = gridTraveler(b, a)
	// hence assign the same value to complementary pair
	gr.routesMap[n] = make(map[int]int)
	gr.routesMap[n][m] = gr.routesMap[m][n] // [TODO] this optimisation is not working - fix it

	fmt.Println(m, n)
	return gr.routesMap[m][n]
}

func getRoutesGridTraveler(m, n int) int {
	routes := NewGridRoutes()
	return routes.gridTraveler(m, n)
}

// for an m, n grid, the depth of binary tree is m + n - 1
// (m,n) node will be broken down till (1, 1) - the last node
// hence, space complexity is O(m+n)
// Time complexsity is 2^(m+n)

// optimisation through memoization
// 1) Certain low level nodes are repeated just like fibonacci
// 2) gridTraveler(m, n) = girdTraveler(n, m)

// NOTE - this of your recursion fuctions as trees and use that tree to implement
// a brute force solution and also look for tends to optimise

// NOTE - Memoization Receipe

// 1. Make it work
// - visualize the problem as a tree
// - implement the tree using recursion
// - test it

// 2. Make it efficient
// - add a memo object - keys represent argument to the function, values represent return values
// enusre memo object is shared in all recursive calls
// - add a new base case which captures the memo
// - store return values in to the memo

// **************************************** //

// Problem 3 - Write a function canSum(targetSum, numbers) that takes in a targetSum and an array of numbers as arguments
// The function should return a boolean indicating whether or not it is possible to generate the targetSum using numbers
// from the array
// Contrainsts
// 1) you may use an element of the array as many times as needed
// 2) you may assume that all input numbers are nonnegative

// [NEW] - this is called table driven testing
// we can use an array of tests to cover multiple testing scenarios in a single test
func Test_canSum(t *testing.T) {
	tests := []struct {
		targetSum int
		nums      []int
		want      bool
	}{
		{7, []int{2, 3}, true},
		{7, []int{5, 3, 4, 7}, true},
		{7, []int{2, 4}, false},
		{8, []int{2, 3, 5}, true},
		{300, []int{7, 14}, false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%v", i), func(t *testing.T) {
			got := canSum(test.targetSum, test.nums)
			assert.Equal(t, test.want, got)
		})
	}
}

func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		newTargetSum := targetSum - numbers[i]
		if newTargetSum < 0 {
			continue
		}
		isPossible := canSum(newTargetSum, numbers)
		if isPossible {
			return true
		}
		// NOTE - we do not want to return false here.
		// We want to return false only after loop is finished
	}

	return false
}

// space and time complexity for above
// m = targetSum and n = length of array
// maximum depth of tree is m in the worst case scenario when we keep deducting 1 on one branch of tree
// branching factor: on each level worst case scenarios is each node makes n more nodes for each element of array
// hence total number of nodes = 1 + n + n^2 + n^3 .... till m levels
// i.e. Time Complexity is O(n^m)
// and Spacce Complexity is O(m)

// Improving time complexity by memoizing above solutions

type canSumWithMemo struct {
	memo map[int]bool
}

func NewCanSumWithMemo() *canSumWithMemo {
	return &canSumWithMemo{
		memo: map[int]bool{},
	}
}

func (sm *canSumWithMemo) canSum(targetSum int, nums []int) bool {
	if targetSum == 0 {
		return true
	}

	for i := 0; i < len(nums); i++ {
		newTargetSum := targetSum - nums[i]
		if newTargetSum < 0 {
			continue
		}
		if val, ok := sm.memo[newTargetSum]; ok {
			return val
		}
		isPossible := sm.canSum(newTargetSum, nums)
		if isPossible {
			sm.memo[newTargetSum] = true
			return true
		} else {
			sm.memo[newTargetSum] = false
		}
	}
	return false
}

func canSumWithMemoFunc(targetSum int, nums []int) bool {
	sm := NewCanSumWithMemo()
	return sm.canSum(targetSum, nums)
}

func Test_canSumWithMemoFunc(t *testing.T) {
	tests := []struct {
		targetSum int
		nums      []int
		want      bool
	}{
		{7, []int{2, 3}, true},
		{7, []int{5, 3, 4, 7}, true},
		{7, []int{2, 4}, false},
		{8, []int{2, 3, 5}, true},
		{300, []int{7, 14}, false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%v", i), func(t *testing.T) {
			got := canSumWithMemoFunc(test.targetSum, test.nums)
			assert.Equal(t, test.want, got)
		})
	}
}

// [NEW]
// Run only a particular test
// go test ./notes -run Test_canSumWithMemoFunc

// Reset testing cache of Golang
// go clean -testcache

// Run test without caching
// go test ./notes -run Test_canSumWithMemoFunc -count=1

// So what is the improved time complexity after memoization?
// Time Complexity O(m * n). We have m levels and at each level there can be a maximum of n nodes for that value
// Space Complexity remains same at O(m)

// **************************************** //

// Problem 4 - howSum
// Write a function 'howSum(targetSum, nums)' that takes in a targetSum and an array of numbers as arguments
// The function should return an array contining any combination of elements that add up to exactly
// The targetSum. If there is no combination that adds up to the targetSum, then return null
// If there are multiple combinations possible, you may return any single one

func Test_howSum(t *testing.T) {
	tests := []struct {
		targetSum int
		nums      []int
		want      []int
	}{
		{7, []int{2, 3}, []int{3, 2, 2}},
		{7, []int{5, 3, 4, 7}, []int{4, 3}},
		{7, []int{2, 4}, nil},
		{8, []int{2, 3, 5}, []int{2, 2, 2, 2}},
		{300, []int{7, 14}, nil},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%v", i), func(t *testing.T) {
			got := howSum(test.targetSum, test.nums)
			assert.Equal(t, test.want, got)
		})
	}
}

func howSum(target int, nums []int) []int {
	// base case
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}
	for _, num := range nums {
		newTarget := target - num
		ret := howSum(newTarget, nums)
		if ret != nil {
			return append(ret, num)
		}
	}
	return nil
}

// Time complexity
// m = target sum, n = length of array
// Recursive tree depth = m
// branching factor = n
// Time Complexity - O(n^m) - these are total recursive calls made
// Space Complexity - O(m) - these are total max recursive calls at a given point

// [NEW]
// In above code append statement time complexity in the worst case
// can be O(n) where the slice size is doubles and all elements are
// compied. the maximum number of elements to be appended in an array
// can be target value itself (i.e. 1 element m times)
// this makes overall time complexity as O(n^m * m)

// Memoizing the solution for better time complexity

type howSumWithMemo struct {
	memo map[int][]int
}

func NewHowSumWithMemo() *howSumWithMemo {
	return &howSumWithMemo{
		memo: make(map[int][]int),
	}
}

func (sm *howSumWithMemo) howSum(target int, nums []int) []int {
	// base case
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}
	for _, num := range nums {
		newTarget := target - num

		var ret []int
		if val, ok := sm.memo[newTarget]; ok {
			ret = val
		} else {
			ret = sm.howSum(newTarget, nums)
			sm.memo[newTarget] = ret
		}

		if ret != nil {
			return append(ret, num)
		}
	}
	return nil
}

func calculateHowSum(target int, nums []int) []int {
	sumWithMemo := NewHowSumWithMemo()
	return sumWithMemo.howSum(target, nums)
}

func Test_calculateHowSum(t *testing.T) {
	tests := []struct {
		targetSum int
		nums      []int
		want      []int
	}{
		{7, []int{2, 3}, []int{3, 2, 2}},
		{7, []int{5, 3, 4, 7}, []int{4, 3}},
		{7, []int{2, 4}, nil},
		{8, []int{2, 3, 5}, []int{2, 2, 2, 2}},
		{300, []int{7, 14}, nil},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%v", i), func(t *testing.T) {
			got := calculateHowSum(test.targetSum, test.nums)
			assert.Equal(t, test.want, got)
		})
	}
}

// Memoized Soulution Time Complexity
// m = target sum, n = length of array
// Recursive tree depth = m
// TBD
// space complexity
// stack depth: O(m)
// memo object space complexity: O(m * m) - m max possible keys with m length array as value
// i.r. Space Complexity: O(m^2)
