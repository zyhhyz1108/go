package main

import (
	"fmt"
	"os"
	"github.com/yourusername/gomoku/pkg/gomoku" // Adjust the import path as necessary
	"github.com/yourusername/gomoku/internal/ui/cli" // Adjust the import path as necessary
)

func main() {
	fmt.Println("Welcome to Gomoku!")
	game := gomoku.NewGame() // Initialize a new game
	cli.StartGame(game) // Start the command line interface for the game
	os.Exit(0)
}