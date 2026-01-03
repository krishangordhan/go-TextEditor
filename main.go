package main

type BufferType int

const (
	Original BufferType = iota
	Add
)

type Piece struct {
	bufferType BufferType
	start      int
	length     int
}

type PieceTable struct {
	original []rune
	add      []rune
	pieces   []Piece
}

func NewPieceTable(text string) *PieceTable {
	pt := &PieceTable{
		original: []rune(text),
		add:      []rune{},
		pieces:   []Piece{},
	}

	if len(text) > 0 {
		pt.pieces = append(pt.pieces, Piece{
			bufferType: Original,
			start:      0,
			length:     len(pt.original),
		})
	}

	return pt
}
