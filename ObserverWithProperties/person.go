package main

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{Observable{make([]Observer, 0)}, age}
}

func (p *Person) Age() int { return p.age }

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChanged{"Age", p.age})
}
