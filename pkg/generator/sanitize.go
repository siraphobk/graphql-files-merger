package generator

import "strings"

// Sanitize remove a string from the content given
func Sanitize(content string, toSanitize string) string {
	sanitizedContent := ""
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		ind := strings.Index(line, toSanitize)
		if ind != -1 { // no substring
			if ind == 0 { // substring at index 0
				line = line[7:] // remove toSanitize from the content
			}
		}
		sanitizedContent += line + "\n"
	}

	return sanitizedContent
}
