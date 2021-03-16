package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/copsterr/graphql-files-merger/pkg/dirman"
	"github.com/copsterr/graphql-files-merger/pkg/generator"
	"github.com/copsterr/graphql-files-merger/pkg/prompt"
)

func main() {
	// get arguments
	inputDirName, outputDirName := argsParser()

	// get current working directory
	cwd, err := dirman.GetCwd()
	if err != nil {
		log.Fatalf("Cannot get cwd: %v", err.Error())
	}

	// read graphql file list
	graphqlFiles, err := dirman.ReadDir(path.Join(cwd, inputDirName))
	if err != nil {
		log.Fatalf("Cannot read dir: %v", err.Error())
	}

	// create output file
	schemasFile, err := os.Create(outputDirName)
	if err != nil {
		log.Fatalf("Cannot create schemas.graphql %v", err.Error())
	}
	defer schemasFile.Close()

	allQueries := []string{}
	allMutations := []string{}

	// loop through each file
	for _, graphqlFile := range graphqlFiles {
		// read content of a file in bytes
		bContent, err := ioutil.ReadFile(path.Join(cwd, inputDirName, graphqlFile))
		if err != nil {
			log.Fatalf("Cannot read content of %s, error: %s", graphqlFile, err.Error())
		}

		// convert bytes to string
		content := string(bContent)
		sanitizedContent := generator.Sanitize(content, "extend ")

		// find type Query (Won't ignore comments)
		sanitizedContent, queries := generator.FindQuery(sanitizedContent)
		allQueries = append(allQueries, queries...)

		// find type Mutation (Won't ignore comments)
		sanitizedContent, mutations := generator.FindMutation(sanitizedContent)
		allMutations = append(allMutations, mutations...)

		// At this point, sanitizedContent doesn't contain any queries or mutations.
		// only type and input are written to the output
		schemasFile.WriteString(sanitizedContent)
	} // end loop through each file

	// write queries
	schemasFile.WriteString("type Query {\n")
	for _, query := range allQueries {
		schemasFile.WriteString(query + "\n")
	}
	schemasFile.WriteString("}\n")

	// write mutations
	schemasFile.WriteString("type Mutation {\n")
	for _, mutation := range allMutations {
		schemasFile.WriteString(mutation + "\n")
	}
	schemasFile.WriteString("}\n")
}

// argsParser returns 2 arguments from stdin
func argsParser() (string, string) {
	// get dir args
	args := os.Args[1:]

	if len(args) == 0 {
		prompt.PrintHelpMessage()
		os.Exit(0)
	}

	inputDirName := args[0]
	if inputDirName == "-h" || inputDirName == "--help" {
		prompt.PrintHelpMessage()
		os.Exit(0)
	}

	outputDirName := args[1]

	return inputDirName, outputDirName
}
