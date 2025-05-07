package main

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(
	c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{
		creature: c}}
}

func (n *NoBonusesModifier) Handle() {
	// nothing here!
}
