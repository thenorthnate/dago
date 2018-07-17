package dago

import (
	// "fmt"
	"encoding/csv"
	"io/ioutil"
	"strings"
)

// ReadFile : Returns a DataFrame with all the data from the file
func ReadFile(filePath string, fileFormat string) DataFrame {
	DF := DataFrame{}
	byteData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return DF
	}
	switch fileFormat {
	case "csv":
		DF = readCSV(byteData, true)
	case "json":
		DF = readJSON(byteData)
	}
	return DF
}

func readCSV(data []byte, header bool) DataFrame {
	DF := DataFrame{level: true}
	stringData := string(data)
	r := csv.NewReader(strings.NewReader(stringData))
	records, err := r.ReadAll()
	if err != nil {
		return DF
	}

	tRecords := map[string][]string{} // transposed records
	headers := []string{}
	if header {
		for _, v := range records[0] {
			headers = append(headers, strings.Trim(v, " "))
		}
		records = records[1:]
	} else {
		for i := 0; i < len(records[0]); i++ {
			headers = append(headers, "")
		}
	}
	recLen := len(records)
	for _, v := range headers {
		tRecords[v] = make([]string, recLen)
	}

	for i, v := range records {
		for j, w := range v {
			tRecords[headers[j]][i] = w
		}
	}
	DF = New(tRecords)
	return DF
}

func readJSON(data []byte) DataFrame {
	DF := DataFrame{}
	return DF
}
