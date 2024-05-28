package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tc := struct {
		i string
		o bool
	}{
		i: "()[]{}",
		o: true,
	}

	ans := isValid(tc.i)
	assert.Equal(t, tc.o, ans)
}

// Why create hm with closing bracket as key?
// because there are two flows in the program - flow 1 pops last element of stack and check against current iteration value,
// and flow 2 appends the value. Since problem statements mentions - "Open brackets must be closed in the correct order."
// all valid cases will have open brackets before close brackets i.e. all open brackets must go in flow2 and close brackets
// in flow 1

func isValid(s string) bool {
	hm := map[rune]rune{'}': '{', ']': '[', ')': '('}
	stack := []rune{}
	for _, chr := range s {
		// close bracket flow
		if val, ok := hm[chr]; ok {
			var topElem rune
			if len(stack) != 0 {
				topElem = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			if topElem != val {
				return false
			}
			// open bracket flow
		} else {
			stack = append(stack, chr)
		}
	}
	return len(stack) == 0
}
