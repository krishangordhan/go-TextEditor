package main

import (
	"log"
	"os"

	"github.com/nsf/termbox-go"
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

	display.Render()

	for {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyCtrlQ {
				break
			}

			if ev.Key == termbox.KeyCtrlS {
				err := editor.Save()
				if err != nil {
					// TODO: Handle save errors.
				}
			}

			switch ev.Key {
			case termbox.KeyArrowLeft:
				editor.MoveCursorLeft()
			case termbox.KeyArrowRight:
				editor.MoveCursorRight()
			case termbox.KeyArrowUp:
				editor.MoveCursorUp()
			case termbox.KeyArrowDown:
				editor.MoveCursorDown()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				editor.Backspace()
			case termbox.KeyDelete:
				editor.Delete()
			case termbox.KeyEnter:
				editor.InsertAtCursor("\n")
			case termbox.KeySpace:
				editor.InsertAtCursor(" ")
			default:
				if ev.Ch != 0 {
					editor.InsertAtCursor(string(ev.Ch))
				}
			}

			display.Render()
		}
	}
}
