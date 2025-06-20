package filehandler

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
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
	timestamp := time.Now().Format("2006-01-02 15:04:05.999999 -0700")
	fullPath := filepath.Join(dirPath, fmt.Sprintf("%s~%s.md", fileName, timestamp))
	file, err := os.Create(fullPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintln(file, content)
}

func localParseInput(input string) []string {
	parts := strings.SplitN(input, "~", 2)
	if len(parts) != 2 {
		return []string{"", ""}
	}
	title := parts[0]
	dateStr := strings.TrimSuffix(parts[1], ".md")
	layout := "2006-01-02 15:04:05.999999 -0700"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Parse error:", err)
		return []string{title, ""}
	}
	dateString := t.Format("02 January 2006 03:04:05 PM MST -0700")
	return []string{title, dateString}
}

func GetFiles() []os.DirEntry {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Join(home, ".dreamshell", "dreams")
	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	sort.Slice(files, func(i, j int) bool {
		filePathI := filepath.Join(dirPath, files[i].Name())
		filePathJ := filepath.Join(dirPath, files[j].Name())
		dateStrI := localParseInput(filePathI)[1]
		dateStrJ := localParseInput(filePathJ)[1]
		const layout = "02 January 2006 03:04:05 PM MST -0700"
		var createdI, createdJ time.Time
		if dateStrI != "" {
			createdI, err = time.Parse(layout, dateStrI)
			if err != nil {
				fmt.Println("Parse error for", filePathI, ":", err)
				return true
			}
		}
		if dateStrJ != "" {
			createdJ, err = time.Parse(layout, dateStrJ)
			if err != nil {
				fmt.Println("Parse error for", filePathJ, ":", err)
				return false
			}
		}
		return createdI.Before(createdJ)
	})
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
