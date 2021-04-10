package board

import "fmt"

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
	Yellow
)

const rowCount = 6
const colCount = 7

type Col struct {
	Col [colCount]FieldVal
}
type Board struct {
	Row  [rowCount]*Col
	Turn FieldVal // who's turn it is
}

func NewEmptyBoard() *Board {
	b := &Board{}
	for r := 0; r < rowCount; r++ {
		b.Row[r] = &Col{}
		for c := 0; c < colCount; c++ {
			b.Row[r].Col[c] = Empty
		}
	}
	b.Turn = Red
	return b
}

func (b *Board) MoveInColumn(col int) error {
	if b.Row[5].Col[col] != Empty {
		return fmt.Errorf("Can't move in that column")
	}

	foundrow := 0
	for row := 0; row <= 5; row++ {
		if b.Row[row].Col[col] == Empty {
			foundrow = row
			break
		}
	}

	b.Row[foundrow].Col[col] = b.Turn
	b.SwitchWhosToMove()
	return nil
}

func (b *Board) SwitchWhosToMove() {
	if b.Turn == Red {
		b.Turn = Yellow
		return
	}
	b.Turn = Red
}
