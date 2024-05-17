package notes

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
