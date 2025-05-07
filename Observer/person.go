package main

// whenever a person catches a cold,
// a doctor must be called
type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}
