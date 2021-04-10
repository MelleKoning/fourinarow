package board

// this package should resemble the board in memory
// we will at first use a very basic implementation
// and later see if we would go for more optimized
// storage to safe memory and speed

// Visualizing the board:
// rows
//  5    | _ | _ | _ | _ | _ | _ | _ |
//  4    | _ | _ | _ | _ | _ | _ | _ |
//  3    | _ | _ | _ | _ | _ | _ | _ |
//  2    | _ | _ | _ | _ | _ | _ | _ |
//  1    | _ | _ | _ | _ | _ | _ | _ |
//  0    | _ | _ | _ | _ | _ | _ | _ |
// col     1   2   3   4   5   6   7
type FieldVal int

const (
	Empty FieldVal = iota
	Red
	Blue
)

const rowCount = 6
const colCount = 7

type Board struct {
	Pos [rowCount][colCount]FieldVal
}

func NewEmptyBoard() *Board {
	b := &Board{}
	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			b.Pos[r][c] = Empty
		}
	}
	return b
}
