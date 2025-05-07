package main

// data common to all modifiers
type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (c *CreatureModifier) Handle(*Query) {
	// nothing here!
}
