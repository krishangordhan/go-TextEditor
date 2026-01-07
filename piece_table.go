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

func (pt *PieceTable) Length() int {
	totalLength := 0
	for _, piece := range pt.pieces {
		totalLength += piece.length
	}
	return totalLength
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

func (pt *PieceTable) Delete(offset, length int) {
	if length == 0 {
		return
	}

	deleteEnd := offset + length
	currentPos := 0
	newPieces := []Piece{}

	for _, piece := range pt.pieces {
		pieceStart := currentPos
		pieceEnd := currentPos + piece.length

		if pieceEnd <= offset {
			newPieces = append(newPieces, piece)
			currentPos = pieceEnd
			continue
		}

		if pieceStart >= deleteEnd {
			newPieces = append(newPieces, piece)
			currentPos = pieceEnd
			continue
		}

		overlapStart := max(pieceStart, offset)
		overlapEnd := min(pieceEnd, deleteEnd)

		if overlapStart > pieceStart {
			leftLength := overlapStart - pieceStart
			newPieces = append(newPieces, Piece{
				bufferType: piece.bufferType,
				start:      piece.start,
				length:     leftLength,
			})
		}

		if overlapEnd < pieceEnd {
			rightStart := piece.start + (overlapEnd - pieceStart)
			rightLength := pieceEnd - overlapEnd
			newPieces = append(newPieces, Piece{
				bufferType: piece.bufferType,
				start:      rightStart,
				length:     rightLength,
			})
		}

		currentPos = pieceEnd
	}

	pt.pieces = newPieces
}

func (pt *PieceTable) GetLineColumn(offset int) (line, col int) {
	text := []rune(pt.String())
	if offset > len(text) {
		offset = len(text)
	}

	line = 0
	col = 0

	for i := 0; i < offset; i++ {
		if text[i] == '\n' {
			line++
			col = 0
		} else {
			col++
		}
	}

	return line, col
}

func (pt *PieceTable) GetOffsetFromLineColumn(targetLine, targetCol int) int {
	text := []rune(pt.String())
	line := 0
	col := 0

	for i := 0; i <= len(text); i++ {
		if line == targetLine && col == targetCol {
			return i
		}

		if i == len(text) {
			break
		}

		if text[i] == '\n' {
			if line == targetLine {
				return i
			}
			line++
			col = 0
		} else {
			col++
		}
	}

	return len(text)
}

func (pt *PieceTable) GetLineLength(lineNum int) int {
	text := []rune(pt.String())
	line := 0
	lineStart := 0

	for i := 0; i <= len(text); i++ {
		if line == lineNum {
			if i == len(text) || text[i] == '\n' {
				return i - lineStart
			}
		}

		if i < len(text) && text[i] == '\n' {
			if line == lineNum {
				return i - lineStart
			}
			line++
			lineStart = i + 1
		}
	}

	return 0
}

func (pt *PieceTable) GetLineCount() int {
	text := []rune(pt.String())
	if len(text) == 0 {
		return 1
	}

	lines := 1
	for _, r := range text {
		if r == '\n' {
			lines++
		}
	}
	return lines
}
