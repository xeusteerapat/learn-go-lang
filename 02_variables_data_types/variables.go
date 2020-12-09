package main

import "fmt"

func main() {
	var number = 3
	var name string = "Teerapat Prommarak"
	name = "Xeus" // re-assigned

	// constant, cannot re-assign
	const bodyTemp = 36.5
	// bodyTemp = 37.7 error

	var isOld = false

	var age uint16 = 34

	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(isOld)
	fmt.Println(age)

	// print formatting
	fmt.Printf("Type: %T Value: %v\n", isOld, isOld)
	fmt.Printf("Type: %T Value: %v\n", age, age)
	fmt.Printf("Type: %T Value: %v\n", bodyTemp, bodyTemp)
	fmt.Printf("Type: %T Value: %v\n", name, name)
	fmt.Printf("Type: %T Value: %v\n", number, number)
}
