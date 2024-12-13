package main

import (
	"fmt"
	"github.com/joao-silva-1007/repo-activity/activity"
	"github.com/joao-silva-1007/repo-activity/csv"
	"log"
	"sort"
)

type KV struct {
	key   string
	value float64
}

func main() {
	csvEntries := csv.ParseFile("commits.csv")
	if len(csvEntries) == 0 {
		log.Panic("no entries read")
	}

	activityMap := activity.ParseActivityPerRepositories(csvEntries)

	// converts the map into a slice where each position has a key and value from the activityMap.
	// This way it is possible to sort by it's activity and still know what repository it is associated with
	var kvSlice []KV
	for key, value := range activityMap {
		kvSlice = append(kvSlice, KV{key: key, value: value})
	}

	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].value > kvSlice[j].value
	})

	// iterates through the 10 repositories with the most activity
	top10Entries := kvSlice[:10]
	for _, entry := range top10Entries {
		fmt.Printf("Repository \"%s\" had %d activity points\n", entry.key, int(entry.value))
	}
}
