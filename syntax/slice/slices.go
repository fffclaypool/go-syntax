/* Slices
  Go Slice is an abstraction over Go Array. Go Array allows you to define variables that
can hold several data items of the same kind but it does not provide any inbuilt method
to increase its size dynamically or get a sub-array of its own. Slices overcome this
limitation. It provides many utility functions required on Array and is widely used in
Go programming.

  To define a slice, you can declare it as an array without specifying its size.
Alternatively, you can use make function to create a slice.

	var numbers []int # a slice of unspecified size#
	# numbers = []int{0,0,0,0,0}
	numbers = make([]int,5,5) # a slice of length 5 and capacity 5
*/

package main

import "fmt"

func main() {
	lenAndcapFunctions()
	nilSlice()
	subslicing()
	appendAndcopyFunctions()
}

func lenAndcapFunctions() {
	/*
		  A slice is an abstraction over array. It actually uses arrays as an underlying
		structure. The len() function returns the elements presents in the slice where
		cap() function returns the capacity of the slice (i.e., how many elements it can
		be accommodate). The following example explains the usage of slice −
	*/

	var numbers = make([]int, 3, 5)
	printSlice(numbers)
}

func nilSlice() {
	// If a slice is declared with no inputs, then by default, it is initialized as nil.
	// Its length and capacity are zero.

	var numbers []int
	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("slice is nil")
	}
}

func subslicing() {
	//   Slice allows lower-bound and upper bound to be specified to get the subslice of
	// it using[lower-bound:upper-bound].

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	fmt.Println("numbers ==", numbers)
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	fmt.Println("numbers[:3] ==", numbers[:3])
	fmt.Println("numbers[:4] ==", numbers[:4])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	numbers3 := numbers[2:5]
	printSlice(numbers3)
}

func appendAndcopyFunctions() {
	// One can increase the capacity of a slice using the append() function. Using
	//  copy()function, the contents of a source slice are copied to a destination slice.

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers) * 2))

	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
