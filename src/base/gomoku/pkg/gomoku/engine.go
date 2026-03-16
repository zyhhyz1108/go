package gomoku

type GameEngine struct {
    Board      *Board
    CurrentPlayer int
    GameOver   bool
}

func NewGameEngine(size int) *GameEngine {
    return &GameEngine{
        Board: NewBoard(size),
        CurrentPlayer: 1,
        GameOver: false,
    }
}

func (g *GameEngine) PlayMove(x, y int) bool {
    if g.GameOver {
        return false
    }
    if g.Board.PlacePiece(x, y, g.CurrentPlayer) {
        if g.Board.CheckWin(g.CurrentPlayer) {
            g.GameOver = true
        } else {
            g.CurrentPlayer = 3 - g.CurrentPlayer // Switch between player 1 and 2
        }
        return true
    }
    return false
}

func (g *GameEngine) Reset() {
    g.Board.Reset()
    g.CurrentPlayer = 1
    g.GameOver = false
}