package main

import (
	"testing"
)

func TestCursor_NewCursor(t *testing.T) {
	cursor := NewCursor()

	if cursor.GetPosition() != 0 {
		t.Errorf("Expected new cursor at position 0, got %d", cursor.GetPosition())
	}
}

func TestCursor_SetPosition(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)

	if cursor.GetPosition() != 5 {
		t.Errorf("Expected cursor at position 5, got %d", cursor.GetPosition())
	}
}

func TestCursor_SetPosition_Negative(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(-5)

	if cursor.GetPosition() != 0 {
		t.Errorf("Expected cursor to clamp to 0, got %d", cursor.GetPosition())
	}
}

func TestCursor_MoveLeft(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)
	cursor.MoveLeft()

	if cursor.GetPosition() != 4 {
		t.Errorf("Expected cursor at position 4, got %d", cursor.GetPosition())
	}
}

func TestCursor_MoveLeft_AtStart(t *testing.T) {
	cursor := NewCursor()
	cursor.MoveLeft()

	if cursor.GetPosition() != 0 {
		t.Errorf("Expected cursor to stay at 0, got %d", cursor.GetPosition())
	}
}

func TestCursor_MoveRight(t *testing.T) {
	cursor := NewCursor()
	cursor.MoveRight(10)

	if cursor.GetPosition() != 1 {
		t.Errorf("Expected cursor at position 1, got %d", cursor.GetPosition())
	}
}

func TestCursor_MoveRight_AtEnd(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(10)
	cursor.MoveRight(10)

	if cursor.GetPosition() != 10 {
		t.Errorf("Expected cursor to stay at 10, got %d", cursor.GetPosition())
	}
}

func TestCursor_NewCursor_NoSelection(t *testing.T) {
	cursor := NewCursor()

	if cursor.HasSelection() {
		t.Error("Expected no selection on new cursor")
	}

	if cursor.GetSelectionAnchor() != -1 {
		t.Errorf("Expected selection anchor -1, got %d", cursor.GetSelectionAnchor())
	}
}

func TestCursor_StartSelection(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)
	cursor.StartSelection()

	if cursor.GetSelectionAnchor() != 5 {
		t.Errorf("Expected selection anchor at 5, got %d", cursor.GetSelectionAnchor())
	}

	cursor.SetPosition(10)
	if !cursor.HasSelection() {
		t.Error("Expected selection after moving cursor")
	}
}

func TestCursor_ClearSelection(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)
	cursor.StartSelection()
	cursor.SetPosition(10)

	if !cursor.HasSelection() {
		t.Error("Expected selection before clear")
	}

	cursor.ClearSelection()

	if cursor.HasSelection() {
		t.Error("Expected no selection after clear")
	}

	if cursor.GetSelectionAnchor() != -1 {
		t.Errorf("Expected selection anchor -1 after clear, got %d", cursor.GetSelectionAnchor())
	}
}

func TestCursor_HasSelection_NoMovement(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)
	cursor.StartSelection()

	if cursor.HasSelection() {
		t.Error("Expected no selection when cursor hasn't moved from anchor")
	}
}

func TestCursor_GetSelection_ForwardSelection(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)
	cursor.StartSelection()
	cursor.SetPosition(10)

	start, end := cursor.GetSelection()
	if start != 5 || end != 10 {
		t.Errorf("Expected selection (5, 10), got (%d, %d)", start, end)
	}
}

func TestCursor_GetSelection_BackwardSelection(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(10)
	cursor.StartSelection()
	cursor.SetPosition(5)

	start, end := cursor.GetSelection()
	if start != 5 || end != 10 {
		t.Errorf("Expected selection (5, 10), got (%d, %d)", start, end)
	}
}

func TestCursor_GetSelection_NoSelection(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(7)

	start, end := cursor.GetSelection()
	if start != 7 || end != 7 {
		t.Errorf("Expected selection (7, 7), got (%d, %d)", start, end)
	}
}

func TestCursor_GetSelection_SingleCharacter(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)
	cursor.StartSelection()
	cursor.MoveRight(10)

	start, end := cursor.GetSelection()
	if start != 5 || end != 6 {
		t.Errorf("Expected selection (5, 6), got (%d, %d)", start, end)
	}
}

func TestCursor_Selection_MoveLeftFromAnchor(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(5)
	cursor.StartSelection()
	cursor.MoveLeft()
	cursor.MoveLeft()

	if !cursor.HasSelection() {
		t.Error("Expected selection after moving left")
	}

	start, end := cursor.GetSelection()
	if start != 3 || end != 5 {
		t.Errorf("Expected selection (3, 5), got (%d, %d)", start, end)
	}
}

func TestCursor_Selection_ExtendBothDirections(t *testing.T) {
	cursor := NewCursor()
	cursor.SetPosition(10)
	cursor.StartSelection()

	cursor.SetPosition(15)
	start, end := cursor.GetSelection()
	if start != 10 || end != 15 {
		t.Errorf("After right move: Expected (10, 15), got (%d, %d)", start, end)
	}

	cursor.SetPosition(5)
	start, end = cursor.GetSelection()
	if start != 5 || end != 10 {
		t.Errorf("After left move: Expected (5, 10), got (%d, %d)", start, end)
	}
}
