package main

import "fmt"

func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}
