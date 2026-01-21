package main

import (
	"testing"
)

type MockClipboard struct {
	content string
}

func (mc *MockClipboard) Copy(text string) error {
	mc.content = text
	return nil
}

func (mc *MockClipboard) Paste() (string, error) {
	return mc.content, nil
}

func TestClipboardManager_NewClipboardManager(t *testing.T) {
	cm := NewClipboardManager()
	if cm == nil {
		t.Error("Expected clipboard manager to be created")
	}
}

func TestClipboardManager_CopyAndPaste_PastesSuccessfully(t *testing.T) {
	var cm Clipboard = &MockClipboard{}
	testText := "Hello, clipboard!"

	err := cm.Copy(testText)
	if err != nil {
		t.Fatalf("Failed to copy to clipboard: %v", err)
	}

	result, err := cm.Paste()
	if err != nil {
		t.Fatalf("Failed to paste from clipboard: %v", err)
	}

	if result != testText {
		t.Errorf("Expected '%s', got '%s'", testText, result)
	}
}

func TestClipboardManager_CopyEmptyString_PastesEmptyString(t *testing.T) {
	var cm Clipboard = &MockClipboard{}

	err := cm.Copy("")
	if err != nil {
		t.Fatalf("Failed to copy empty string to clipboard: %v", err)
	}

	result, err := cm.Paste()
	if err != nil {
		t.Fatalf("Failed to paste from clipboard: %v", err)
	}

	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestClipboardManager_MultipleOperations_PastesSuccessfully(t *testing.T) {
	var cm Clipboard = &MockClipboard{}

	err := cm.Copy("First")
	if err != nil {
		t.Fatalf("Failed first copy: %v", err)
	}

	result, err := cm.Paste()
	if err != nil {
		t.Fatalf("Failed first paste: %v", err)
	}
	if result != "First" {
		t.Errorf("First paste: expected 'First', got '%s'", result)
	}

	err = cm.Copy("Second")
	if err != nil {
		t.Fatalf("Failed second copy: %v", err)
	}

	result, err = cm.Paste()
	if err != nil {
		t.Fatalf("Failed second paste: %v", err)
	}
	if result != "Second" {
		t.Errorf("Second paste: expected 'Second', got '%s'", result)
	}
}
