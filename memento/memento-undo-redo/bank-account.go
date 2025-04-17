package main

import "fmt"

type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{balance: balance}
	b.changes = append(b.changes, &Memento{balance})
	return b
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = %d, current = %d", b.balance, b.current)
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := Memento{b.balance}
	b.changes = append(b.changes, &m)
	b.current++
	fmt.Println("Deposited", amount, "Balance =", b.balance)
	return &m
}git

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}
