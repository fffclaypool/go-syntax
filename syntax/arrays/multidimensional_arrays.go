/* Multidimensional Arrays
  A two-dimensional array is the simplest form of a multidimensional array. A two-dimensional
array is, in essence, a list of one-dimensional arrays. To declare a two-dimensional integer
array of size [x, y], you would write something as follows −

	var arrayName [ x ][ y ] variable_type

  Where variable_type can be any valid Go data type and arrayName will be a valid Go
identifier. A two-dimensional array can be think as a table which will have x number of rows
and y number of columns.
*/

package main

func main() {
	accessingTwodimensionalArrayElements()
}

func accessingTwodimensionalArrayElements() [5][2]int {
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var b = [5][2]int{}
	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			b[i][j] = a[i][j]
		}
	}

	return b
}
