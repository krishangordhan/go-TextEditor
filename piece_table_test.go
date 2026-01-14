package main

import (
	"testing"
)

func TestNewPieceTable_WithText(t *testing.T) {
	text := "Hello World"
	pt := NewPieceTable(text)

	if string(pt.original) != text {
		t.Errorf("Expected original buffer to be %q, got %q", text, string(pt.original))
	}

	if len(pt.add) != 0 {
		t.Errorf("Expected add buffer to be empty, got length %d", len(pt.add))
	}

	if len(pt.pieces) != 1 {
		t.Errorf("Expected 1 piece, got %d", len(pt.pieces))
	}

	piece := pt.pieces[0]
	if piece.bufferType != Original {
		t.Errorf("Expected piece to reference Original buffer, got %v", piece.bufferType)
	}

	if piece.start != 0 {
		t.Errorf("Expected piece to start at 0, got %d", piece.start)
	}

	if piece.length != len(pt.original) {
		t.Errorf("Expected piece length to be %d, got %d", len(pt.original), piece.length)
	}
}

func TestNewPieceTable_EmptyText(t *testing.T) {
	pt := NewPieceTable("")

	if len(pt.original) != 0 {
		t.Errorf("Expected original buffer to be empty, got length %d", len(pt.original))
	}

	if len(pt.add) != 0 {
		t.Errorf("Expected add buffer to be empty, got length %d", len(pt.add))
	}

	if len(pt.pieces) != 0 {
		t.Errorf("Expected 0 pieces, got %d", len(pt.pieces))
	}
}

func TestNewPieceTable_UnicodeText(t *testing.T) {
	text := "Hello 世界 🌍"
	pt := NewPieceTable(text)

	if string(pt.original) != text {
		t.Errorf("Expected original buffer to be %q, got %q", text, string(pt.original))
	}

	expectedRuneCount := 10
	if len(pt.original) != expectedRuneCount {
		t.Errorf("Expected %d runes, got %d", expectedRuneCount, len(pt.original))
	}

	if pt.pieces[0].length != expectedRuneCount {
		t.Errorf("Expected piece length to be %d, got %d", expectedRuneCount, pt.pieces[0].length)
	}
}

func TestPieceTable_String(t *testing.T) {
	text := "Hello World"
	pt := NewPieceTable(text)

	result := pt.String()
	if result != text {
		t.Errorf("Expected String() to return %q, got %q", text, result)
	}
}

func TestPieceTable_String_Empty(t *testing.T) {
	pt := NewPieceTable("")

	result := pt.String()
	if result != "" {
		t.Errorf("Expected String() to return empty string, got %q", result)
	}
}

func TestPieceTable_String_Unicode(t *testing.T) {
	text := "Hello 世界 🌍"
	pt := NewPieceTable(text)

	result := pt.String()
	if result != text {
		t.Errorf("Expected String() to return %q, got %q", text, result)
	}
}

func TestPieceTable_Length(t *testing.T) {
	pt := NewPieceTable("Hello")

	if pt.Length() != 5 {
		t.Errorf("Expected length 5, got %d", pt.Length())
	}
}

func TestPieceTable_Length_AfterInsert(t *testing.T) {
	pt := NewPieceTable("Hello")
	pt.Insert(5, " World")

	if pt.Length() != 11 {
		t.Errorf("Expected length 11, got %d", pt.Length())
	}
}

func TestPieceTable_Length_AfterDelete(t *testing.T) {
	pt := NewPieceTable("Hello World")
	pt.Delete(5, 6)

	if pt.Length() != 5 {
		t.Errorf("Expected length 5, got %d", pt.Length())
	}
}

func TestPieceTable_Insert_AtBeginning(t *testing.T) {
	pt := NewPieceTable("World")
	pt.Insert(0, "Hello ")

	result := pt.String()
	expected := "Hello World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Insert_AtEnd(t *testing.T) {
	pt := NewPieceTable("Hello")
	pt.Insert(5, " World")

	result := pt.String()
	expected := "Hello World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Insert_InMiddle(t *testing.T) {
	pt := NewPieceTable("Hello World")
	pt.Insert(6, "Beautiful ")

	result := pt.String()
	expected := "Hello Beautiful World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Insert_Multiple(t *testing.T) {
	pt := NewPieceTable("ac")
	pt.Insert(1, "b")
	pt.Insert(3, "d")

	result := pt.String()
	expected := "abcd"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Insert_EmptyString(t *testing.T) {
	pt := NewPieceTable("Hello")
	pt.Insert(2, "")

	result := pt.String()
	expected := "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Insert_Unicode(t *testing.T) {
	pt := NewPieceTable("Hello ")
	pt.Insert(6, "世界 🌍")

	result := pt.String()
	expected := "Hello 世界 🌍"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Delete_FromBeginning(t *testing.T) {
	pt := NewPieceTable("Hello World")
	pt.Delete(0, 6)

	result := pt.String()
	expected := "World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Delete_FromEnd(t *testing.T) {
	pt := NewPieceTable("Hello World")
	pt.Delete(5, 6)

	result := pt.String()
	expected := "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Delete_FromMiddle(t *testing.T) {
	pt := NewPieceTable("Hello Beautiful World")
	pt.Delete(6, 10)

	result := pt.String()
	expected := "Hello World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Delete_EntireText(t *testing.T) {
	pt := NewPieceTable("Hello")
	pt.Delete(0, 5)

	result := pt.String()
	expected := ""
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Delete_PartialPiece(t *testing.T) {
	pt := NewPieceTable("Hello World")
	pt.Delete(2, 3)

	result := pt.String()
	expected := "He World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Delete_ZeroLength(t *testing.T) {
	pt := NewPieceTable("Hello")
	pt.Delete(2, 0)

	result := pt.String()
	expected := "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Delete_AfterInsert(t *testing.T) {
	pt := NewPieceTable("Hello World")
	pt.Insert(6, "Beautiful ")
	pt.Delete(0, 6)

	result := pt.String()
	expected := "Beautiful World"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestPieceTable_Substring_SuccessShorterMessage(t *testing.T) {
	pt := NewPieceTable("Hello World")

	result := pt.Substring(0, 5)
	if result != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", result)
	}

	result = pt.Substring(6, 11)
	if result != "World" {
		t.Errorf("Expected 'World', got '%s'", result)
	}

	result = pt.Substring(0, 11)
	if result != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", result)
	}
}

func TestPieceTable_Substring_WithInsert(t *testing.T) {
	pt := NewPieceTable("Hello World")
	pt.Insert(5, " Beautiful")

	result := pt.Substring(5, 15)
	if result != " Beautiful" {
		t.Errorf("Expected ' Beautiful', got '%s'", result)
	}

	result = pt.Substring(0, 21)
	if result != "Hello Beautiful World" {
		t.Errorf("Expected 'Hello Beautiful World', got '%s'", result)
	}
}

func TestPieceTable_Substring_WithDelete(t *testing.T) {
	pt := NewPieceTable("Hello Beautiful World")
	pt.Delete(6, 10)

	result := pt.Substring(0, 11)
	if result != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", result)
	}
}

func TestPieceTable_Substring_InvalidRange(t *testing.T) {
	pt := NewPieceTable("Hello")

	result := pt.Substring(-1, 5)
	if result != "" {
		t.Errorf("Expected empty string for negative start, got '%s'", result)
	}

	result = pt.Substring(0, 100)
	if result != "" {
		t.Errorf("Expected empty string for end beyond length, got '%s'", result)
	}

	result = pt.Substring(5, 2)
	if result != "" {
		t.Errorf("Expected empty string for start > end, got '%s'", result)
	}
}
