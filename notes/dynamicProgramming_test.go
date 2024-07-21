package notes

import (
	"fmt"
	"testing"
)

// Video - freecodecamp course - https://youtu.be/oBt53YbR9Kk?si=F2G4-EMaQGogTuJe

// Part 1 - memoization
// Part 2 - tabulation

// Problem 1 - Write a function fib(n) that takes in a number as an argument. The function should return the n-th number of the Fibonacci sequence.

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

// Problem 3 - Write a function canSum(targetSum, numbers) that takes in a targetSum and an array of numbers as arguments
// The function should return a boolean indicating whether or not it is possible to generate the targetSum using numbers
// from the array
// Contrainsts
// 1) you may use an element of the array as many times as needed
// 2) you may assume that all input numbers are nonnegative

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
