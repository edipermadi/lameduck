package pieces

type Piece int

const (
	None        Piece = iota
	BlackBishop Piece = iota
	BlackKing   Piece = iota
	BlackKnight Piece = iota
	BlackPawn   Piece = iota
	BlackQueen  Piece = iota
	BlackRook   Piece = iota
	WhiteBishop Piece = iota
	WhiteKing   Piece = iota
	WhiteKnight Piece = iota
	WhitePawn   Piece = iota
	WhiteQueen  Piece = iota
	WhiteRook   Piece = iota
)

func AllPieces() []Piece {
	return []Piece{
		BlackBishop,
		BlackKing,
		BlackKnight,
		BlackPawn,
		BlackQueen,
		BlackRook,
		WhiteBishop,
		WhiteKing,
		WhiteKnight,
		WhitePawn,
		WhiteQueen,
		WhiteRook,
	}
}

var codes = map[Piece]string{
	BlackBishop: "b",
	BlackKing:   "k",
	BlackKnight: "n",
	BlackPawn:   "p",
	BlackQueen:  "q",
	BlackRook:   "r",
	WhiteBishop: "B",
	WhiteKing:   "K",
	WhiteKnight: "N",
	WhitePawn:   "P",
	WhiteQueen:  "Q",
	WhiteRook:   "R",
}

var figures = map[Piece]string{
	BlackBishop: "♝",
	BlackKing:   "♚",
	BlackKnight: "♞",
	BlackPawn:   "♟",
	BlackQueen:  "♛",
	BlackRook:   "♜",
	WhiteBishop: "♗",
	WhiteKing:   "♔",
	WhiteKnight: "♘",
	WhitePawn:   "♙",
	WhiteQueen:  "♕",
	WhiteRook:   "♖",
}

var names = map[Piece]string{
	BlackBishop: "BlackBishop",
	BlackKing:   "BlackKing",
	BlackKnight: "BlackKnight",
	BlackPawn:   "BlackPawn",
	BlackQueen:  "BlackQueen",
	BlackRook:   "BlackRook",
	WhiteBishop: "WhiteBishop",
	WhiteKing:   "WhiteKing",
	WhiteKnight: "WhiteKnight",
	WhitePawn:   "WhitePawn",
	WhiteQueen:  "WhiteQueen",
	WhiteRook:   "WhiteRook",
}

func (p Piece) Code() string {
	return codes[p]
}

func (p Piece) Figure() string {
	return figures[p]
}

func (p Piece) String() string {
	return names[p]
}
