package gomoku

type Board struct {
	Size   int
	Cells  [][]rune
}

func NewBoard(size int) *Board {
	cells := make([][]rune, size)
	for i := range cells {
		cells[i] = make([]rune, size)
	}
	return &Board{
		Size:  size,
		Cells: cells,
	}
}

func (b *Board) Draw() {
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell == 0 {
				print(". ")
			} else {
				print(string(cell) + " ")
			}
		}
		println()
	}
}

func (b *Board) PlaceStone(x, y int, stone rune) bool {
	if x < 0 || x >= b.Size || y < 0 || y >= b.Size || b.Cells[x][y] != 0 {
		return false
	}
	b.Cells[x][y] = stone
	return true
}