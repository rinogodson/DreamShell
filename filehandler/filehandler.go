package filehandler

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(fileName string, content string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dirPath := filepath.Join(home, ".dreamshell", "dreams")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fullPath := filepath.Join(dirPath, fileName+".md")
	file, err := os.Create(fullPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintln(file, content)
}

func GetFiles() []os.DirEntry {
	home, err := os.UserHomeDir()
	files, err := os.ReadDir(filepath.Join(home, ".dreamshell", "dreams"))
	if err != nil {
		panic(err)
	}
	return files
}

func GetContent(input string) string {
  textContent, err := os.ReadFile(input) 
	if err != nil {
		panic(err)
	}
	return string(textContent)
}
