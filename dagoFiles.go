package dago

import (
	// "fmt"
	"encoding/csv"
	"io/ioutil"
	"strings"
	"unicode"
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
	// TODO: Check if each header character is a letter or not using unicode.IsLetter() bool
	// must loop through a string to get a rune value for each input to IsLetter
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
			headers = append(headers, formatStringName(v))
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

func formatStringName(name string) string {
	name = strings.TrimSpace(name)
	outString := ""
	for _, runeVal := range name {
		if unicode.IsLetter(runeVal) || unicode.IsNumber(runeVal) {
			outString += string(runeVal)
		}
	}
	return outString
}

/*
file, err := os.Open("filetoread.txt")
if err != nil {
  fmt.Println(err)
  return
}
defer file.Close()

scanner := bufio.NewScanner(file)
scanner.Split(bufio.ScanLines)

// This is our buffer now
var lines []string

for scanner.Scan() {
  lines = append(lines, scanner.Text())
}

fmt.Println("read lines:")
for _, line := range lines {
  fmt.Println(line)
}

*/

func readJSON(data []byte) DataFrame {
	DF := DataFrame{}
	return DF
}
