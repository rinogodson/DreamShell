package filehandler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func ParseDream(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) < 2 {
		return []string{"", "", ""}
	}

	titleLine := strings.TrimSpace(lines[0])
	title := strings.TrimPrefix(titleLine, "#")
	title = strings.TrimSpace(title)

	bodyLines := lines[1 : len(lines)-1]
	body := strings.Join(bodyLines, "\n")

	tagWords := strings.Fields(lines[len(lines)-1])
	for i, tag := range tagWords {
		tagWords[i] = "#" + tag
	}
	tags := strings.Join(tagWords, " ")

	return []string{title, body, tags}
}
