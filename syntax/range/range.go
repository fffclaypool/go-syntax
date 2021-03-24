/* Range
  The range keyword is used in for loop to iterate over items of an array,
slice, channel or map. With array and slices, it returns the index of the
item as integer. With maps, it returns the key of the next key-value pair.
Range either returns one value or two. If only one value is used on the left
of a range expression, it is the 1st value in the following table.
*/

package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}
