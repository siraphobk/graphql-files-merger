package dirman

// Directory management package

import (
	"io/ioutil"
	"os"
)

func GetCwd() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return cwd, nil
}

func ReadDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// get dir names
	dirNames := []string{}
	for _, file := range files {
		dirNames = append(dirNames, file.Name())
	}
	return dirNames, nil
}
