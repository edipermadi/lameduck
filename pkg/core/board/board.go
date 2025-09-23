package board

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"

	"github.com/edipermadi/lameduck/pkg/core/pieces"
)

type Board [12]BitBoard

func New() Board {
	var board Board
	for _, piece := range pieces.AllPieces() {
		board[piece-1] = NewBitBoard(piece)
	}
	return board
}

func (c Board) ToBase64() string {
	return base64.StdEncoding.EncodeToString(c.Bytes())
}

func (c Board) Bytes() []byte {
	var buffer bytes.Buffer
	for _, v := range c {
		_ = binary.Write(&buffer, binary.BigEndian, v)
	}
	return buffer.Bytes()
}

func (c Board) Movements() []Board {
	return nil
}

func (c Board) Grid() Grid {
	var grid Grid

	for _, piece := range pieces.AllPieces() {
		for _, position := range c[piece-1].Positions() {
			grid.Set(position, piece)
		}
	}
	
	return grid
}
