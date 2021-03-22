/* Strings
  Strings, which are widely used in Go programming, are a readonly slice of bytes.
In the Go programming language, strings are slices. The Go platform provides various
libraries to manipulate strings.

	# unicode
	# regexp
	# strings

*/

package main

import "strings"

func stringsoperation() {
	creatingStrings()
	stringLength()
	concatenatingStrings()
}

func creatingStrings() []byte {
	var greeting = "Hello world!"
	var slice []byte

	for i := 0; i < len(greeting); i++ {
		slice = append(slice, greeting[i])
	}

	return slice
}

func stringLength() int {
	var greeting = "Hello world!"
	return len(greeting)
}

func concatenatingStrings() string {
	greetings := []string{"Hello", "world!"}
	return strings.Join(greetings, " ")
}
