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

func TestEditor_DeleteAtCursor_DeleteAt0_ShouldNotChangeText(t *testing.T) {
	editor := NewEditor("Hello")
	editor.DeleteAtCursor(0)

	expected := "Hello"
	if editor.GetText() != expected {
		t.Errorf("Expected %q, got %q", expected, editor.GetText())
	}

	if editor.GetCursorPosition() != 0 {
		t.Errorf("Expected cursor at 0, got %d", editor.GetCursorPosition())
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

func TestEditor_MoveCursorUp_ShouldUpdateCursorPosition(t *testing.T) {
	editor := NewEditor("Line 1\nLine 2\nLine 3")
	editor.SetCursorPosition(10)

	editor.MoveCursorUp()

	expected := 3
	if editor.GetCursorPosition() != expected {
		t.Errorf("Expected cursor at %d, got %d", expected, editor.GetCursorPosition())
	}
}

func TestEditor_MoveCursorDown_ShouldUpdateCursorPosition(t *testing.T) {
	editor := NewEditor("Line 1\nLine 2\nLine 3")
	editor.SetCursorPosition(3)

	editor.MoveCursorDown()

	expected := 10
	if editor.GetCursorPosition() != expected {
		t.Errorf("Expected cursor at %d, got %d", expected, editor.GetCursorPosition())
	}
}

func TestEditor_MoveCursorUpAtStart_ShouldNotChangeCursorPosition(t *testing.T) {
	editor := NewEditor("Line 1\nLine 2")
	editor.MoveCursorUp()

	if editor.GetCursorPosition() != 0 {
		t.Errorf("Expected cursor at 0, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_MoveCursorDownAtEnd_ShouldNotChangeCursorPosition(t *testing.T) {
	editor := NewEditor("Line 1\nLine 2")
	editor.SetCursorPosition(13)

	editor.MoveCursorDown()

	if editor.GetCursorPosition() != 13 {
		t.Errorf("Expected cursor at 13, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_MoveCursorUpShortLine_ShouldUpdateCursorPosition(t *testing.T) {
	editor := NewEditor("Short\nMuch longer line")
	editor.SetCursorPosition(15)

	editor.MoveCursorUp()

	expected := 5
	if editor.GetCursorPosition() != expected {
		t.Errorf("Expected cursor at %d, got %d", expected, editor.GetCursorPosition())
	}
}

func TestEditor_MoveCursorDown_ShouldRememberCursorLocation(t *testing.T) {
	editor := NewEditor("Long line here\nX\nAnother long line")
	editor.SetCursorPosition(5)

	editor.MoveCursorDown()
	editor.MoveCursorDown()

	expected := 22
	if editor.GetCursorPosition() != expected {
		t.Errorf("Expected cursor at %d, got %d", expected, editor.GetCursorPosition())
	}
}

func TestEditor_NewEditorFromFile_LoadsFileContent(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/test.txt"
	content := "Line 1\nLine 2\nLine 3"

	err := writeTestFile(tmpFile, content)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	editor, err := NewEditorFromFile(tmpFile)
	if err != nil {
		t.Fatalf("NewEditorFromFile failed: %v", err)
	}

	if editor.GetText() != content {
		t.Errorf("Expected text %q, got %q", content, editor.GetText())
	}

	if editor.GetFileManager().GetFilePath() != tmpFile {
		t.Errorf("Expected file path %q, got %q", tmpFile, editor.GetFileManager().GetFilePath())
	}

	if editor.GetFileManager().IsDirty() {
		t.Error("Expected file to not be dirty after loading")
	}
}

func TestEditor_NewEditorFromFile_NonExistentFile(t *testing.T) {
	_, err := NewEditorFromFile("/nonexistent/file.txt")
	if err == nil {
		t.Error("Expected error when loading non-existent file")
	}
}

func TestEditor_Save_WritesContentToFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/test.txt"
	initialContent := "Initial content"

	err := writeTestFile(tmpFile, initialContent)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	editor, err := NewEditorFromFile(tmpFile)
	if err != nil {
		t.Fatalf("NewEditorFromFile failed: %v", err)
	}

	editor.SetCursorPosition(editor.buffer.Length())
	editor.InsertAtCursor("\nNew line")

	if !editor.GetFileManager().IsDirty() {
		t.Error("Expected file to be dirty after modification")
	}

	err = editor.Save()
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	if editor.GetFileManager().IsDirty() {
		t.Error("Expected file to not be dirty after save")
	}

	savedContent := readTestFile(t, tmpFile)
	expectedContent := initialContent + "\nNew line"
	if savedContent != expectedContent {
		t.Errorf("Expected saved content %q, got %q", expectedContent, savedContent)
	}
}

func TestEditor_SaveAs_WritesToNewFile(t *testing.T) {
	editor := NewEditor("Original content")
	tmpDir := t.TempDir()
	newFile := tmpDir + "/new_file.txt"

	err := editor.SaveAs(newFile)
	if err != nil {
		t.Fatalf("SaveAs failed: %v", err)
	}

	if editor.GetFileManager().GetFilePath() != newFile {
		t.Errorf("Expected file path %q, got %q", newFile, editor.GetFileManager().GetFilePath())
	}

	if editor.GetFileManager().IsDirty() {
		t.Error("Expected file to not be dirty after SaveAs")
	}

	savedContent := readTestFile(t, newFile)
	if savedContent != "Original content" {
		t.Errorf("Expected saved content %q, got %q", "Original content", savedContent)
	}
}

func TestEditor_InsertMarksFileDirty(t *testing.T) {
	editor := NewEditor("Test")
	editor.InsertAtCursor("x")

	if !editor.GetFileManager().IsDirty() {
		t.Error("Expected file to be dirty after insert")
	}
}

func TestEditor_BackspaceMarksFileDirty(t *testing.T) {
	editor := NewEditor("Test")
	editor.SetCursorPosition(1)
	editor.Backspace()

	if !editor.GetFileManager().IsDirty() {
		t.Error("Expected file to be dirty after backspace")
	}
}

func TestEditor_DeleteMarksFileDirty(t *testing.T) {
	editor := NewEditor("Test")
	editor.Delete()

	if !editor.GetFileManager().IsDirty() {
		t.Error("Expected file to be dirty after delete")
	}
}

func writeTestFile(path, content string) error {
	return writeFile(path, content)
}

func readTestFile(t *testing.T, path string) string {
	content, err := readFile(path)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}
	return content
}

func writeFile(path, content string) error {
	fm := NewFileManagerWithPath(path)
	return fm.WriteFile(content)
}

func readFile(path string) (string, error) {
	fm := NewFileManagerWithPath(path)
	return fm.ReadFile()
}

func TestEditor_InsertAtCursor_PushesToUndoStack(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	if len(editor.undoStack) != 0 {
		t.Errorf("Expected empty undo stack, got length %d", len(editor.undoStack))
	}

	editor.InsertAtCursor(" World")

	if len(editor.undoStack) != 1 {
		t.Errorf("Expected undo stack length 1, got %d", len(editor.undoStack))
	}

	if len(editor.redoStack) != 0 {
		t.Errorf("Expected empty redo stack, got length %d", len(editor.redoStack))
	}
}

func TestEditor_MultipleInserts_BuildsUndoStack(t *testing.T) {
	editor := NewEditor("")

	editor.InsertAtCursor("H")
	editor.InsertAtCursor("i")

	if len(editor.undoStack) != 2 {
		t.Errorf("Expected undo stack length 2, got %d", len(editor.undoStack))
	}

	if editor.GetText() != "Hi" {
		t.Errorf("Expected 'Hi', got '%s'", editor.GetText())
	}
}

func TestEditor_Backspace_PushesToUndoStack(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	editor.Backspace()

	if len(editor.undoStack) != 1 {
		t.Errorf("Expected undo stack length 1, got %d", len(editor.undoStack))
	}

	if editor.GetText() != "Hell" {
		t.Errorf("Expected 'Hell', got '%s'", editor.GetText())
	}
}

func TestEditor_Delete_PushesToUndoStack(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(0)

	editor.Delete()

	if len(editor.undoStack) != 1 {
		t.Errorf("Expected undo stack length 1, got %d", len(editor.undoStack))
	}

	if editor.GetText() != "ello" {
		t.Errorf("Expected 'ello', got '%s'", editor.GetText())
	}
}

func TestEditor_NewOperation_ClearsRedoStack(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	// Simulate redo stack having content
	cmd := NewInsertCommand(editor.buffer, editor.cursor, "test", 0)
	editor.redoStack = append(editor.redoStack, cmd)

	if len(editor.redoStack) != 1 {
		t.Errorf("Setup: Expected redo stack length 1, got %d", len(editor.redoStack))
	}

	// New operation should clear redo stack
	editor.InsertAtCursor("!")

	if len(editor.redoStack) != 0 {
		t.Errorf("Expected redo stack cleared, got length %d", len(editor.redoStack))
	}
}

func TestEditor_DeleteAtCursor_PushesToUndoStack(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(6)

	editor.DeleteAtCursor(5)

	if len(editor.undoStack) != 1 {
		t.Errorf("Expected undo stack length 1, got %d", len(editor.undoStack))
	}

	if editor.GetText() != "Hello " {
		t.Errorf("Expected 'Hello ', got '%s'", editor.GetText())
	}
}
