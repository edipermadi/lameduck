package board

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/edipermadi/lameduck/pkg/core/pieces"
	"github.com/edipermadi/lameduck/pkg/core/positions"
)

type Board [12]BitBoard

func New() Board {
	var board Board
	for _, piece := range pieces.AllPieces() {
		board[piece-1] = NewBitBoard(piece)
	}
	return board
}

func (b Board) ToBase64() string {
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func (b Board) Bytes() []byte {
	var buffer bytes.Buffer
	for _, v := range b {
		_ = binary.Write(&buffer, binary.BigEndian, v)
	}
	return buffer.Bytes()
}

func (b Board) Movements() []Board {
	return nil
}

func (b Board) Pieces() []pieces.Piece {
	allPositions := positions.AllPositions()
	allPieces := pieces.AllPieces()
	result := make([]pieces.Piece, len(allPositions))
	for _, position := range allPositions {
		for _, piece := range allPieces {
			if bitmask := b[piece-1]; bitmask.Has(position) {
				result[position] = piece
				break
			}
		}
	}
	return result
}

func (b Board) PiecesMap() map[positions.Position]pieces.Piece {
	allPositions := positions.AllPositions()
	allPieces := pieces.AllPieces()
	result := make(map[positions.Position]pieces.Piece)
	for _, position := range allPositions {
		for _, piece := range allPieces {
			if bitmask := b[piece-1]; bitmask.Has(position) {
				result[position] = piece
				break
			}
		}
	}
	return result
}

func (b Board) Grid() Grid {
	var grid Grid

	for _, piece := range pieces.AllPieces() {
		for _, position := range b[piece-1].Positions() {
			grid.Set(position, piece)
		}
	}

	return grid
}

func (b Board) FEN() string {
	var sb strings.Builder
	g := b.Grid()
	for i := 0; i < 8; i++ {
		rank := rune('8' - i)
		emptyCtr := 0
		for j := 0; j < 8; j++ {
			file := rune('a' + j)
			position := positions.New(file, rank)
			piece := g.Get(position)
			if piece == pieces.None {
				emptyCtr += 1
			} else {
				if emptyCtr > 0 {
					sb.WriteString(strconv.Itoa(emptyCtr))
					emptyCtr = 0
				}

				sb.WriteString(piece.Code())
			}
		}
		if emptyCtr > 0 {
			sb.WriteString(strconv.Itoa(emptyCtr))
		}
		if i < 7 {
			sb.WriteRune('/')
		}
	}

	return sb.String()
}
