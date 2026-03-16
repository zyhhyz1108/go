package gomoku

type Game struct {
	Board      [][]rune
	CurrentTurn rune
}

const (
	Empty   rune = '.'
	Player1 rune = 'X'
	Player2 rune = 'O'
)

func NewGame(size int) *Game {
	board := make([][]rune, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = Empty
		}
	}
	return &Game{
		Board:      board,
		CurrentTurn: Player1,
	}
}

func (g *Game) PlayMove(x, y int) bool {
	if g.Board[x][y] == Empty {
		g.Board[x][y] = g.CurrentTurn
		g.switchTurn()
		return true
	}
	return false
}

func (g *Game) switchTurn() {
	if g.CurrentTurn == Player1 {
		g.CurrentTurn = Player2
	} else {
		g.CurrentTurn = Player1
	}
}

func (g *Game) CheckWin(x, y int) bool {
	return g.checkDirection(x, y, 1, 0) || // Horizontal
		g.checkDirection(x, y, 0, 1) || // Vertical
		g.checkDirection(x, y, 1, 1) || // Diagonal \
		g.checkDirection(x, y, 1, -1)   // Diagonal /
}

func (g *Game) checkDirection(x, y, dx, dy int) bool {
	count := 1
	count += g.countInDirection(x, y, dx, dy)
	count += g.countInDirection(x, y, -dx, -dy)
	return count >= 5
}

func (g *Game) countInDirection(x, y, dx, dy int) int {
	count := 0
	player := g.Board[x][y]
	for i := 1; i < 5; i++ {
		newX := x + i*dx
		newY := y + i*dy
		if newX < 0 || newY < 0 || newX >= len(g.Board) || newY >= len(g.Board) || g.Board[newX][newY] != player {
			break
		}
		count++
	}
	return count
}