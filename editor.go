package main

type Editor struct {
	buffer      *PieceTable
	cursor      *Cursor
	desiredCol  int
	fileManager *FileManager
	clipboard   Clipboard
	undoStack   []Command
	redoStack   []Command
}

func NewEditor(text string) *Editor {
	return &Editor{
		buffer:      NewPieceTable(text),
		cursor:      NewCursor(),
		desiredCol:  0,
		fileManager: NewFileManager(),
		clipboard:   NewClipboardManager(),
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
		clipboard:   NewClipboardManager(),
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
	e.moveCursorLeft(false)
}

func (e *Editor) MoveCursorLeftWithSelection() {
	e.moveCursorLeft(true)
}

func (e *Editor) MoveCursorRight() {
	e.moveCursorRight(false)
}

func (e *Editor) MoveCursorRightWithSelection() {
	e.moveCursorRight(true)
}

func (e *Editor) moveCursorLeft(withSelection bool) {
	if withSelection && !e.cursor.HasSelection() {
		e.cursor.StartSelection()
	}
	e.cursor.MoveLeft()
	_, col := e.buffer.GetLineColumn(e.cursor.GetPosition())
	e.desiredCol = col
}

func (e *Editor) moveCursorRight(withSelection bool) {
	if withSelection && !e.cursor.HasSelection() {
		e.cursor.StartSelection()
	}
	e.cursor.MoveRight(e.buffer.Length())
	_, col := e.buffer.GetLineColumn(e.cursor.GetPosition())
	e.desiredCol = col
}

func (e *Editor) MoveCursorUp() {
	e.moveCursorUp(false)
}

func (e *Editor) MoveCursorUpWithSelection() {
	e.moveCursorUp(true)
}

func (e *Editor) moveCursorUp(withSelection bool) {
	e.moveCursorVertical(-1, withSelection)
}

func (e *Editor) MoveCursorDown() {
	e.moveCursorDown(false)
}

func (e *Editor) MoveCursorDownWithSelection() {
	e.moveCursorDown(true)
}

func (e *Editor) moveCursorDown(withSelection bool) {
	e.moveCursorVertical(1, withSelection)
}

func (e *Editor) moveCursorVertical(direction int, withSelection bool) {
	if withSelection && !e.cursor.HasSelection() {
		e.cursor.StartSelection()
	}

	pos := e.cursor.GetPosition()
	line, col := e.buffer.GetLineColumn(pos)

	if e.desiredCol == 0 {
		e.desiredCol = col
	}

	lineCount := e.buffer.GetLineCount()
	targetLine := line + direction

	if targetLine < 0 {
		e.cursor.SetPosition(0)
		e.desiredCol = 0
		return
	}
	if targetLine >= lineCount {
		e.cursor.SetPosition(e.buffer.Length())
		_, col := e.buffer.GetLineColumn(e.buffer.Length())
		e.desiredCol = col
		return
	}

	targetCol := e.desiredCol
	lineLength := e.buffer.GetLineLength(targetLine)
	if targetCol > lineLength {
		targetCol = lineLength
	}

	newPos := e.buffer.GetOffsetFromLineColumn(targetLine, targetCol)
	e.cursor.SetPosition(newPos)
}

func (e *Editor) HasSelection() bool {
	return e.cursor.HasSelection()
}

func (e *Editor) GetSelection() (int, int) {
	return e.cursor.GetSelection()
}

func (e *Editor) ClearSelection() {
	e.cursor.ClearSelection()
}

func (e *Editor) Copy() error {
	if !e.cursor.HasSelection() {
		return nil
	}

	start, end := e.cursor.GetSelection()
	text := e.buffer.Substring(start, end)
	return e.clipboard.Copy(text)
}

func (e *Editor) InsertAtCursor(text string) {
	if e.cursor.HasSelection() {
		start, end := e.cursor.GetSelection()
		length := end - start
		deleteCmd := NewDeleteCommand(e.buffer, e.cursor, start, length)
		e.executeCommand(deleteCmd)
		e.cursor.ClearSelection()
	}

	pos := e.cursor.GetPosition()
	cmd := NewInsertCommand(e.buffer, e.cursor, text, pos)
	e.executeCommand(cmd)
}

func (e *Editor) DeleteAtCursor(length int) {
	if length <= 0 {
		return
	}
	pos := e.cursor.GetPosition()
	cmd := NewDeleteCommand(e.buffer, e.cursor, pos, length)
	e.executeCommand(cmd)
}

func (e *Editor) Backspace() {
	if e.cursor.HasSelection() {
		e.deleteSelection()
		return
	}

	pos := e.cursor.GetPosition()
	if pos > 0 {
		cmd := NewDeleteCommand(e.buffer, e.cursor, pos-1, 1)
		e.executeCommand(cmd)
	}
}

func (e *Editor) Delete() {
	if e.cursor.HasSelection() {
		e.deleteSelection()
		return
	}

	pos := e.cursor.GetPosition()
	if pos < e.buffer.Length() {
		cmd := NewDeleteCommand(e.buffer, e.cursor, pos, 1)
		e.executeCommand(cmd)
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
	if cmd := e.popStack(&e.undoStack); cmd != nil {
		cmd.Undo()
		e.redoStack = append(e.redoStack, cmd)
	}
}

func (e *Editor) Redo() {
	if cmd := e.popStack(&e.redoStack); cmd != nil {
		cmd.Execute()
		e.undoStack = append(e.undoStack, cmd)
	}
}

func (e *Editor) executeCommand(cmd Command) {
	cmd.Execute()
	e.undoStack = append(e.undoStack, cmd)
	e.redoStack = make([]Command, 0)
	e.fileManager.MarkDirty()
}

func (e *Editor) deleteSelection() {
	start, end := e.cursor.GetSelection()
	length := end - start
	cmd := NewDeleteCommand(e.buffer, e.cursor, start, length)
	e.executeCommand(cmd)
	e.cursor.ClearSelection()
}

func (e *Editor) popStack(stack *[]Command) Command {
	if len(*stack) == 0 {
		return nil
	}
	lastIndex := len(*stack) - 1
	cmd := (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]
	return cmd
}
