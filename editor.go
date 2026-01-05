package main

type Editor struct {
	buffer *PieceTable
	cursor *Cursor
}

func NewEditor(text string) *Editor {
	return &Editor{
		buffer: NewPieceTable(text),
		cursor: NewCursor(),
	}
}

func (e *Editor) GetText() string {
	return e.buffer.String()
}

func (e *Editor) GetCursorPosition() int {
	return e.cursor.GetPosition()
}

func (e *Editor) SetCursorPosition(pos int) {
	if pos < 0 {
		pos = 0
	}
	length := e.buffer.Length()
	if pos > length {
		pos = length
	}
	e.cursor.SetPosition(pos)
}

func (e *Editor) MoveCursorLeft() {
	e.cursor.MoveLeft()
}

func (e *Editor) MoveCursorRight() {
	e.cursor.MoveRight(e.buffer.Length())
}

func (e *Editor) InsertAtCursor(text string) {
	pos := e.cursor.GetPosition()
	e.buffer.Insert(pos, text)
	e.cursor.SetPosition(pos + len([]rune(text)))
}

func (e *Editor) DeleteAtCursor(length int) {
	if length <= 0 {
		return
	}
	pos := e.cursor.GetPosition()
	e.buffer.Delete(pos, length)
}

func (e *Editor) Backspace() {
	pos := e.cursor.GetPosition()
	if pos > 0 {
		e.buffer.Delete(pos-1, 1)
		e.cursor.SetPosition(pos - 1)
	}
}

func (e *Editor) Delete() {
	pos := e.cursor.GetPosition()
	if pos < e.buffer.Length() {
		e.buffer.Delete(pos, 1)
	}
}
