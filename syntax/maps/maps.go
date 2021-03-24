/* Maps
  Go provides another important data type named map which maps unique keys to
values. A key is an object that you use to retrieve a value at a later date.
Given a key and a value, you can store the value in a Map object. After the
value is stored, you can retrieve it by using its key.
*/

package main

import "fmt"

func main() {
	definingMaps()
	deleteFunction()
}

func definingMaps() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}

func deleteFunction() {
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("Original map")

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")

	fmt.Println("Entry for France is deleted")
	fmt.Println("Updated map")

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}
