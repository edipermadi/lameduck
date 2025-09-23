package board_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/edipermadi/lameduck/pkg/core/board"
	"github.com/edipermadi/lameduck/pkg/core/pieces"
	"github.com/stretchr/testify/require"
)

func TestBitBoard_Positions(t *testing.T) {
	type testCase struct {
		Piece             pieces.Piece
		ExpectedPositions []string
	}

	testCases := []testCase{
		{Piece: pieces.WhiteRook, ExpectedPositions: []string{"a1", "h1"}},
		{Piece: pieces.WhiteKnight, ExpectedPositions: []string{"b1", "g1"}},
		{Piece: pieces.WhiteBishop, ExpectedPositions: []string{"c1", "f1"}},
		{Piece: pieces.WhiteQueen, ExpectedPositions: []string{"d1"}},
		{Piece: pieces.WhiteKing, ExpectedPositions: []string{"e1"}},
		{Piece: pieces.WhitePawn, ExpectedPositions: []string{"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"}},
		{Piece: pieces.BlackRook, ExpectedPositions: []string{"a8", "h8"}},
		{Piece: pieces.BlackKnight, ExpectedPositions: []string{"b8", "g8"}},
		{Piece: pieces.BlackBishop, ExpectedPositions: []string{"c8", "f8"}},
		{Piece: pieces.BlackQueen, ExpectedPositions: []string{"d8"}},
		{Piece: pieces.BlackKing, ExpectedPositions: []string{"e8"}},
		{Piece: pieces.BlackPawn, ExpectedPositions: []string{"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_Positions", tc.Piece), func(t *testing.T) {
			b := board.NewBitBoard(tc.Piece)
			positions := make([]string, 0)
			for _, position := range b.Positions() {
				positions = append(positions, position.String())
			}
			require.Equal(t, tc.ExpectedPositions, positions)
			t.Logf("%s positions are [%s]", tc.Piece, strings.Join(positions, ", "))
		})
	}
}
