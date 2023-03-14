package main

import "fmt"

// Global Variables
var (
	s string  = "SampleString"
	i int     = 100
	f float64 = 10.123
	b bool    = true
)

func GetVariables() {
	fmt.Printf("%T -> %s\n", s, s)
	fmt.Printf("%T -> %d\n", i, i)
	fmt.Printf("%T -> %f\n", f, f)
	fmt.Printf("%T\n", b)
}

type data struct {
	name string
}

type student struct {
	id   int
	name string
}

func (d data) GetData() {
	fmt.Println(d.name)
}

// Caller function
func (s student) DisplayStudent() {
	fmt.Printf("Student ID = %d\n", s.id)
	fmt.Printf("Student Name = %s\n", s.name)
}
