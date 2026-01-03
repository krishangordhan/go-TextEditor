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

func (pt *PieceTable) String() string {
	totalLength := 0
	for _, piece := range pt.pieces {
		totalLength += piece.length
	}

	result := make([]rune, 0, totalLength)
	for _, piece := range pt.pieces {
		var buffer []rune
		if piece.bufferType == Original {
			buffer = pt.original
		} else {
			buffer = pt.add
		}

		result = append(result, buffer[piece.start:piece.start+piece.length]...)
	}

	return string(result)
}

func (pt *PieceTable) Insert(offset int, text string) {
	if len(text) == 0 {
		return
	}

	addStart := len(pt.add)
	pt.add = append(pt.add, []rune(text)...)
	textLength := len([]rune(text))

	newPiece := Piece{
		bufferType: Add,
		start:      addStart,
		length:     textLength,
	}

	currentPos := 0
	for i, piece := range pt.pieces {
		pieceEnd := currentPos + piece.length

		if offset == currentPos {
			pt.pieces = append(pt.pieces[:i], append([]Piece{newPiece}, pt.pieces[i:]...)...)
			return
		}

		if offset > currentPos && offset < pieceEnd {
			splitAt := offset - currentPos

			leftPiece := Piece{
				bufferType: piece.bufferType,
				start:      piece.start,
				length:     splitAt,
			}

			rightPiece := Piece{
				bufferType: piece.bufferType,
				start:      piece.start + splitAt,
				length:     piece.length - splitAt,
			}

			pt.pieces = append(pt.pieces[:i], append([]Piece{leftPiece, newPiece, rightPiece}, pt.pieces[i+1:]...)...)
			return
		}

		currentPos = pieceEnd
	}

	pt.pieces = append(pt.pieces, newPiece)
}
