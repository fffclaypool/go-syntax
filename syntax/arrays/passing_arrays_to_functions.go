/* Passing arrays to functions
  If you want to pass a single-dimension array as an argument in a function, you
would have to declare function formal parameter in one of following two ways and
all two declaration methods produce similar results because each tells the compiler
that an integer array is going to be received. Similar way you can pass multi-dimensional
array as formal parameters.
*/

package main

func passingarraystofunctions() {
	basicOperations()
}

func basicOperations() {
	var balance = []int{1000, 2, 3, 17, 50}
	getAverage(balance, 5)
}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)
	return avg
}
