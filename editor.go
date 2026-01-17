package main

type Editor struct {
	buffer      *PieceTable
	cursor      *Cursor
	desiredCol  int
	fileManager *FileManager
	undoStack   []Command
	redoStack   []Command
}

func NewEditor(text string) *Editor {
	return &Editor{
		buffer:      NewPieceTable(text),
		cursor:      NewCursor(),
		desiredCol:  0,
		fileManager: NewFileManager(),
		undoStack:   make([]Command, 0),
		redoStack:   make([]Command, 0),
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
		undoStack:   make([]Command, 0),
		redoStack:   make([]Command, 0),
	}, nil
}

func (e *Editor) GetText() string {
	return e.buffer.String()
}

func (e *Editor) GetBuffer() *PieceTable {
	return e.buffer
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
	cmd := NewInsertCommand(e.buffer, e.cursor, text, pos)
	cmd.Execute()
	e.undoStack = append(e.undoStack, cmd)
	e.redoStack = make([]Command, 0)
	e.fileManager.MarkDirty()
}

func (e *Editor) DeleteAtCursor(length int) {
	if length <= 0 {
		return
	}
	pos := e.cursor.GetPosition()
	cmd := NewDeleteCommand(e.buffer, e.cursor, pos, length)
	cmd.Execute()
	e.undoStack = append(e.undoStack, cmd)
	e.redoStack = make([]Command, 0)
	e.fileManager.MarkDirty()
}

func (e *Editor) Backspace() {
	pos := e.cursor.GetPosition()
	if pos > 0 {
		cmd := NewDeleteCommand(e.buffer, e.cursor, pos-1, 1)
		cmd.Execute()
		e.undoStack = append(e.undoStack, cmd)
		e.redoStack = make([]Command, 0)
		e.fileManager.MarkDirty()
	}
}

func (e *Editor) Delete() {
	pos := e.cursor.GetPosition()
	if pos < e.buffer.Length() {
		cmd := NewDeleteCommand(e.buffer, e.cursor, pos, 1)
		cmd.Execute()
		e.undoStack = append(e.undoStack, cmd)
		e.redoStack = make([]Command, 0)
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

func (e *Editor) Undo() {
	if len(e.undoStack) == 0 {
		return
	}

	lastIndex := len(e.undoStack) - 1
	cmd := e.undoStack[lastIndex]
	e.undoStack = e.undoStack[:lastIndex]

	cmd.Undo()

	e.redoStack = append(e.redoStack, cmd)
}
