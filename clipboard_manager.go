package main

import "github.com/atotto/clipboard"

type Clipboard interface {
	Copy(text string) error
	Paste() (string, error)
}

type ClipboardManager struct{}

func NewClipboardManager() *ClipboardManager {
	return &ClipboardManager{}
}

func (cm *ClipboardManager) Copy(text string) error {
	return clipboard.WriteAll(text)
}

func (cm *ClipboardManager) Paste() (string, error) {
	return clipboard.ReadAll()
}
