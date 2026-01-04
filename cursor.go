package main

type Cursor struct {
	position int
}

func NewCursor() *Cursor {
	return &Cursor{position: 0}
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
