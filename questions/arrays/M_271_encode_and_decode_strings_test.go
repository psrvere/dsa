package arrays

// question - what should be the delimiter

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Codec struct {
	wordLength []int // for storing word length in Solution 1
}

// SOLUTION 1 - use an array to store length of each word and use this length to retrieve the word from combined string

func TestS1(t *testing.T) {
	testcase := struct {
		i []string
		o []string
	}{
		i: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
		o: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
	}

	c := Codec{}
	encoded := c.Encode_S1(testcase.i)
	decoded := c.Decode_S1(encoded)
	assert.Equal(t, true, reflect.DeepEqual(decoded, testcase.o))
}

func (c *Codec) Encode_S1(strs []string) string {
	c.wordLength = []int{}
	var s string
	for _, str := range strs {
		s = s + str
		c.wordLength = append(c.wordLength, len(str))
	}
	return s
}

func (c *Codec) Decode_S1(str string) []string {
	i := 0
	words := []string{}
	for _, l := range c.wordLength {
		word := str[i : i+l]
		words = append(words, word)
		i = i + l
	}
	return words
}

// Solution 1a
// strings are primitive type, read only i.e. every mutation of it will create a new string
// use strings.Builder for efficient concatenation of strings

func TestS1a(t *testing.T) {
	testcase := struct {
		i []string
		o []string
	}{
		i: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
		o: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
	}

	c := Codec{}
	encoded := c.Encode_S1a(testcase.i)
	decoded := c.Decode_S1a(encoded)
	assert.Equal(t, true, reflect.DeepEqual(decoded, testcase.o))
}

func (c *Codec) Encode_S1a(strs []string) string {
	c.wordLength = []int{}
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		c.wordLength = append(c.wordLength, len(str))
	}
	return sb.String()
}

func (c *Codec) Decode_S1a(str string) []string {
	// Also added an optimisaton to handle empty word array efficiently
	if len(c.wordLength) == 0 {
		return []string{}
	}

	i := 0
	words := []string{}
	for _, l := range c.wordLength {
		word := str[i : i+l]
		words = append(words, word)
		i = i + l
	}
	return words
}

// Notes on strings.Builder - implements bytes array underneath
// Solution 1b - implement your own version of strings.Builder to solve the problem

func TestS1b(t *testing.T) {
	testcase := struct {
		i []string
		o []string
	}{
		i: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
		o: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
	}

	c := Codec{}
	encoded := c.Encode_S1b(testcase.i)
	decoded := c.Decode_S1b(encoded)
	assert.Equal(t, true, reflect.DeepEqual(decoded, testcase.o))
}

func (c *Codec) Encode_S1b(strs []string) string {
	c.wordLength = []int{}
	sb := []byte{}
	for _, str := range strs {
		sb = append(sb, str...)
		c.wordLength = append(c.wordLength, len(str))
	}
	return string(sb)
}

func (c *Codec) Decode_S1b(str string) []string {
	// Also added an optimisaton to handle empty word array efficiently
	if len(c.wordLength) == 0 {
		return []string{}
	}

	i := 0
	words := []string{}
	for _, l := range c.wordLength {
		word := str[i : i+l]
		words = append(words, word)
		i = i + l
	}
	return words
}

// ****************************************** //

// Solution 2 - Using a non ASCII delimiter
// use unicode charater which is not part of ASCII like pie, sigma, delta
func TestS2(t *testing.T) {
	testcase := struct {
		i []string
		o []string
	}{
		i: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
		o: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
	}

	c := Codec{}
	encoded := c.Encode_S2(testcase.i)
	decoded := c.Decode_S2(encoded)
	assert.Equal(t, true, reflect.DeepEqual(decoded, testcase.o))
}

var pi string = fmt.Sprintf("%c", 0x03C0)

func (c *Codec) Encode_S2(strs []string) string {
	var encoded string
	for i, s := range strs {
		if i == 0 {
			encoded = s
			continue
		}
		encoded = encoded + pi + s
	}
	return encoded
}

func (c *Codec) Decode_S2(str string) []string {
	return strings.Split(str, pi)
}

// ****************************************** //

// Solution 3 - using escaping

func TestS3(t *testing.T) {
	testcases := []struct {
		i []string
		o []string
	}{
		{
			i: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
			o: []string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
		},
		{
			i: []string{"V", "Grz/"},
			o: []string{"V", "Grz/"},
		},
	}

	c := Codec{}
	for _, tc := range testcases {
		encoded := c.Encode_S3(tc.i)
		decoded := c.Decode_S3(encoded)
		assert.Equal(t, true, reflect.DeepEqual(decoded, tc.o))
	}
}

func (c *Codec) Encode_S3(strs []string) string {
	var encoded string
	for _, str := range strs {
		idx := strings.Index(str, "/")
		if idx == -1 {
			encoded = encoded + str + "/:"
			continue
		}
		// escape the forward slash
		str = str[:idx] + "/" + str[idx:]
		encoded = encoded + str + "/:"
	}
	return encoded
}

// scenarios
// no '/' - no problem
// '/' exists
//     - start: /abc ecoding --> //abc/: decoding --> search for //, replace with /
//     - mid: a/bc encoding --> a//bc/: decoding --> search for //, replace with /
//     - end: abc/ --> abc///: decoding --> above logic will fail here
// '//' exists - above logic will fail in all cases
//    - start: //abc encoding --> ///abc/: decoding --> above logic will fail here
//    - mid:
//    - end:
// '/:' exists - above logic will fail in the end case
//    - start: /:abc encoding --> //:abc/:

func (c *Codec) Decode_S3(strs string) []string {
	// first, remove escaped forward slash
	var arr []string
	i := 0
	for {
		i = strings.Index(strs, "//")
		if i == -1 {
			break
		}

		if strs[i:i+3] == "//:" {
			// extract the word
			arr = append(arr, strs[:i+1])
			strs = strs[i+3:]
			if strs == "" {
				break
			}
		} else if strs[i:i+2] == "//" {
			// replace '//' with '/'
			strs = strs[:i] + strs[i+1:]
		}
	}
	return arr
}
