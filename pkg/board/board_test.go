package board_test

import (
	"testing"

	"github.com/MelleKoning/fourinarow/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	brd := board.NewEmptyBoard()

	assert.Equal(t, board.Empty, brd.Row[0].Col[0])
}

func TestMakeCenterMove(t *testing.T) {
	brd := board.NewEmptyBoard()
	// default first move goes to Red
	err := brd.MoveInColumn(3)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, board.Red, brd.Row[0].Col[3])
	assert.Equal(t, board.Yellow, brd.Turn) // check that Yellow has the move now
}

func TestSevenMovesOverFlowsColumn(t *testing.T) {
	brd := board.NewEmptyBoard()

	var err error
	for move := 0; move <= 6; move++ {
		err = brd.MoveInColumn(3)
	}
	assert.NotNil(t, err)
}
