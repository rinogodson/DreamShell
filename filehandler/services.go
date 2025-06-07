package filehandler

import (
	"regexp"
	"strings"
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
