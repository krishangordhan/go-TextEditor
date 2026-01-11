package main

type Editor struct {
	buffer      *PieceTable
	cursor      *Cursor
	desiredCol  int
	fileManager *FileManager
}

func NewEditor(text string) *Editor {
	return &Editor{
		buffer:      NewPieceTable(text),
		cursor:      NewCursor(),
		desiredCol:  0,
		fileManager: NewFileManager(),
	}
}

func NewEditorFromFile(filePath string) (*Editor, error) {
	fm := NewFileManagerWithPath(filePath)
	content, err := fm.ReadFile()
	if err != nil {
		return nil, err
	}

	return &Editor{
		buffer:      NewPieceTable(content),
		cursor:      NewCursor(),
		desiredCol:  0,
		fileManager: fm,
	}, nil
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
	_, col := e.buffer.GetLineColumn(e.cursor.GetPosition())
	e.desiredCol = col
}

func (e *Editor) MoveCursorRight() {
	e.cursor.MoveRight(e.buffer.Length())
	_, col := e.buffer.GetLineColumn(e.cursor.GetPosition())
	e.desiredCol = col
}

func (e *Editor) MoveCursorUp() {
	pos := e.cursor.GetPosition()
	line, col := e.buffer.GetLineColumn(pos)

	if e.desiredCol == 0 {
		e.desiredCol = col
	}

	if line > 0 {
		targetLine := line - 1
		targetCol := e.desiredCol

		lineLength := e.buffer.GetLineLength(targetLine)
		if targetCol > lineLength {
			targetCol = lineLength
		}

		newPos := e.buffer.GetOffsetFromLineColumn(targetLine, targetCol)
		e.cursor.SetPosition(newPos)
	} else {
		e.cursor.SetPosition(0)
		e.desiredCol = 0
	}
}

func (e *Editor) MoveCursorDown() {
	pos := e.cursor.GetPosition()
	line, col := e.buffer.GetLineColumn(pos)

	if e.desiredCol == 0 {
		e.desiredCol = col
	}

	lineCount := e.buffer.GetLineCount()

	if line < lineCount-1 {
		targetLine := line + 1
		targetCol := e.desiredCol

		lineLength := e.buffer.GetLineLength(targetLine)
		if targetCol > lineLength {
			targetCol = lineLength
		}

		newPos := e.buffer.GetOffsetFromLineColumn(targetLine, targetCol)
		e.cursor.SetPosition(newPos)
	} else {
		e.cursor.SetPosition(e.buffer.Length())
		_, col := e.buffer.GetLineColumn(e.buffer.Length())
		e.desiredCol = col
	}
}

func (e *Editor) InsertAtCursor(text string) {
	pos := e.cursor.GetPosition()
	e.buffer.Insert(pos, text)
	e.cursor.SetPosition(pos + len([]rune(text)))
	e.fileManager.MarkDirty()
}

func (e *Editor) DeleteAtCursor(length int) {
	if length <= 0 {
		return
	}
	pos := e.cursor.GetPosition()
	e.buffer.Delete(pos, length)
	e.fileManager.MarkDirty()
}

func (e *Editor) Backspace() {
	pos := e.cursor.GetPosition()
	if pos > 0 {
		e.buffer.Delete(pos-1, 1)
		e.cursor.SetPosition(pos - 1)
		e.fileManager.MarkDirty()
	}
}

func (e *Editor) Delete() {
	pos := e.cursor.GetPosition()
	if pos < e.buffer.Length() {
		e.buffer.Delete(pos, 1)
		e.fileManager.MarkDirty()
	}
}

func (e *Editor) GetFileManager() *FileManager {
	return e.fileManager
}

func (e *Editor) Save() error {
	return e.fileManager.WriteFile(e.buffer.String())
}

func (e *Editor) SaveAs(filePath string) error {
	e.fileManager.SetFilePath(filePath)
	return e.fileManager.WriteFile(e.buffer.String())
}
