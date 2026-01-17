package main

type Cursor struct {
	position        int
	selectionAnchor int
}

func NewCursor() *Cursor {
	return &Cursor{
		position:        0,
		selectionAnchor: -1,
	}
}

func (c *Cursor) GetPosition() int {
	return c.position
}

func (c *Cursor) SetPosition(pos int) {
	if pos < 0 {
		pos = 0
	}
	c.position = pos
}

func (c *Cursor) MoveLeft() {
	if c.position > 0 {
		c.position--
	}
}

func (c *Cursor) MoveRight(maxLength int) {
	if c.position < maxLength {
		c.position++
	}
}

func (c *Cursor) StartSelection() {
	c.selectionAnchor = c.position
}

func (c *Cursor) ClearSelection() {
	c.selectionAnchor = -1
}

func (c *Cursor) HasSelection() bool {
	return c.selectionAnchor >= 0 && c.selectionAnchor != c.position
}

func (c *Cursor) GetSelection() (int, int) {
	if !c.HasSelection() {
		return c.position, c.position
	}

	if c.selectionAnchor < c.position {
		return c.selectionAnchor, c.position
	}
	return c.position, c.selectionAnchor
}

func (c *Cursor) GetSelectionAnchor() int {
	return c.selectionAnchor
}
