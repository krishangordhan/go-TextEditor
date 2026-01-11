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
			if i == cursorPos {
				termbox.SetCell(x, y, ' ', fg, bg)
			}
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

func (d *Display) RenderWithPrompt(prompt, input string) {
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
			if i == cursorPos {
				termbox.SetCell(x, y, ' ', fg, bg)
			}
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

	d.renderPrompt(prompt, input)

	termbox.Flush()
}

func (d *Display) renderPrompt(prompt, input string) {
	_, height := termbox.Size()
	promptY := height - 1

	fullPrompt := prompt + input
	x := 0
	for _, r := range fullPrompt {
		termbox.SetCell(x, promptY, r, termbox.ColorWhite, termbox.ColorBlue)
		x++
	}
	termbox.SetCell(x, promptY, ' ', termbox.ColorBlack, termbox.ColorWhite)
}

func (d *Display) renderStatusBar() {
	_, height := termbox.Size()
	statusY := height - 1

	status := " Ctrl+S: Save | Ctrl+W: Save As | Ctrl+Q: Quit | Arrows: Move"
	x := 0
	for _, r := range status {
		termbox.SetCell(x, statusY, r, termbox.ColorBlack, termbox.ColorWhite)
		x++
	}
}
