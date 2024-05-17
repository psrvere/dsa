package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

// arrays and slices are homogeneous data structures

func main() {

	// ***** ARRAYS ****** //

	var a [5]int
	fmt.Println("length:", len(a))
	// size of array must be a constant expression i.e. known at compile time
	// once declared, size of an array can not be changed

	// create an array with custom indices
	fmt.Println("creating custom array....")
	type currency int
	const (
		USD currency = iota
		EUR
		GBP
		INR
	)
	customArray := [...]string{USD: "$", EUR: "€", GBP: "£", INR: "₹"}
	fmt.Println("customArray: ", customArray)

	// specifying element value at the time of creation of array
	fmt.Println("creating array with a different default value......")
	diffDefualtValue1 := [3]int{2}
	fmt.Println("example1 - different first value: ", diffDefualtValue1)
	diffDefualtValue2 := [3]int{2: -1}
	fmt.Println("example2 - different last value: ", diffDefualtValue2)
	diffDefualtValue2[0] = 5

	// comparing two arrays
	// two arrays are equal if size and type of array are equal and all the elements at ith index are equal
	fmt.Println("comparing arrays...")
	c1 := [2]int{}
	c2 := [...]int{0, 0}
	c3 := [...]int{1: 3}
	fmt.Println(c1, c2, c3, c1 == c2, c2 == c3, c3 == c1)

	// passing arrays to functions
	// arrays are passed to parameters by copy, hence passing large arrays is inefficient
	// use pointers to an array for efficiency. This will allow function to mutate the original array

	// a function that prints differnt bits in two sha256 hashes
	fmt.Println(".... comparing bits delta between two sha256 hashes ....")
	b1 := sha256.Sum256([]byte{'A'})
	b2 := sha256.Sum256([]byte{'B'})
	byteDiffHash(b1, b2)

	// a function to print SHA256, SHA384 or SHA512 hashes using a string input and a cli flag
	fmt.Println(".... printing hash of requested string input and hashing algo to use (default: SHA256) ....")
	input := "golang"
	algo := os.Args[1]
	printHash(input, algo)

	// ***** SLICES ****** //

}

func byteDiffHash(b1 [32]byte, b2 [32]byte) {
	fmt.Printf("b1: %x\nb2: %x\n", b1, b2)
	var d int
	for i, v := range b1 {
		if v != b2[i] {
			d++
		}
	}
	fmt.Printf("count of different bits: %d\n", d)

}

func printHash(input string, algo string) {
	i := []byte(input)
	if algo == "sha512" {
		h := sha512.Sum512(i)
		fmt.Printf("hash: %x\n", h)
	} else if algo == "sha384" {
		h := sha512.Sum384(i)
		fmt.Printf("hash: %x\n", h)
	} else {
		h := sha256.Sum256(i)
		fmt.Printf("hash: %x\n", h)
	}
}
