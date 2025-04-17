package main

import "fmt"

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	ba.Undo()
	fmt.Println("unod 1", ba)
	ba.Undo()
	fmt.Println("unod 2", ba)
	ba.Redo()
	fmt.Println("redo", ba)
}
