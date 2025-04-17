package main

import "fmt"

func main() {
	b := PersonBuilder{}
	p := b.Called("Alex").WorksAsA("Developer").Build()
	fmt.Println(*p)
}
