/* Structures
  Go arrays allow you to define variables that can hold several data items of the
same kind. Structure is another user-defined data type available in Go programming,
which allows you to combine data items of different kinds. Structures are used to
represent a record. Suppose you want to keep track of the books in a library. You
might want to track the following attributes of each book −

	Title
	Author
	Subject
	Book ID

  In such a scenario, structures are highly useful.

  To define a structure, you must use type and struct statements. The struct statement
defines a new data type, with multiple members for your program. The type statement
binds a name with the type which is struct in our case. The format of the struct
statement is as follows −

	type struct_variable_type struct {
	  member definition;
	  member definition;
	  ...
	  member definition;
	}

  Once a structure type is defined, it can be used to declare variables of that type
using the following syntax.

	variable_name := structure_variable_type {value1, value2...valuen}
*/

package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func structures() {
	accessingStructureMembers()
	structuresAsFunctionArguments()
}

func accessingStructureMembers() {
	var Book1 Books
	var Book2 Books

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	/* print Book1 info */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* print Book2 info */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}

func structuresAsFunctionArguments() {
	var Book1 Books
	var Book2 Books

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	/* print Book1 info */
	printBook(Book1)
	printBookbyPointers(&Book1)

	/* print Book2 info */
	printBook(Book2)
	printBookbyPointers(&Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBookbyPointers(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
