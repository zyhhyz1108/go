package gomoku

import (
	"testing"
)

func TestInitializeBoard(t *testing.T) {
	board := NewBoard()
	if len(board) != 15 || len(board[0]) != 15 {
		t.Errorf("Expected board size 15x15, got %dx%d", len(board), len(board[0]))
	}
}

func TestPlacePiece(t *testing.T) {
	board := NewBoard()
	err := board.PlacePiece(7, 7, 'X')
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if board[7][7] != 'X' {
		t.Errorf("Expected piece 'X' at (7, 7), got %c", board[7][7])
	}
}

func TestPlacePieceOutOfBounds(t *testing.T) {
	board := NewBoard()
	err := board.PlacePiece(15, 15, 'X')
	if err == nil {
		t.Errorf("Expected error for out of bounds placement, got none")
	}
}

func TestCheckWin(t *testing.T) {
	board := NewBoard()
	board.PlacePiece(7, 7, 'X')
	board.PlacePiece(7, 8, 'X')
	board.PlacePiece(7, 9, 'X')
	board.PlacePiece(7, 10, 'X')
	board.PlacePiece(7, 11, 'X')

	if !board.CheckWin(7, 11, 'X') {
		t.Errorf("Expected 'X' to win, but did not detect win")
	}
}