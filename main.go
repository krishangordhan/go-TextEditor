package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	var editor *Editor
	var err error

	if len(os.Args) > 1 {
		filePath := os.Args[1]
		editor, err = NewEditorFromFile(filePath)
		if err != nil {
			log.Fatalf("Failed to open file %s: %v", filePath, err)
		}
	} else {
		editor = NewEditor("Hello World!\nThis is a simple text editor.\nTry editing this text!")
	}

	display := NewDisplay(editor)

	err = display.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer display.Close()

	inputMode := false
	inputPrompt := ""
	inputBuffer := ""
	confirmQuit := false

	display.Render()

	for {
		ev := display.GetScreen().PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if confirmQuit {
				if ev.Rune() == 'y' || ev.Rune() == 'Y' {
					return
				} else if ev.Rune() == 'n' || ev.Rune() == 'N' || ev.Key() == tcell.KeyEscape {
					confirmQuit = false
					display.Render()
				}
				continue
			}

			if inputMode {
				if ev.Key() == tcell.KeyEscape {
					inputMode = false
					inputBuffer = ""
					display.Render()
				} else if ev.Key() == tcell.KeyEnter {
					if inputBuffer != "" {
						err := editor.SaveAs(inputBuffer)
						if err != nil {
							// TODO: Handle save errors.
						}
					}
					inputMode = false
					inputBuffer = ""
					display.Render()
				} else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
					if len(inputBuffer) > 0 {
						runes := []rune(inputBuffer)
						inputBuffer = string(runes[:len(runes)-1])
					}
					display.RenderWithPrompt(inputPrompt, inputBuffer)
				} else if ev.Rune() != 0 {
					inputBuffer += string(ev.Rune())
					display.RenderWithPrompt(inputPrompt, inputBuffer)
				}
				continue
			}

			if ev.Key() == tcell.KeyCtrlQ || (ev.Modifiers()&tcell.ModCtrl != 0 && ev.Rune() == 'q') {
				if editor.GetFileManager().IsDirty() {
					confirmQuit = true
					display.RenderWithPrompt("Unsaved changes! Are you sure you want to quit? (y/n): ", "")
					continue
				}
				return
			}

			if ev.Key() == tcell.KeyCtrlS || (ev.Modifiers()&tcell.ModCtrl != 0 && ev.Rune() == 's') {
				err := editor.Save()
				if err != nil {
					// TODO: Handle save errors.
				}
				display.Render()
				continue
			}

			if ev.Key() == tcell.KeyCtrlW || (ev.Modifiers()&tcell.ModCtrl != 0 && ev.Rune() == 'w') {
				inputMode = true
				inputPrompt = "Save as: "
				inputBuffer = ""
				display.RenderWithPrompt(inputPrompt, inputBuffer)
				continue
			}

			if ev.Key() == tcell.KeyCtrlZ || (ev.Modifiers()&tcell.ModCtrl != 0 && ev.Rune() == 'z') {
				editor.Undo()
				display.Render()
				continue
			}

			if ev.Key() == tcell.KeyCtrlY || (ev.Modifiers()&tcell.ModCtrl != 0 && ev.Rune() == 'y') {
				editor.Redo()
				display.Render()
				continue
			}

			if ev.Key() == tcell.KeyCtrlC || (ev.Modifiers()&tcell.ModCtrl != 0 && ev.Rune() == 'c') {
				err := editor.Copy()
				if err != nil {
					// TODO: Handle copy errors.
				}
				display.Render()
				continue
			}

			if ev.Key() == tcell.KeyCtrlV || (ev.Modifiers()&tcell.ModCtrl != 0 && ev.Rune() == 'v') {
				err := editor.Paste()
				if err != nil {
					// TODO: Handle paste errors.
				}
				display.Render()
				continue
			}

			switch ev.Key() {
			case tcell.KeyLeft:
				if ev.Modifiers()&tcell.ModShift != 0 {
					editor.MoveCursorLeftWithSelection()
				} else {
					editor.ClearSelection()
					editor.MoveCursorLeft()
				}
			case tcell.KeyRight:
				if ev.Modifiers()&tcell.ModShift != 0 {
					editor.MoveCursorRightWithSelection()
				} else {
					editor.ClearSelection()
					editor.MoveCursorRight()
				}
			case tcell.KeyUp:
				if ev.Modifiers()&tcell.ModShift != 0 {
					editor.MoveCursorUpWithSelection()
				} else {
					editor.ClearSelection()
					editor.MoveCursorUp()
				}
			case tcell.KeyDown:
				if ev.Modifiers()&tcell.ModShift != 0 {
					editor.MoveCursorDownWithSelection()
				} else {
					editor.ClearSelection()
					editor.MoveCursorDown()
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				editor.Backspace()
			case tcell.KeyDelete:
				editor.Delete()
			case tcell.KeyEnter:
				editor.InsertNewlineWithIndent()
			default:
				if ev.Rune() == ' ' {
					editor.InsertAtCursor(" ")
				} else if ev.Rune() != 0 {
					editor.InsertAtCursor(string(ev.Rune()))
				}
			}

			display.Render()
		case *tcell.EventResize:
			display.Render()
		}
	}
}
