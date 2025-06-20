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

	fullPath := filepath.Join(dirPath, fileName+".md")
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

	dateFields := strings.Fields(dateStr)
	if len(dateFields) < 5 {
		fmt.Println("Date string format is invalid")
		return []string{title, ""}
	}

	layout := "2006-01-02 15:04:05.999999 -0700"
	datetimeStr := fmt.Sprintf("%s %s %s", dateFields[0], dateFields[1], dateFields[2])
	t, err := time.Parse(layout, datetimeStr)
	if err != nil {
		fmt.Println("Parse error:", err)
		return []string{title, ""}
	}

	day := fmt.Sprintf("%02d", t.Day())
	month := t.Month().String()
	year := fmt.Sprintf("%d", t.Year())
	hour := t.Format("03")
	minute := t.Format("04")
	second := t.Format("05")
	ampm := t.Format("PM")
	timeStr := fmt.Sprintf("%s:%s:%s", hour, minute, second)
	timezone := dateFields[3]
	offset := dateFields[2]

	result := []string{day, month, year, timeStr, ampm, timezone, offset}
	dateString := strings.Join(result, " ")
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
		createdI, err := time.Parse(layout, dateStrI)
		if err != nil {
			panic(err)
		}
		createdJ, err := time.Parse(layout, dateStrJ)
		if err != nil {
			panic(err)
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
