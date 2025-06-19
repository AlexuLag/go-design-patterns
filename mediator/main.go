package main

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	luis := NewPerson("luis")

	room.Join(john)
	room.Join(luis)

	john.Say("hi room")
	luis.Say("oh, hey john")

	simeon := NewPerson("Simeon")
	room.Join(simeon)
	simeon.Say("hi everyone!")

	luis.PrivateMessage("Simeon", "glad you could join us!")
}
