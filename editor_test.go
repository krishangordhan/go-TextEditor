package main

import (
	"testing"
)

func TestEditor_NewEditor_CreatesObjectWithInitialText(t *testing.T) {
	editor := NewEditor("Hello")

	if editor.GetText() != "Hello" {
		t.Errorf("Expected text %q, got %q", "Hello", editor.GetText())
	}

	if editor.GetCursorPosition() != 0 {
		t.Errorf("Expected cursor at 0, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_SetCursorPosition_SetsPositionCorrectly(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(3)

	if editor.GetCursorPosition() != 3 {
		t.Errorf("Expected cursor at 3, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_SetCursorPosition_SetNegativeShouldSet0(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(-5)

	if editor.GetCursorPosition() != 0 {
		t.Errorf("Expected cursor clamped to 0, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_SetCursorPosition_SetBeyondEndShouldSetMax(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(100)

	if editor.GetCursorPosition() != 5 {
		t.Errorf("Expected cursor clamped to 5, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_MoveCursorLeft_ShouldMoveLeft(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(3)
	editor.MoveCursorLeft()

	if editor.GetCursorPosition() != 2 {
		t.Errorf("Expected cursor at 2, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_MoveCursorRight_ShouldMoveRight(t *testing.T) {
	editor := NewEditor("Hello")
	editor.MoveCursorRight()

	if editor.GetCursorPosition() != 1 {
		t.Errorf("Expected cursor at 1, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_InsertAtCursor_ShouldInsertText(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(6)
	editor.InsertAtCursor("Beautiful ")

	expected := "Hello Beautiful World"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 16 {
		t.Errorf("Expected cursor at 16, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_InsertAtCursorAtStart_ShouldInsertTextAtStart(t *testing.T) {
	editor := NewEditor("World")
	editor.InsertAtCursor("Hello ")

	expected := "Hello World"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 6 {
		t.Errorf("Expected cursor at 6, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_InsertAtCursorAtEnd_ShouldInsertTextAtEnd(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)
	editor.InsertAtCursor(" World")

	expected := "Hello World"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 11 {
		t.Errorf("Expected cursor at 11, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_DeleteAtCursor_ShouldDeleteText(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(6)
	editor.DeleteAtCursor(6)

	expected := "Hello "
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 6 {
		t.Errorf("Expected cursor at 6, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Backspace_ShouldDeleteCharacterBeforeCursor(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)
	editor.Backspace()

	expected := "Hell"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 4 {
		t.Errorf("Expected cursor at 4, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Backspace_AtStart_ShouldNotChangeTextOrCursor(t *testing.T) {
	editor := NewEditor("Hello")
	editor.Backspace()

	expected := "Hello"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 0 {
		t.Errorf("Expected cursor at 0, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Delete_ShouldDeleteCharacterAtCursor(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(4)
	editor.Delete()

	expected := "Hell"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 4 {
		t.Errorf("Expected cursor at 4, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Delete_AtEnd_ShouldNotChangeTextOrCursor(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)
	editor.Delete()

	expected := "Hello"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 5 {
		t.Errorf("Expected cursor at 5, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_ComplexEditing(t *testing.T) {
	editor := NewEditor("Hello World")

	editor.SetCursorPosition(11)
	editor.InsertAtCursor("!")

	editor.SetCursorPosition(5)
	editor.InsertAtCursor(",")

	editor.Backspace()

	expected := "Hello World!"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}
}

func TestEditor_TypeAndBackspace_ShouldTypeAndDeleteCharacters(t *testing.T) {
	editor := NewEditor("")

	editor.InsertAtCursor("H")
	editor.InsertAtCursor("e")
	editor.InsertAtCursor("l")
	editor.InsertAtCursor("l")
	editor.InsertAtCursor("o")

	if editor.GetText() != "Hello" {
		t.Errorf("Expected %q, got %q", "Hello", editor.GetText())
	}

	editor.Backspace()
	editor.Backspace()

	expected := "Hel"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 3 {
		t.Errorf("Expected cursor at 3, got %d", editor.GetCursorPosition())
	}
}
