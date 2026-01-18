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

func TestEditor_Undo_EmptyStack(t *testing.T) {
	editor := NewEditor("Hello")

	editor.Undo()

	if editor.GetText() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", editor.GetText())
	}

	if len(editor.undoStack) != 0 {
		t.Errorf("Expected empty undo stack, got length %d", len(editor.undoStack))
	}
}

func TestEditor_Undo_SingleInsert(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)
	editor.InsertAtCursor(" World")

	if editor.GetText() != "Hello World" {
		t.Errorf("Before undo: Expected 'Hello World', got '%s'", editor.GetText())
	}

	editor.Undo()

	if editor.GetText() != "Hello" {
		t.Errorf("After undo: Expected 'Hello', got '%s'", editor.GetText())
	}

	if editor.GetCursorPosition() != 5 {
		t.Errorf("After undo: Expected cursor at 5, got %d", editor.GetCursorPosition())
	}

	if len(editor.undoStack) != 0 {
		t.Errorf("Expected empty undo stack, got length %d", len(editor.undoStack))
	}

	if len(editor.redoStack) != 1 {
		t.Errorf("Expected redo stack length 1, got %d", len(editor.redoStack))
	}
}

func TestEditor_Undo_SingleDelete(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)
	editor.Delete()

	if editor.GetText() != "HelloWorld" {
		t.Errorf("Before undo: Expected 'HelloWorld', got '%s'", editor.GetText())
	}

	editor.Undo()

	if editor.GetText() != "Hello World" {
		t.Errorf("After undo: Expected 'Hello World', got '%s'", editor.GetText())
	}

	if editor.GetCursorPosition() != 5 {
		t.Errorf("After undo: Expected cursor at 5, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Undo_SingleBackspace(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)
	editor.Backspace()

	if editor.GetText() != "Hell" {
		t.Errorf("Before undo: Expected 'Hell', got '%s'", editor.GetText())
	}

	editor.Undo()

	if editor.GetText() != "Hello" {
		t.Errorf("After undo: Expected 'Hello', got '%s'", editor.GetText())
	}

	if editor.GetCursorPosition() != 5 {
		t.Errorf("After undo: Expected cursor at 5, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Undo_MultipleInserts(t *testing.T) {
	editor := NewEditor("")

	editor.InsertAtCursor("H")
	editor.InsertAtCursor("e")
	editor.InsertAtCursor("l")
	editor.InsertAtCursor("l")
	editor.InsertAtCursor("o")

	if editor.GetText() != "Hello" {
		t.Errorf("Before undo: Expected 'Hello', got '%s'", editor.GetText())
	}

	editor.Undo()
	if editor.GetText() != "Hell" {
		t.Errorf("After 1st undo: Expected 'Hell', got '%s'", editor.GetText())
	}

	editor.Undo()
	if editor.GetText() != "Hel" {
		t.Errorf("After 2nd undo: Expected 'Hel', got '%s'", editor.GetText())
	}

	if len(editor.undoStack) != 3 {
		t.Errorf("Expected undo stack length 3, got %d", len(editor.undoStack))
	}

	if len(editor.redoStack) != 2 {
		t.Errorf("Expected redo stack length 2, got %d", len(editor.redoStack))
	}
}

func TestEditor_Undo_ComplexSequence(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	editor.InsertAtCursor(" World")
	if editor.GetText() != "Hello World" {
		t.Errorf("After insert: Expected 'Hello World', got '%s'", editor.GetText())
	}

	editor.SetCursorPosition(6)
	editor.DeleteAtCursor(5)
	if editor.GetText() != "Hello " {
		t.Errorf("After delete: Expected 'Hello ', got '%s'", editor.GetText())
	}

	editor.Undo()
	if editor.GetText() != "Hello World" {
		t.Errorf("After undo delete: Expected 'Hello World', got '%s'", editor.GetText())
	}

	editor.Undo()
	if editor.GetText() != "Hello" {
		t.Errorf("After undo insert: Expected 'Hello', got '%s'", editor.GetText())
	}

	if editor.GetCursorPosition() != 5 {
		t.Errorf("Final cursor: Expected 5, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Undo_AllOperations(t *testing.T) {
	editor := NewEditor("")

	editor.InsertAtCursor("a")
	editor.InsertAtCursor("b")
	editor.InsertAtCursor("c")

	editor.Undo()
	editor.Undo()
	editor.Undo()

	if editor.GetText() != "" {
		t.Errorf("Expected empty text, got '%s'", editor.GetText())
	}

	if len(editor.undoStack) != 0 {
		t.Errorf("Expected empty undo stack, got length %d", len(editor.undoStack))
	}

	if len(editor.redoStack) != 3 {
		t.Errorf("Expected redo stack length 3, got %d", len(editor.redoStack))
	}
}

func TestEditor_Undo_PreservesCursorColumn(t *testing.T) {
	editor := NewEditor("Hello\nWorld")
	editor.SetCursorPosition(11)

	editor.InsertAtCursor("!")

	if editor.GetCursorPosition() != 12 {
		t.Errorf("After insert: Expected cursor at 12, got %d", editor.GetCursorPosition())
	}

	editor.Undo()

	if editor.GetCursorPosition() != 11 {
		t.Errorf("After undo: Expected cursor at 11, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Redo_EmptyStack(t *testing.T) {
	editor := NewEditor("Hello")

	editor.Redo()

	if editor.GetText() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", editor.GetText())
	}

	if len(editor.redoStack) != 0 {
		t.Errorf("Expected empty redo stack, got length %d", len(editor.redoStack))
	}
}

func TestEditor_Redo_SingleInsert(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)
	editor.InsertAtCursor(" World")

	editor.Undo()
	if editor.GetText() != "Hello" {
		t.Errorf("After undo: Expected 'Hello', got '%s'", editor.GetText())
	}

	editor.Redo()

	if editor.GetText() != "Hello World" {
		t.Errorf("After redo: Expected 'Hello World', got '%s'", editor.GetText())
	}

	if editor.GetCursorPosition() != 11 {
		t.Errorf("After redo: Expected cursor at 11, got %d", editor.GetCursorPosition())
	}

	if len(editor.redoStack) != 0 {
		t.Errorf("Expected empty redo stack, got length %d", len(editor.redoStack))
	}

	if len(editor.undoStack) != 1 {
		t.Errorf("Expected undo stack length 1, got %d", len(editor.undoStack))
	}
}

func TestEditor_Redo_SingleDelete(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)
	editor.Delete()

	editor.Undo()
	if editor.GetText() != "Hello World" {
		t.Errorf("After undo: Expected 'Hello World', got '%s'", editor.GetText())
	}

	editor.Redo()

	if editor.GetText() != "HelloWorld" {
		t.Errorf("After redo: Expected 'HelloWorld', got '%s'", editor.GetText())
	}

	if editor.GetCursorPosition() != 5 {
		t.Errorf("After redo: Expected cursor at 5, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Redo_MultipleOperations(t *testing.T) {
	editor := NewEditor("")

	editor.InsertAtCursor("H")
	editor.InsertAtCursor("e")
	editor.InsertAtCursor("l")
	editor.InsertAtCursor("l")
	editor.InsertAtCursor("o")

	editor.Undo()
	editor.Undo()
	editor.Undo()

	if editor.GetText() != "He" {
		t.Errorf("After undos: Expected 'He', got '%s'", editor.GetText())
	}

	editor.Redo()
	if editor.GetText() != "Hel" {
		t.Errorf("After 1st redo: Expected 'Hel', got '%s'", editor.GetText())
	}

	editor.Redo()
	if editor.GetText() != "Hell" {
		t.Errorf("After 2nd redo: Expected 'Hell', got '%s'", editor.GetText())
	}

	if len(editor.redoStack) != 1 {
		t.Errorf("Expected redo stack length 1, got %d", len(editor.redoStack))
	}

	if len(editor.undoStack) != 4 {
		t.Errorf("Expected undo stack length 4, got %d", len(editor.undoStack))
	}
}

func TestEditor_Redo_ComplexSequence(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	editor.InsertAtCursor(" World")

	editor.SetCursorPosition(6)
	editor.DeleteAtCursor(5)

	if editor.GetText() != "Hello " {
		t.Errorf("After operations: Expected 'Hello ', got '%s'", editor.GetText())
	}

	editor.Undo()
	editor.Undo()

	if editor.GetText() != "Hello" {
		t.Errorf("After undos: Expected 'Hello', got '%s'", editor.GetText())
	}

	editor.Redo()
	if editor.GetText() != "Hello World" {
		t.Errorf("After 1st redo: Expected 'Hello World', got '%s'", editor.GetText())
	}

	editor.Redo()
	if editor.GetText() != "Hello " {
		t.Errorf("After 2nd redo: Expected 'Hello ', got '%s'", editor.GetText())
	}
}

func TestEditor_Redo_AllOperations(t *testing.T) {
	editor := NewEditor("")

	editor.InsertAtCursor("a")
	editor.InsertAtCursor("b")
	editor.InsertAtCursor("c")

	editor.Undo()
	editor.Undo()
	editor.Undo()

	if editor.GetText() != "" {
		t.Errorf("After undos: Expected empty, got '%s'", editor.GetText())
	}

	editor.Redo()
	editor.Redo()
	editor.Redo()

	if editor.GetText() != "abc" {
		t.Errorf("After redos: Expected 'abc', got '%s'", editor.GetText())
	}

	if len(editor.redoStack) != 0 {
		t.Errorf("Expected empty redo stack, got length %d", len(editor.redoStack))
	}

	if len(editor.undoStack) != 3 {
		t.Errorf("Expected undo stack length 3, got %d", len(editor.undoStack))
	}
}

func TestEditor_Redo_ClearedByNewOperation(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	editor.InsertAtCursor(" World")
	editor.Undo()

	if len(editor.redoStack) != 1 {
		t.Errorf("Before new op: Expected redo stack length 1, got %d", len(editor.redoStack))
	}

	editor.InsertAtCursor("!")

	if len(editor.redoStack) != 0 {
		t.Errorf("After new op: Expected empty redo stack, got length %d", len(editor.redoStack))
	}

	if editor.GetText() != "Hello!" {
		t.Errorf("Expected 'Hello!', got '%s'", editor.GetText())
	}
}

func TestEditor_UndoRedo_PreservesCursor(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	editor.InsertAtCursor(" World")

	initialCursor := editor.GetCursorPosition()
	if initialCursor != 11 {
		t.Errorf("After insert: Expected cursor at 11, got %d", initialCursor)
	}

	editor.Undo()
	if editor.GetCursorPosition() != 5 {
		t.Errorf("After undo: Expected cursor at 5, got %d", editor.GetCursorPosition())
	}

	editor.Redo()
	if editor.GetCursorPosition() != 11 {
		t.Errorf("After redo: Expected cursor at 11, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_RedoAfterPartialUndo(t *testing.T) {
	editor := NewEditor("")

	editor.InsertAtCursor("a")
	editor.InsertAtCursor("b")
	editor.InsertAtCursor("c")
	editor.InsertAtCursor("d")

	editor.Undo()
	editor.Undo()

	if editor.GetText() != "ab" {
		t.Errorf("After undos: Expected 'ab', got '%s'", editor.GetText())
	}

	editor.Redo()

	if editor.GetText() != "abc" {
		t.Errorf("After redo: Expected 'abc', got '%s'", editor.GetText())
	}

	if len(editor.undoStack) != 3 {
		t.Errorf("Expected undo stack length 3, got %d", len(editor.undoStack))
	}

	if len(editor.redoStack) != 1 {
		t.Errorf("Expected redo stack length 1, got %d", len(editor.redoStack))
	}
}

func TestEditor_MoveCursorLeftWithSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)

	editor.MoveCursorLeftWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection to be active")
	}

	start, end := editor.GetSelection()
	if start != 4 || end != 5 {
		t.Errorf("Expected selection (4, 5), got (%d, %d)", start, end)
	}
}

func TestEditor_MoveCursorRightWithSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)

	editor.MoveCursorRightWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection to be active")
	}

	start, end := editor.GetSelection()
	if start != 5 || end != 6 {
		t.Errorf("Expected selection (5, 6), got (%d, %d)", start, end)
	}
}

func TestEditor_MoveCursorUpWithSelection(t *testing.T) {
	editor := NewEditor("Hello\nWorld")
	editor.SetCursorPosition(8)

	editor.MoveCursorUpWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection to be active")
	}

	start, end := editor.GetSelection()
	// Moving up from position 8 (col 2 in line 1) goes to position 2 (col 2 in line 0)
	if start != 2 || end != 8 {
		t.Errorf("Expected selection (2, 8), got (%d, %d)", start, end)
	}
}

func TestEditor_MoveCursorDownWithSelection(t *testing.T) {
	editor := NewEditor("Hello\nWorld")
	editor.SetCursorPosition(2)

	editor.MoveCursorDownWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection to be active")
	}

	start, end := editor.GetSelection()
	if start != 2 || end != 8 {
		t.Errorf("Expected selection (2, 8), got (%d, %d)", start, end)
	}
}

func TestEditor_ExtendSelectionMultipleMoves(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)

	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()

	start, end := editor.GetSelection()
	if start != 5 || end != 8 {
		t.Errorf("Expected selection (5, 8), got (%d, %d)", start, end)
	}
}

func TestEditor_SelectionBackward(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)

	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()

	start, end := editor.GetSelection()
	if start != 3 || end != 5 {
		t.Errorf("Expected selection (3, 5), got (%d, %d)", start, end)
	}
}

func TestEditor_ClearSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)

	editor.MoveCursorRightWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection before clear")
	}

	editor.ClearSelection()

	if editor.HasSelection() {
		t.Error("Expected no selection after clear")
	}
}

func TestEditor_HasSelection_NoSelection(t *testing.T) {
	editor := NewEditor("Hello World")

	if editor.HasSelection() {
		t.Error("Expected no selection in new editor")
	}
}

func TestEditor_GetSelection_NoSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)

	start, end := editor.GetSelection()
	if start != 5 || end != 5 {
		t.Errorf("Expected (5, 5) when no selection, got (%d, %d)", start, end)
	}
}

func TestEditor_SelectionBidirectional(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(5)

	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()

	start, end := editor.GetSelection()
	if start != 5 || end != 7 {
		t.Errorf("After forward: Expected (5, 7), got (%d, %d)", start, end)
	}

	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()

	start, end = editor.GetSelection()
	if start != 3 || end != 5 {
		t.Errorf("After backward: Expected (3, 5), got (%d, %d)", start, end)
	}
}

func TestEditor_SelectMultipleLines(t *testing.T) {
	editor := NewEditor("Line1\nLine2\nLine3")
	editor.SetCursorPosition(0)

	editor.MoveCursorDownWithSelection()
	editor.MoveCursorDownWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection across multiple lines")
	}

	start, end := editor.GetSelection()
	if start != 0 || end != 12 {
		t.Errorf("Expected selection (0, 12), got (%d, %d)", start, end)
	}
}

func TestEditor_Backspace_DeletesSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(0)

	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection before backspace")
	}

	editor.Backspace()

	if editor.GetText() != "World" {
		t.Errorf("Expected 'World', got '%s'", editor.GetText())
	}

	if editor.HasSelection() {
		t.Error("Expected no selection after backspace")
	}

	if editor.GetCursorPosition() != 0 {
		t.Errorf("Expected cursor at 0, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Delete_DeletesSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(6)

	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()
	editor.MoveCursorRightWithSelection()

	if !editor.HasSelection() {
		t.Error("Expected selection before delete")
	}

	editor.Delete()

	if editor.GetText() != "Hello " {
		t.Errorf("Expected 'Hello ', got '%s'", editor.GetText())
	}

	if editor.HasSelection() {
		t.Error("Expected no selection after delete")
	}

	if editor.GetCursorPosition() != 6 {
		t.Errorf("Expected cursor at 6, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Backspace_BackwardSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(6)

	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()
	editor.MoveCursorLeftWithSelection()

	editor.Backspace()

	if editor.GetText() != "World" {
		t.Errorf("Expected 'World', got '%s'", editor.GetText())
	}

	if editor.GetCursorPosition() != 0 {
		t.Errorf("Expected cursor at 0, got %d", editor.GetCursorPosition())
	}
}

func TestEditor_Delete_MultiLineSelection(t *testing.T) {
	editor := NewEditor("Line1\nLine2\nLine3")
	editor.SetCursorPosition(0)

	editor.MoveCursorDownWithSelection()
	editor.MoveCursorDownWithSelection()

	start, end := editor.GetSelection()

	editor.Delete()

	expected := "Line3"
	if editor.GetText() != expected {
		t.Errorf("Expected '%s', got '%s' (deleted range %d-%d)", expected, editor.GetText(), start, end)
	}

	if editor.HasSelection() {
		t.Error("Expected no selection after delete")
	}
}

func TestEditor_Backspace_UndoDeletesSelection(t *testing.T) {
	editor := NewEditor("Hello World")
	editor.SetCursorPosition(0)

	for i := 0; i < 6; i++ {
		editor.MoveCursorRightWithSelection()
	}

	editor.Backspace()

	if editor.GetText() != "World" {
		t.Errorf("After delete: Expected 'World', got '%s'", editor.GetText())
	}

	editor.Undo()

	if editor.GetText() != "Hello World" {
		t.Errorf("After undo: Expected 'Hello World', got '%s'", editor.GetText())
	}
}

func TestEditor_Backspace_NoSelection(t *testing.T) {
	editor := NewEditor("Hello")
	editor.SetCursorPosition(5)

	editor.Backspace()

	if editor.GetText() != "Hell" {
		t.Errorf("Expected 'Hell', got '%s'", editor.GetText())
	}
}
