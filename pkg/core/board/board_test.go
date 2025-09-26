package board_test

import (
	"testing"

	"github.com/edipermadi/lameduck/pkg/core/board"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	b := board.New()
	grid := b.Grid()
	view := grid.String()
	t.Logf("view: \n%s\n", view)
}

func TestFEN(t *testing.T) {
	b := board.New()
	fen := b.FEN()
	require.Equal(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", fen)
	t.Logf("fen: %s", fen)
}
