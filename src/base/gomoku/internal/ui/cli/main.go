package main

import (
	"fmt"
	"github.com/yourusername/gomoku/pkg/gomoku"
)

func main() {
	fmt.Println("Welcome to Gomoku!")
	game := gomoku.NewGame()
	game.Start()
}