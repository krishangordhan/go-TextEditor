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
