package main

import (
	"fmt"
)

func main() {
	ba := BankAccount{100}
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)
	fmt.Println(ba)

	ba.Restore(m1)
	fmt.Println(ba)

	ba.Restore(m2)
	fmt.Println(ba)
}
