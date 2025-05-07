package main

import (
	"fmt"
)

type Observer interface {
	Notify(data interface{})
}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s",
		data.(string))
}

func main() {

	ds := &DoctorService{}

	p := NewPerson("john")
	y := NewPerson("Lizz")

	p.Subscribe(ds)
	y.Subscribe(ds)

	p.CatchACold()
	y.CatchACold()
}
