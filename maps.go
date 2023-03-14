package main

import "fmt"

//Declare and assign a map
var m2 = map[string]int{"foo": 1, "bar": 2}

func DisplayMaps() {

	//Declare a map
	m := make(map[string]int)

	m["One"] = 1
	m["two"] = 2

	fmt.Println(m)
	fmt.Println(m2)
}

// Function to create a map and initialize with student data
// returns map[string]int reference
func SetMap() map[string]int {
	m := map[string]int{
		"student_1": 1,
		"student_2": 2,
		"student_3": 3,
		"student_4": 4,
		"student_5": 5,
	}
	return m
}
