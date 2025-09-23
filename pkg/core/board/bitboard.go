package board

import (
	"github.com/edipermadi/lameduck/pkg/core/pieces"
	"github.com/edipermadi/lameduck/pkg/core/positions"
)

type BitBoard uint64

func NewBitBoard(piece pieces.Piece) BitBoard {
	switch piece {
	case pieces.WhiteRook:
		return BitBoard(positions.A1.Mask() | positions.H1.Mask())
	case pieces.WhiteKnight:
		return BitBoard(positions.B1.Mask() | positions.G1.Mask())
	case pieces.WhiteBishop:
		return BitBoard(positions.C1.Mask() | positions.F1.Mask())
	case pieces.WhiteQueen:
		return BitBoard(positions.D1.Mask())
	case pieces.WhitePawn:
		return BitBoard(positions.A2.Mask() |
			positions.B2.Mask() |
			positions.C2.Mask() |
			positions.D2.Mask() |
			positions.E2.Mask() |
			positions.F2.Mask() |
			positions.G2.Mask() |
			positions.H2.Mask())
	case pieces.WhiteKing:
		return BitBoard(positions.E1.Mask())
	case pieces.BlackRook:
		return BitBoard(positions.A8.Mask() | positions.H8.Mask())
	case pieces.BlackKnight:
		return BitBoard(positions.B8.Mask() | positions.G8.Mask())
	case pieces.BlackBishop:
		return BitBoard(positions.C8.Mask() | positions.F8.Mask())
	case pieces.BlackQueen:
		return BitBoard(positions.D8.Mask())
	case pieces.BlackKing:
		return BitBoard(positions.E8.Mask())
	case pieces.BlackPawn:
		return BitBoard(positions.A7.Mask() |
			positions.B7.Mask() |
			positions.C7.Mask() |
			positions.D7.Mask() |
			positions.E7.Mask() |
			positions.F7.Mask() |
			positions.G7.Mask() |
			positions.H7.Mask())
	default:
		return 0
	}
}

func (b BitBoard) Positions() []positions.Position {
	value := uint64(b)
	result := make([]positions.Position, 0)
	for _, position := range positions.AllPositions() {
		if value&position.Mask() > 0 {
			result = append(result, position)
		}
	}
	return result
}
