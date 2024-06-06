package practice

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Encode and Decode strings
// Approach - 1 Escaping

type Codec struct {
}

// use / as escaping character
// cases:
// case 1 - no / in string --> find and replace blank space with /:
// case 2 - / present in string in mid --> find and replace / with //
// case 3 - / present at the starting of a word - example: He is /smart --> He/:is///:smart
// case 4 - / present at the end of a word - example: He is/ smart --> He/:is///:smart
// case 5 - multiple / present // --> //// - escape each / present --> total number of / will always be even

// Encodes a list of strings to a single string.
func (codec *Codec) Encode_6July(strs []string) string {
	for i := range strs {
		str := strs[i]
		idx := 0

		// replace all '/' with '//'
		for idx != -1 {
			idx := strings.Index(str, "/")
			if idx != -1 {
				str = str[:idx] + "/" + str[idx:]
				strs[i] = str
			}
		}

		// add '/:' in the and
		str = str + "/:"
		strs[i] = str
	}
	return strings.Join(strs, "")
}

// first try - it is complex - not tested.
// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	for {
		idx := strings.Index(strs, "/:")
		if idx == -1 {
			break
		}
		if string(strs[idx+1]) == ":" {
			if idx+2 < len(strs) {
				strs = strs[:idx-1] + strs[idx+2:]
			} else {
				strs = strs[:idx-1]
			}
		} else { // it is '/'
			// case 1 - ////a even number of slashes followed by another character
			// case 2 - /////: even number of slashes followed by /:
			j := 1
			for {
				if strs[idx+j] == '/' && strs[idx+j+1] != ':' { // possible more forward slashes
					j++
					continue
				} else if strs[idx+j] == '/' && strs[idx+j+1] == ':' { // end of word
					if j%2 != 0 {
						log.Fatal("j is not even, strs: ", strs)
					}

					slashes := ""
					for k := 1; k <= j; k++ {
						slashes += "/"
					}

					var rhalf string
					if idx+j+1 < len(strs) {
						rhalf = strs[idx+j+1:]
					} else {
						rhalf = ""
					}

					strs = strs[:idx] + slashes + rhalf
					break
				} else { // end of forward slashes followed by another character
					if j%2 != 0 {
						log.Fatal("j is not even, strs: ", strs)
					}

					slashes := ""
					for k := 1; k <= j; k++ {
						slashes += "/"
					}

					strs = strs[:idx] + slashes
					break
				}
			}
		}
	}
	return []string{}
}

func Test_EncodeDecode_6July2(t *testing.T) {
	tc := struct {
		i []string
		o []string
	}{
		i: []string{"V", "Grz/"},
		o: []string{"V", "Grz/"},
	}

	codec := Codec{}
	encoded := codec.Encode_6June2(tc.i)
	decoded := codec.Decode_6June2(encoded)
	assert.Equal(t, tc.o, decoded)
}

func (codec *Codec) Encode_6June2(strs []string) string {
	for i := range strs {
		str := strs[i]
		new := ""
		for j := range str {
			if str[j] == '/' {
				// str = str[:j] + "//" + str[j:]
				new += "//"
			} else {
				new += string(str[j])
			}
		}
		new += "/:"
		strs[i] = new
	}
	return strings.Join(strs, "")
}

// after seeing solution
func (codec *Codec) Decode_6June2(strs string) []string {
	response := []string{}
	currentString := ""
	i := 0
	for i < len(strs)-1 {
		twoChars := string(strs[i]) + string(strs[i+1])
		if twoChars == "/:" {
			response = append(response, currentString)
			currentString = ""
			i += 2
		} else if twoChars == "//" {
			currentString += "/"
			i += 2
		} else { // strs[i] not /, it is other character
			currentString += string(strs[i])
			i++
		}
	}
	return response
}

// Learning from above solution
// it was more intuitive to handle one character at a time
// if charter is / replace with // etc
// dealing wih multiple charters str[i:i+2] etc involved handling many case
// and also added complexity of handling i+2 < len(str)
