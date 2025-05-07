package main

import "fmt"

type Person struct {
	name, position string
	salary         int
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dmitri").WorksAsA("dev").Salary(15000).Build()
	fmt.Println(*p)
}
