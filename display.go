package main

import (
	"fmt"
	"path/filepath"

	"github.com/nsf/termbox-go"
)

type Display struct {
	editor  *Editor
	scrollX int
	scrollY int
}

func NewDisplay(editor *Editor) *Display {
	return &Display{
		editor:  editor,
		scrollX: 0,
		scrollY: 0,
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
	width, height := termbox.Size()
	statusY := height - 1

	fm := d.editor.GetFileManager()
	filename := "[No Name]"
	if fm.HasFile() {
		filename = filepath.Base(fm.GetFilePath())
	}

	modifiedIndicator := ""
	if fm.IsDirty() {
		modifiedIndicator = " [+]"
	}

	cursorPos := d.editor.GetCursorPosition()
	line, col := d.editor.GetBuffer().GetLineColumn(cursorPos)

	leftStatus := fmt.Sprintf(" %s%s | Ln %d, Col %d", filename, modifiedIndicator, line+1, col)

	rightStatus := "Ctrl+S: Save | Ctrl+W: Save As | Ctrl+Q: Quit "

	for i := 0; i < width; i++ {
		termbox.SetCell(i, statusY, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}

	x := 0
	for _, r := range leftStatus {
		if x >= width {
			break
		}
		termbox.SetCell(x, statusY, r, termbox.ColorBlack, termbox.ColorWhite)
		x++
	}

	rightX := width - len(rightStatus)
	if rightX < x {
		rightX = x
	}
	for i, r := range rightStatus {
		if rightX+i >= width {
			break
		}
		termbox.SetCell(rightX+i, statusY, r, termbox.ColorBlack, termbox.ColorWhite)
	}
}

// TODO: Find a better way to do cursor line/col tracking, a lot of duplication.
func (d *Display) getCursorLineCol() (int, int) {
	cursorPos := d.editor.GetCursorPosition()
	return d.editor.GetBuffer().GetLineColumn(cursorPos)
}

func (d *Display) adjustScrollForCursor() {
	_, height := termbox.Size()
	visibleLines := height - 1 // Hard code 1 line for status bar. Yes its a magic number. Fuck off, i'll fix it later.

	cursorLine, _ := d.getCursorLineCol()

	margin := 3

	if cursorLine >= d.scrollY+visibleLines-margin {
		d.scrollY = cursorLine - visibleLines + margin + 1
	}

	if cursorLine < d.scrollY+margin {
		d.scrollY = max(cursorLine-margin, 0)
	}
}
