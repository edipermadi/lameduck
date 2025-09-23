package board

import (
	"fmt"
	"strings"

	"github.com/edipermadi/lameduck/pkg/core/pieces"
	"github.com/edipermadi/lameduck/pkg/core/positions"
)

type Grid [8][8]pieces.Piece

func (g *Grid) Set(position positions.Position, piece pieces.Piece) {
	file := int(position.File() - 'a')
	rank := int(position.Rank() - '1')
	g[rank][file] = piece
}

func (g *Grid) String() string {
	var sb strings.Builder
	sb.WriteString("┌───┬───┬───┬───┬───┬───┬───┬───┐\n")
	for file, files := range g {
		for _, piece := range files {
			sb.WriteString(fmt.Sprintf("│%*s ", 2, piece.Figure()))
		}
		sb.WriteString("│\n")
		if file < 7 {
			sb.WriteString("├───┼───┼───┼───┼───┼───┼───┼───┤\n")
		}
	}
	sb.WriteString("└───┴───┴───┴───┴───┴───┴───┴───┘\n")
	return sb.String()
}
