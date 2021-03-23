/* Pointers
  Pointers in Go are easy and fun to learn. Some Go programming tasks are performed more
easily with pointers, and other tasks, such as call by reference, cannot be performed
without using pointers.
*/

package main

import "fmt"

func pointers() int {
	/* actual variable declaration */
	var a int = 20

	/* pointer variable declaration */
	var ip *int

	/* store address of a in pointer variable */
	ip = &a
	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	/* nil pointer */
	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)

	return *ip
}
