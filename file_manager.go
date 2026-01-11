package main

import (
	"errors"
	"os"
)

type FileManager struct {
	filePath string
	isDirty  bool
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func NewFileManagerWithPath(path string) *FileManager {
	return &FileManager{
		filePath: path,
	}
}

func (fm *FileManager) GetFilePath() string {
	return fm.filePath
}

func (fm *FileManager) SetFilePath(path string) {
	fm.filePath = path
}

func (fm *FileManager) IsDirty() bool {
	return fm.isDirty
}

func (fm *FileManager) MarkDirty() {
	fm.isDirty = true
}

func (fm *FileManager) MarkClean() {
	fm.isDirty = false
}

func (fm *FileManager) HasFile() bool {
	return fm.filePath != ""
}

func (fm *FileManager) ReadFile() (string, error) {
	if fm.filePath == "" {
		return "", errors.New("no file path set")
	}

	data, err := os.ReadFile(fm.filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (fm *FileManager) WriteFile(content string) error {
	if fm.filePath == "" {
		return errors.New("no file path set")
	}

	err := os.WriteFile(fm.filePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	fm.MarkClean()
	return nil
}
