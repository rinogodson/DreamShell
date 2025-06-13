package filehandler

import (
	"regexp"
	"strings"
	"time"
)

func TagValidator(tagString string) bool {
	tagString = strings.TrimSpace(tagString)
	if tagString == "" {
		return false
	}
	validPattern := regexp.MustCompile(`^(#\w+)(\s+#\w+)*$`)
	return validPattern.MatchString(tagString)
}

func ExtractTags(content string) []string {
	tagPattern := regexp.MustCompile(`#\w+`)
	tags := tagPattern.FindAllString(content, -1)
	for i, tag := range tags {
		tags[i] = strings.TrimPrefix(tag, "#")
	}
	return tags
}

func ParseInput(input string) []string {
	if len(input) == 1 {
		return []string{"", ""}
	}
	parts := strings.Split(input, "~")
	title := parts[0]
	dateStr := parts[1]
	dateStr = strings.TrimSuffix(dateStr, ".md")

	parseLayout := "2006-01-02 15:04:05.999999 -0700 MST m=+999.999999"
	t, _ := time.Parse(parseLayout, dateStr)

	outputLayout := "January, 2, 2025, 03:04 PM"
	formattedDate := t.Format(outputLayout)

	return []string{title, formattedDate}
}

