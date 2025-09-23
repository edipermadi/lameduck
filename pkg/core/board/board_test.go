package board_test

import (
	"testing"

	"github.com/edipermadi/lameduck/pkg/core/board"
)

func TestInit(t *testing.T) {
	b := board.New()
	grid := b.Grid()
	view := grid.String()
	t.Logf("view: \n%s\n", view)
}
