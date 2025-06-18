package main

import "fmt"

type chess struct {
	turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

func (c *chess) Start() {
	fmt.Println("Starting a game of chess.")
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n",
		c.turn, c.currentPlayer)
	c.currentPlayer = (c.currentPlayer + 1) % 2
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}
