package notes

import "fmt"

func stringNotes() {

	// 1. Using elipsis with Strings

	s := "abc"
	sArr := []string{}
	bArr := []byte{}

	// using elipsis on string to unpack char - DOES NOT work
	// sArr = append(sArr, s...) - inavalid syntax

	// usng elipsis on string to upack bytes of string
	bArr = append(bArr, s...) // valid syntax

	// to unpack characters of string, use for loop with range
	for _, char := range s {
		sArr = append(sArr, string(char)) // char is a rune type here
	}

	// 2. mutating string is expsensive if done several times. this is because string ia primitive type,
	// ready only i.e. every mutation creates new string. use byte array instead.
}

func Example_byteString() {
	// create a byte of character "a"
	// first approach - use bracket notation
	var str string = "a"
	x := str[0]
	fmt.Printf("value: %x, type: %T", x, x)

	// second approach - use character notation
	var b byte = 'a' // or b := byte(':')
	fmt.Printf("value: %x, type: %T", b, b)

	// but can not create byte of "" empty string
	// charter literal doesn't exist for empty string
	// bracket notation will not work for empty string
	// we can create byte 0 which is empty byte
	var zero byte = 0
	fmt.Printf("value: %x, type: %T", zero, zero)
}
