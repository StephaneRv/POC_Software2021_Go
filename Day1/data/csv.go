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
	fmt.Println((string(content)))
	fileContent := string(content)
	splitFile, err := LineToCSV(fileContent)
	return splitFile, err
}
