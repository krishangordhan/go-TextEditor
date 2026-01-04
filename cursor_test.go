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
