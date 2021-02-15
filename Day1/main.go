package main

import (
	"SoftwareGoDay1/data"
	"SoftwareGoDay1/humanity"
)

func main() {
	data.ReadFile("file.csv")
	humanity.NewHumansFromCsvFile("file.csv")
}
