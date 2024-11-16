package main

import "fmt"

func calculateArea(width, height float64) float64 {
	return width * height
}
func getPersonInfo() (string, int) {
	return "Ayub Subagiya", 21
}

func main() {
	fmt.Println("Hello, World!")

	var name string = "Ayub Subagiya"
	age := 21
	var isStudent bool = true

	fmt.Printf("Name: %s\nAge: %d\nIs Student: %t\n", name, age, isStudent)

	// calculate area
	area := calculateArea(20, 10)
	fmt.Printf("Area: %.2f\n", area)

	// get person info
	name, age = getPersonInfo()
	fmt.Printf("Name: %s\nAge: %d\n", name, age)
}
