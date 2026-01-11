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

	inputMode := false
	inputPrompt := ""
	inputBuffer := ""

	display.Render()

	for {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			if inputMode {
				if ev.Key == termbox.KeyEsc {
					inputMode = false
					inputBuffer = ""
					display.Render()
				} else if ev.Key == termbox.KeyEnter {
					if inputBuffer != "" {
						err := editor.SaveAs(inputBuffer)
						if err != nil {
							// TODO: Handle save errors.
						}
					}
					inputMode = false
					inputBuffer = ""
					display.Render()
				} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
					if len(inputBuffer) > 0 {
						runes := []rune(inputBuffer)
						inputBuffer = string(runes[:len(runes)-1])
					}
					display.RenderWithSaveAsPrompt(inputPrompt, inputBuffer)
				} else if ev.Ch != 0 {
					inputBuffer += string(ev.Ch)
					display.RenderWithSaveAsPrompt(inputPrompt, inputBuffer)
				}
				continue
			}

			if ev.Key == termbox.KeyCtrlQ {
				break
			}

			if ev.Key == termbox.KeyCtrlS {
				err := editor.Save()
				if err != nil {
					// TODO: Handle save errors.
				}
			}

			if ev.Key == termbox.KeyCtrlW {
				inputMode = true
				inputPrompt = "Save as: "
				inputBuffer = ""
				display.RenderWithSaveAsPrompt(inputPrompt, inputBuffer)
				continue
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
