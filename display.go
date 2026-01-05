package main

import (
	"github.com/nsf/termbox-go"
)

type Display struct {
	editor *Editor
}

func NewDisplay(editor *Editor) *Display {
	return &Display{
		editor: editor,
	}
}

func (d *Display) Init() error {
	return termbox.Init()
}

func (d *Display) Close() {
	termbox.Close()
}

func (d *Display) Render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	text := d.editor.GetText()
	cursorPos := d.editor.GetCursorPosition()

	x, y := 0, 0
	for i, r := range []rune(text) {
		fg := termbox.ColorDefault
		bg := termbox.ColorDefault

		if i == cursorPos {
			bg = termbox.ColorWhite
			fg = termbox.ColorBlack
		}

		if r == '\n' {
			y++
			x = 0
			continue
		}

		termbox.SetCell(x, y, r, fg, bg)
		x++
	}

	if cursorPos == len([]rune(text)) {
		termbox.SetCell(x, y, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}

	d.renderStatusBar()

	termbox.Flush()
}

func (d *Display) renderStatusBar() {
	_, height := termbox.Size()
	statusY := height - 1

	status := " Ctrl+Q: Quit | Arrows: Move | Type to insert | Backspace: Delete"
	x := 0
	for _, r := range status {
		termbox.SetCell(x, statusY, r, termbox.ColorBlack, termbox.ColorWhite)
		x++
	}
}
