package csv

import (
	"errors"
	"strconv"
	"time"
)

// Entry contains all fields present in a certain csv format
type Entry struct {
	Timestamp  time.Time
	Username   string
	Repository string
	Files      int
	Additions  int
	Deletions  int
}

// NewEntry creates an Entry based on the raw text fields read from a csv file
func NewEntry(csvLine []string) (*Entry, error) {
	if len(csvLine) != 6 {
		return nil, errors.New("line must have 6 elements")
	}
	unixTime, err := strconv.ParseInt(csvLine[0], 0, 0)
	if err != nil {
		return nil, err
	}
	timestampTime := time.Unix(unixTime, 0)
	filesNumber, err := strconv.ParseInt(csvLine[3], 0, 0)
	if err != nil {
		return nil, err
	}
	additionsNumber, err := strconv.ParseInt(csvLine[4], 0, 0)
	if err != nil {
		return nil, err
	}
	deletionsNumber, err := strconv.ParseInt(csvLine[5], 0, 0)
	if err != nil {
		return nil, err
	}
	return &Entry{
		Timestamp:  timestampTime,
		Username:   csvLine[1],
		Repository: csvLine[2],
		Files:      int(filesNumber),
		Additions:  int(additionsNumber),
		Deletions:  int(deletionsNumber),
	}, nil
}
