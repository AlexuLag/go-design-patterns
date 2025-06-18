package main

type Game interface {
	Start()
	HaveWinner() bool
	TakeTurn()
	WinningPlayer() int
}
