package main

import (
	"fmt"
)

type Observer interface {
	Notify(data interface{})
}

type PropertyChanged struct {
	Name  string
	Value interface{}
}

func main() {
	p := NewPerson(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
