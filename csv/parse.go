package csv

import (
	"encoding/csv"
	"log"
	"os"
)

// ParseFile reads a csv file's content and iterates over it, parsing each of the entries into an Entry
func ParseFile(fileName string) []Entry {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("can't open file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Panic("no entries found in file", err)
	}
	linesWithoutHeader := lines[1:]
	var arr []Entry

	for _, line := range linesWithoutHeader {
		entry, err := NewEntry(line)
		if err == nil {
			arr = append(arr, *entry)
		} else {
			log.Print("found error creating line", err)
		}
	}
	return arr
}
