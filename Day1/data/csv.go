package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func LineToCSV(line string) ([]string, error) {
	splitFile := strings.Split(line, ",")

	if len(splitFile) == 0 {
		return nil, fmt.Errorf("Bad file format")
	}
	return splitFile, nil
}

func ReadFile(path string) ([]string, error) {
	content, err := ioutil.ReadFile("file.csv")

	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(content)
	lines := strings.Split(fileContent, "\n")
	return lines, err
}
