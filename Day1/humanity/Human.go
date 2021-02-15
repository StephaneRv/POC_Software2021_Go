package humanity

import (
	"SoftwareGoDay1/data"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Human struct {
	Name    string
	Age     int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	var Err error
	h := Human{Name: csv[0], Country: csv[2]}
	h.Age, Err = strconv.Atoi((csv[1]))
	return &h, Err
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	var tab []*Human
	_data, err := data.ReadFile(path)

	if err != nil {
		return nil, err
	}
	for i := 0; i < len(_data); i++ {
		csv, csvErr := data.LineToCSV(_data[i])
		if len(csv) != 3 || csvErr != nil {
			return nil, csvErr
		}
		_Age, ageErr := strconv.Atoi(csv[1])
		if ageErr != nil {
			return nil, ageErr
		}
		humanLine := &Human{Name: csv[0], Age: _Age, Country: csv[0]}
		tab = append(tab, humanLine)
	}
	return tab, nil
}

func NewHumansFromJsonFile(path string) ([]*Human, error) {
	var tab []*Human

	_data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(_data, &tab)
	return tab, nil
}
