package generator

import (
	"regexp"
	"strings"
)

func FindQuery(content string) (string, []string) {
	queryList := []string{}

	regexpQuery := regexp.MustCompile(`(?s)type Query \{(.*?)\}`)
	q := regexpQuery.FindString(content)
	if len(q) > 0 {
		s := strings.Split(q, "\n")
		for i := 1; i < len(s)-1; i++ {
			queryList = append(queryList, s[i]) // put to list
		}
		content = strings.ReplaceAll(content, q, "") // remove from content
		content = strings.ReplaceAll(content, "\n\n", "\n")
	}

	return content, queryList
}

func FindMutation(content string) (string, []string) {
	mutationList := []string{}

	regexpMutation := regexp.MustCompile(`(?s)type Mutation \{(.*?)\}`)
	m := regexpMutation.FindString(content)
	if len(m) > 0 {
		s := strings.Split(m, "\n")
		for i := 1; i < len(s)-1; i++ {
			mutationList = append(mutationList, s[i]) // put to list
		}
		content = strings.ReplaceAll(content, m, "") // remove from content
		content = strings.ReplaceAll(content, "\n\n", "\n")
	}

	return content, mutationList
}
