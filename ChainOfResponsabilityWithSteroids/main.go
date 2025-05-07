package main

import (
  "fmt"
  "sync"
)

type Argument int

const (
  Attack Argument = iota
  Defense
)

type Query struct {
  CreatureName string
  WhatToQuery Argument
  Value int
}

type Observer interface {
  Handle(*Query)
}

type Observable interface {
  Subscribe(o Observer)
  Unsubscribe(o Observer)
  Fire(q *Query)
}


func main() {
  game := &Game{sync.Map{}}
  goblin := NewCreature(game, "Strong Goblin", 2, 2)
  fmt.Println(goblin.String())

  {
    m := NewDoubleAttackModifier(game, goblin)
    fmt.Println(goblin.String())
    m.Close()
  }

  fmt.Println(goblin.String())
}