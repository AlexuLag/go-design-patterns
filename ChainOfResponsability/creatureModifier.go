package main

type CreatureModifier struct {
	creature *Creature
	next     Modifier // singly linked list
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}
