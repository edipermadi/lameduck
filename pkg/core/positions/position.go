package positions

type Position int

const (
	A1 Position = iota
	B1 Position = iota
	C1 Position = iota
	D1 Position = iota
	E1 Position = iota
	F1 Position = iota
	G1 Position = iota
	H1 Position = iota

	A2 Position = iota
	B2 Position = iota
	C2 Position = iota
	D2 Position = iota
	E2 Position = iota
	F2 Position = iota
	G2 Position = iota
	H2 Position = iota

	A3 Position = iota
	B3 Position = iota
	C3 Position = iota
	D3 Position = iota
	E3 Position = iota
	F3 Position = iota
	G3 Position = iota
	H3 Position = iota

	A4 Position = iota
	B4 Position = iota
	C4 Position = iota
	D4 Position = iota
	E4 Position = iota
	F4 Position = iota
	G4 Position = iota
	H4 Position = iota

	A5 Position = iota
	B5 Position = iota
	C5 Position = iota
	D5 Position = iota
	E5 Position = iota
	F5 Position = iota
	G5 Position = iota
	H5 Position = iota

	A6 Position = iota
	B6 Position = iota
	C6 Position = iota
	D6 Position = iota
	E6 Position = iota
	F6 Position = iota
	G6 Position = iota
	H6 Position = iota

	A7 Position = iota
	B7 Position = iota
	C7 Position = iota
	D7 Position = iota
	E7 Position = iota
	F7 Position = iota
	G7 Position = iota
	H7 Position = iota

	A8 Position = iota
	B8 Position = iota
	C8 Position = iota
	D8 Position = iota
	E8 Position = iota
	F8 Position = iota
	G8 Position = iota
	H8 Position = iota
)

func New(file rune, rank rune) Position {
	rankInt := int(rank - '1')
	fileInt := int(file - 'a')
	return Position((rankInt * 8) + fileInt)
}

func (p Position) Mask() uint64 {
	return 1 << p
}

var names = map[Position]string{
	A1: "a1", B1: "b1", C1: "c1", D1: "d1", E1: "e1", F1: "f1", G1: "g1", H1: "h1",
	A2: "a2", B2: "b2", C2: "c2", D2: "d2", E2: "e2", F2: "f2", G2: "g2", H2: "h2",
	A3: "a3", B3: "b3", C3: "c3", D3: "d3", E3: "e3", F3: "f3", G3: "g3", H3: "h3",
	A4: "a4", B4: "b4", C4: "c4", D4: "d4", E4: "e4", F4: "f4", G4: "g4", H4: "h4",
	A5: "a5", B5: "b5", C5: "c5", D5: "d5", E5: "e5", F5: "f5", G5: "g5", H5: "h5",
	A6: "a6", B6: "b6", C6: "c6", D6: "d6", E6: "e6", F6: "f6", G6: "g6", H6: "h6",
	A7: "a7", B7: "b7", C7: "c7", D7: "d7", E7: "e7", F7: "f7", G7: "g7", H7: "h7",
	A8: "a8", B8: "b8", C8: "c8", D8: "d8", E8: "e8", F8: "f8", G8: "g8", H8: "h8",
}

func (p Position) String() string {
	return names[p]
}

func (p Position) File() rune {
	return 'a' + rune(p%8)
}

func (p Position) Rank() rune {
	return '1' + rune(p/8)
}

func AllPositions() []Position {
	return []Position{
		A1, B1, C1, D1, E1, F1, G1, H1,
		A2, B2, C2, D2, E2, F2, G2, H2,
		A3, B3, C3, D3, E3, F3, G3, H3,
		A4, B4, C4, D4, E4, F4, G4, H4,
		A5, B5, C5, D5, E5, F5, G5, H5,
		A6, B6, C6, D6, E6, F6, G6, H6,
		A7, B7, C7, D7, E7, F7, G7, H7,
		A8, B8, C8, D8, E8, F8, G8, H8,
	}
}
