package main

import (
	"fmt"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
)

const TabWidth = 4

type Display struct {
	editor  *Editor
	scrollX int
	scrollY int
	screen  tcell.Screen
}

func NewDisplay(editor *Editor) *Display {
	return &Display{
		editor:  editor,
		scrollX: 0,
		scrollY: 0,
		screen:  nil,
	}
}

func (d *Display) GetScreen() tcell.Screen {
	return d.screen
}

func (d *Display) Init() error {
	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := screen.Init(); err != nil {
		return err
	}
	d.screen = screen
	return nil
}

func (d *Display) Close() {
	if d.screen != nil {
		d.screen.Fini()
	}
}

func (d *Display) Render() {
	d.renderEditor()
	d.renderStatusBar()
	d.screen.Show()
}

func (d *Display) RenderWithPrompt(prompt, input string) {
	d.renderEditor()
	d.renderPrompt(prompt, input)
	d.screen.Show()
}

func (d *Display) getLineNumberWidth() int {
	lineCount := d.editor.GetBuffer().GetLineCount()
	width := 1
	for n := lineCount; n >= 10; n /= 10 {
		width++
	}
	return width + 1
}

func (d *Display) renderEditor() {
	d.adjustScrollForCursor()
	d.screen.Clear()

	width, height := d.screen.Size()
	visibleLines := height - 1 // Hard code 1 line for status bar. Yes its a magic number.
	lineNumWidth := d.getLineNumberWidth()
	visibleCols := width - lineNumWidth

	defStyle := tcell.StyleDefault
	lineNumStyle := defStyle.Foreground(tcell.ColorYellow)

	for i := range visibleLines {
		lineNum := d.scrollY + i + 1
		lineText := fmt.Sprintf("%*d ", lineNumWidth-1, lineNum)
		for j, r := range lineText {
			d.screen.SetContent(j, i, r, nil, lineNumStyle)
		}
	}

	text := d.editor.GetText()
	cursorPos := d.editor.GetCursorPosition()
	hasSelection := d.editor.HasSelection()
	selStart, selEnd := d.editor.GetSelection()

	y := 0
	lineNum := 0
	colNum := 0

	for i, r := range []rune(text) {
		if lineNum < d.scrollY {
			switch r {
			case '\n':
				lineNum++
				colNum = 0
			case '\t':
				colNum += TabWidth - (colNum % TabWidth)
			default:
				colNum++
			}
			continue
		}

		if y >= visibleLines {
			break
		}

		style := defStyle

		if hasSelection && i >= selStart && i < selEnd {
			style = style.Background(tcell.ColorTeal).Foreground(tcell.ColorBlack)
		}

		if i == cursorPos {
			style = style.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
		}

		if r == '\n' {
			if i == cursorPos && colNum >= d.scrollX && colNum < d.scrollX+visibleCols {
				d.screen.SetContent(lineNumWidth+colNum-d.scrollX, y, ' ', nil, style)
			}
			y++
			lineNum++
			colNum = 0
			continue
		}

		if r == '\t' {
			tabStop := TabWidth - (colNum % TabWidth)
			for range tabStop {
				if colNum >= d.scrollX && colNum < d.scrollX+visibleCols {
					d.screen.SetContent(lineNumWidth+colNum-d.scrollX, y, ' ', nil, style)
				}
				colNum++
			}
			continue
		}

		if colNum >= d.scrollX && colNum < d.scrollX+visibleCols {
			d.screen.SetContent(lineNumWidth+colNum-d.scrollX, y, r, nil, style)
		}

		colNum++
	}

	if cursorPos == len([]rune(text)) {
		if y < visibleLines && colNum >= d.scrollX && colNum < d.scrollX+visibleCols {
			cursorStyle := defStyle.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
			d.screen.SetContent(lineNumWidth+colNum-d.scrollX, y, ' ', nil, cursorStyle)
		}
	}
}

func (d *Display) renderPrompt(prompt, input string) {
	_, height := d.screen.Size()
	promptY := height - 1

	promptStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
	cursorStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

	fullPrompt := prompt + input
	x := 0
	for _, r := range fullPrompt {
		d.screen.SetContent(x, promptY, r, nil, promptStyle)
		x++
	}
	d.screen.SetContent(x, promptY, ' ', nil, cursorStyle)
}

func (d *Display) renderStatusBar() {
	width, height := d.screen.Size()
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

	rightStatus := "Ctrl+C: Copy | Ctrl+V: Paste | Ctrl+Z: Undo | Ctrl+Y: Redo | Ctrl+S: Save | Ctrl+Q: Quit "

	statusStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

	for i := 0; i < width; i++ {
		d.screen.SetContent(i, statusY, ' ', nil, statusStyle)
	}

	x := 0
	for _, r := range leftStatus {
		if x >= width {
			break
		}
		d.screen.SetContent(x, statusY, r, nil, statusStyle)
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
		d.screen.SetContent(rightX+i, statusY, r, nil, statusStyle)
	}
}

// TODO: Find a better way to do cursor line/col tracking, a lot of duplication.
func (d *Display) getCursorLineCol() (int, int) {
	cursorPos := d.editor.GetCursorPosition()
	return d.editor.GetBuffer().GetLineColumn(cursorPos)
}

func (d *Display) adjustScrollForCursor() {
	width, height := d.screen.Size()
	visibleLines := height - 1 // Hard code 1 line for status bar. Yes its a magic number. FUck off.
	lineNumWidth := d.getLineNumberWidth()
	visibleCols := width - lineNumWidth

	cursorLine, cursorCol := d.getCursorLineCol()

	margin := 3

	if cursorLine >= d.scrollY+visibleLines-margin {
		d.scrollY = cursorLine - visibleLines + margin + 1
	}

	if cursorLine < d.scrollY+margin {
		d.scrollY = max(cursorLine-margin, 0)
	}

	if cursorCol >= d.scrollX+visibleCols-margin {
		d.scrollX = cursorCol - visibleCols + margin + 1
	}

	if cursorCol < d.scrollX+margin {
		d.scrollX = max(cursorCol-margin, 0)
	}
}
