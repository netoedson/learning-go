// The package Clause
package main

// The Import Declaration
import "fmt" // Internal tool - Reformats source code to adhere to the standard

// Source Body
// variable declaration
var name string

type person struct {
	name     string
	lastname string
	age      int
}

func main() {
	fmt.Println("Hello World!")

	// the type is assigned based on the literal type value
	value := 10

	// Printing the type and the value
	fmt.Printf("%T - %v\n", value, value)

	// This will print a zero-value. For string is ""
	fmt.Printf("%T - %v\n", name, name)

	// Declaring a slice.
	slice := []int{1, 2, 3}
	// Declaring an array
	array := [3]int{4, 3, 2}

	// Range is use to know the elements inside the slice
	// _ is a blank identifier (it avoids error about the use of declared identifies)
	for _, v := range slice {
		fmt.Println(v)
	}

	for index := 0; index < len(array); index++ {
		fmt.Println(array[index] * 2)
	}

	amap := map[string]int{
		"good":    0,
		"neutral": 1,
		"evil":    2,
	}

	simpleFunc(amap)

	person1 := person{"John", "Doe", 26}

	person2 := person{
		name:     "Dow",
		lastname: "John",
		age:      62,
	}

	fmt.Println(person1, "\n", person2)
}

func simpleFunc(value map[string]int) {
	fmt.Println(value)
}
