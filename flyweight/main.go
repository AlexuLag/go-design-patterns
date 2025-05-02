package main

import (
	"fmt"
)

var allNames []string

func main() {

	john2 := NewUser("John Doe")
	jane2 := NewUser("Jane Doe")

	fmt.Println(john2.FullName())
	fmt.Println(jane2.FullName())

}
