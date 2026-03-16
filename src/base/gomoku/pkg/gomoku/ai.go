package gomoku

import "math/rand"

// AI represents a simple AI for the Gomoku game.
type AI struct {
	board *Board
}

// NewAI creates a new AI instance.
func NewAI(board *Board) *AI {
	return &AI{board: board}
}

// MakeMove makes a move for the AI on the board.
func (ai *AI) MakeMove() (int, int) {
	// Simple AI logic: randomly choose an empty spot on the board
	var x, y int
	for {
		x = rand.Intn(ai.board.Size)
		y = rand.Intn(ai.board.Size)
		if ai.board.IsEmpty(x, y) {
			break
		}
	}
	return x, y
}