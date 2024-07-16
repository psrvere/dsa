package notes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tail Recursion
// Video - https://www.youtube.com/watch?v=_JtPhF8MshA&ab_channel=Computerphile

// Reguar recursion Example
func Test_getFactorial(t *testing.T) {
	tc := struct {
		i int
		o int
	}{
		i: 5,
		o: 120,
	}

	ans := getFactorial(5)
	assert.Equal(t, tc.o, ans)
}

func getFactorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * getFactorial(n-1)
}

// Above approach will have high number of stack entrie i.e high memory usage
// Because we will go from N to 1 functions and then come out multiplying each layer

// A better appproach is Tail Recursion - see following example
// As we keeping going into deeper level, we also pass accumulator value
// so that when base case is found, the value is returned
// This is more efficient in terms of memory
// In earlier solution we had to make a large intermediate

func Test_getFactorial2(t *testing.T) {
	tc := struct {
		i int
		o int
	}{
		i: 5,
		o: 120,
	}

	ans := getFactorial2(tc.i)
	assert.Equal(t, tc.o, ans)
}

func getFactorial2(n int) int {
	return getFactorialTail(n, 1)
}

func getFactorialTail(n int, sum int) int {
	if n == 1 {
		return sum // return accumulator value instead of number
	}

	return getFactorialTail(n-1, sum*n)
}

// Above solution only helps when the underlying language compiler supports
// Tail Call Optimisation
// In golang it is supported in a limited manner
// more here - https://stackoverflow.com/questions/12102675/tail-call-optimization-in-go

// Another example - fibonacci number

// Normal recursion
// F(0) = 0 - base case
// F(1) = 1 - base case
// F(n) = F(n-1) + F(n-2) - double recursion i.e. two recursive calls is another source of inefficiency
// in addition to the fact that computer has to remember to add the result of these two

// Tail Recursion
// F(n) = go n (0,1)
// go 0 (a,b) = a
// go 1 (a,b) = b
// go n (a,b) = go (n-1) (b, a+b)
// Understand bracket notation with following example
// (2, 3) --> here a = 2 and b = 3. step forward like this --> (3, 5)

// Implementation
// 0, 1, 1, 2, 3, 5, 8, 13, 21.....
// Note:
// 0 is zeroth element of series
// 1 is first element of series
// hence, 21 is 8th element of series
func Test_getFibonacciNumber(t *testing.T) {
	tc := struct {
		i int
		o int
	}{
		i: 8,
		o: 13,
	}

	ans := getFibonacciNumber(tc.i)
	assert.Equal(t, tc.o, ans)
}

func getFibonacciNumber(n int) int {
	return getFibNumberByTailRec(n, 0, 1)
}

func getFibNumberByTailRec(n, a, b int) int {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	return getFibNumberByTailRec(n-1, b, a+b) // single recursion call
}

// What is the difference between head recursion and tail recurion?

// Head recursion happens at the beginning of the method before other processing
// Tail recursion happens at the end of the method after other processing
