package main

import "testing"

func TestCommand_NewInsertCommand(t *testing.T) {
	buffer := NewPieceTable("Hello World")
	cursor := NewCursor()
	cursor.SetPosition(5)

	cmd := NewInsertCommand(buffer, cursor, ", there", 5)

	cmd.Execute()

	if buffer.String() != "Hello, there World" {
		t.Errorf("Expected 'Hello, there World', got '%s'", buffer.String())
	}

	if cursor.GetPosition() != 12 {
		t.Errorf("Expected cursor at 12, got %d", cursor.GetPosition())
	}

	cmd.Undo()

	if buffer.String() != "Hello World" {
		t.Errorf("After undo, expected 'Hello World', got '%s'", buffer.String())
	}

	if cursor.GetPosition() != 5 {
		t.Errorf("After undo, expected cursor at 5, got %d", cursor.GetPosition())
	}
}

func TestCommand_NewDeleteCommand(t *testing.T) {
	buffer := NewPieceTable("Hello World")
	cursor := NewCursor()
	cursor.SetPosition(5)

	cmd := NewDeleteCommand(buffer, cursor, 5, 6)

	cmd.Execute()

	if buffer.String() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", buffer.String())
	}

	if cursor.GetPosition() != 5 {
		t.Errorf("Expected cursor at 5, got %d", cursor.GetPosition())
	}

	cmd.Undo()

	if buffer.String() != "Hello World" {
		t.Errorf("After undo, expected 'Hello World', got '%s'", buffer.String())
	}

	if cursor.GetPosition() != 5 {
		t.Errorf("After undo, expected cursor at 5, got %d", cursor.GetPosition())
	}
}

func TestCommand_InsertCommandAtBeginning(t *testing.T) {
	buffer := NewPieceTable("World")
	cursor := NewCursor()
	cursor.SetPosition(0)

	cmd := NewInsertCommand(buffer, cursor, "Hello ", 0)
	cmd.Execute()

	if buffer.String() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", buffer.String())
	}

	cmd.Undo()

	if buffer.String() != "World" {
		t.Errorf("After undo, expected 'World', got '%s'", buffer.String())
	}
}

func TestCommand_DeleteCommandAtBeginning(t *testing.T) {
	buffer := NewPieceTable("Hello World")
	cursor := NewCursor()
	cursor.SetPosition(0)

	cmd := NewDeleteCommand(buffer, cursor, 0, 6)
	cmd.Execute()

	if buffer.String() != "World" {
		t.Errorf("Expected 'World', got '%s'", buffer.String())
	}

	cmd.Undo()

	if buffer.String() != "Hello World" {
		t.Errorf("After undo, expected 'Hello World', got '%s'", buffer.String())
	}
}

func TestCommand_MultipleInsertCommands(t *testing.T) {
	buffer := NewPieceTable("")
	cursor := NewCursor()

	cmd1 := NewInsertCommand(buffer, cursor, "Hello", 0)
	cmd1.Execute()

	cmd2 := NewInsertCommand(buffer, cursor, " ", 5)
	cmd2.Execute()

	cmd3 := NewInsertCommand(buffer, cursor, "World", 6)
	cmd3.Execute()

	if buffer.String() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", buffer.String())
	}

	cmd3.Undo()
	if buffer.String() != "Hello " {
		t.Errorf("After undo 3, expected 'Hello ', got '%s'", buffer.String())
	}

	cmd2.Undo()
	if buffer.String() != "Hello" {
		t.Errorf("After undo 2, expected 'Hello', got '%s'", buffer.String())
	}

	cmd1.Undo()
	if buffer.String() != "" {
		t.Errorf("After undo 1, expected '', got '%s'", buffer.String())
	}
}

func TestCommand_MultipleDeleteCommands(t *testing.T) {
	buffer := NewPieceTable("Hello World")
	cursor := NewCursor()

	cmd1 := NewDeleteCommand(buffer, cursor, 5, 1)
	cmd1.Execute()

	if buffer.String() != "HelloWorld" {
		t.Errorf("Expected 'HelloWorld', got '%s'", buffer.String())
	}

	cmd2 := NewDeleteCommand(buffer, cursor, 5, 5)
	cmd2.Execute()

	if buffer.String() != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", buffer.String())
	}

	cmd2.Undo()
	if buffer.String() != "HelloWorld" {
		t.Errorf("After undo 2, expected 'HelloWorld', got '%s'", buffer.String())
	}

	cmd1.Undo()
	if buffer.String() != "Hello World" {
		t.Errorf("After undo 1, expected 'Hello World', got '%s'", buffer.String())
	}
}
