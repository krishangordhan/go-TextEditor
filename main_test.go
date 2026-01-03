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

func TestString_ReturnsText(t *testing.T) {
	text := "Hello World"
	pt := NewPieceTable(text)

	result := pt.String()
	if result != text {
		t.Errorf("Expected String() to return %q, got %q", text, result)
	}
}

func TestString_ReturnsEmpty(t *testing.T) {
	pt := NewPieceTable("")

	result := pt.String()
	if result != "" {
		t.Errorf("Expected String() to return empty string, got %q", result)
	}
}

func TestString_ReturnsUnicode(t *testing.T) {
	text := "Hello 世界 🌍"
	pt := NewPieceTable(text)

	result := pt.String()
	if result != text {
		t.Errorf("Expected String() to return %q, got %q", text, result)
	}
}
