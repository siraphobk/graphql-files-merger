package prompt

import "fmt"

func PrintHelpMessage() {
	msg := `GRAPHQL-FILES-MERGER merges a list of .graphql files in the directory
and renders them into a single .graphql file.

Usage:
	graphql-files-merger <input-dir-name> <output-file-name> 

Options:
	-h --help Show this message
`
	fmt.Print(msg)
}
