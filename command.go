package main

type Command interface {
	Execute()
	Undo()
}

type InsertCommand struct {
	buffer       *PieceTable
	cursor       *Cursor
	text         string
	position     int
	cursorBefore int
	cursorAfter  int
}

func NewInsertCommand(buffer *PieceTable, cursor *Cursor, text string, position int) *InsertCommand {
	return &InsertCommand{
		buffer:       buffer,
		cursor:       cursor,
		text:         text,
		position:     position,
		cursorBefore: cursor.GetPosition(),
	}
}

func (c *InsertCommand) Execute() {
	c.buffer.Insert(c.position, c.text)
	c.cursor.SetPosition(c.position + len(c.text))
	c.cursorAfter = c.cursor.GetPosition()
}

func (c *InsertCommand) Undo() {
	c.buffer.Delete(c.position, len(c.text))
	c.cursor.SetPosition(c.cursorBefore)
}

type DeleteCommand struct {
	buffer       *PieceTable
	cursor       *Cursor
	position     int
	length       int
	deletedText  string
	cursorBefore int
	cursorAfter  int
}

func NewDeleteCommand(buffer *PieceTable, cursor *Cursor, position int, length int) *DeleteCommand {
	return &DeleteCommand{
		buffer:       buffer,
		cursor:       cursor,
		position:     position,
		length:       length,
		cursorBefore: cursor.GetPosition(),
	}
}

func (c *DeleteCommand) Execute() {
	c.deletedText = c.buffer.Substring(c.position, c.position+c.length)
	c.buffer.Delete(c.position, c.length)
	c.cursor.SetPosition(c.position)
	c.cursorAfter = c.cursor.GetPosition()
}

func (c *DeleteCommand) Undo() {
	c.buffer.Insert(c.position, c.deletedText)
	c.cursor.SetPosition(c.cursorBefore)
}
