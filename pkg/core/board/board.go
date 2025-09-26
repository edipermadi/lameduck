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

func (c Board) FEN() string {
	var sb strings.Builder
	g := c.Grid()
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
