package main

import "fmt"

type DoubleAttackModifier struct {
	CreatureModifier
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name,
		"attack...")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{
		creature: c}}
}
