package board_test

import (
	"testing"

	"github.com/MelleKoning/fourinarow/pkg/board"
	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	brd := board.NewEmptyBoard()

	assert.Equal(t, board.Empty, brd.Pos[0][0])

}
