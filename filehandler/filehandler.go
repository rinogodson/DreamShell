package filehandler

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(fileName string) {
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
	fmt.Fprintln(file, "Hello")
}
