package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileManager_NewFileManager(t *testing.T) {
	fm := NewFileManager()
	if fm.GetFilePath() != "" {
		t.Errorf("Expected empty file path, got %s", fm.GetFilePath())
	}
	if fm.IsDirty() {
		t.Error("Expected isDirty to be false")
	}
}

func TestFileManager_NewFileManagerWithPath(t *testing.T) {
	path := "test.txt"
	fm := NewFileManagerWithPath(path)
	if fm.GetFilePath() != path {
		t.Errorf("Expected file path %s, got %s", path, fm.GetFilePath())
	}
	if fm.IsDirty() {
		t.Error("Expected isDirty to be false")
	}
}

func TestFileManager_SetFilePath_SetsFilePath(t *testing.T) {
	fm := NewFileManager()
	path := "new_file.txt"
	fm.SetFilePath(path)
	if fm.GetFilePath() != path {
		t.Errorf("Expected file path %s, got %s", path, fm.GetFilePath())
	}
}

func TestFileManager_MarkDirtyAndClean_FileIsDirty(t *testing.T) {
	fm := NewFileManager()
	if fm.IsDirty() {
		t.Error("Expected isDirty to be false initially")
	}

	fm.MarkDirty()
	if !fm.IsDirty() {
		t.Error("Expected isDirty to be true after MarkDirty")
	}

	fm.MarkClean()
	if fm.IsDirty() {
		t.Error("Expected isDirty to be false after MarkClean")
	}
}

func TestFileManager_HasFile_Success(t *testing.T) {
	fm := NewFileManager()
	if fm.HasFile() {
		t.Error("Expected HasFile to be false with empty path")
	}

	fm.SetFilePath("test.txt")
	if !fm.HasFile() {
		t.Error("Expected HasFile to be true with non-empty path")
	}
}

func TestFileManager_ReadFile_Success(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "Hello, World!\nThis is a test file."

	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fm := NewFileManagerWithPath(tmpFile)
	readContent, err := fm.ReadFile()
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if readContent != content {
		t.Errorf("Expected content %q, got %q", content, readContent)
	}
}

func TestFileManager_ReadFile_NoPath(t *testing.T) {
	fm := NewFileManager()
	_, err := fm.ReadFile()
	if err == nil {
		t.Error("Expected error when reading with no file path set")
	}
}

func TestFileManager_ReadFile_NonExistent(t *testing.T) {
	fm := NewFileManagerWithPath("/nonexistent/path/file.txt")
	_, err := fm.ReadFile()
	if err == nil {
		t.Error("Expected error when reading non-existent file")
	}
}

func TestFileManager_WriteFile_Success(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "Test content\nWith multiple lines."

	fm := NewFileManagerWithPath(tmpFile)
	fm.MarkDirty()

	err := fm.WriteFile(content)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	readContent, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}

	if string(readContent) != content {
		t.Errorf("Expected content %q, got %q", content, string(readContent))
	}

	if fm.IsDirty() {
		t.Error("Expected isDirty to be false after successful write")
	}
}

func TestFileManager_WriteFile_NoPath(t *testing.T) {
	fm := NewFileManager()
	err := fm.WriteFile("test content")
	if err == nil {
		t.Error("Expected error when writing with no file path set")
	}
}

func TestFileManager_WriteFile_InvalidPath(t *testing.T) {
	fm := NewFileManagerWithPath("/invalid/path/that/does/not/exist/file.txt")
	err := fm.WriteFile("test content")
	if err == nil {
		t.Error("Expected error when writing to invalid path")
	}
}

func TestFileManager_ReadWriteRoundTrip(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "roundtrip.txt")
	originalContent := "Original content\nWith newlines\nAnd special chars: @#$%"

	fm := NewFileManagerWithPath(tmpFile)
	err := fm.WriteFile(originalContent)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	readContent, err := fm.ReadFile()
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if readContent != originalContent {
		t.Errorf("Round-trip failed. Expected %q, got %q", originalContent, readContent)
	}
}
