package activity

import (
	"github.com/joao-silva-1007/repo-activity/csv"
	"time"
)

// RepoNameToActivityMap correlates a repository name with its activity value
type RepoNameToActivityMap = map[string]float64

const (
	pointsForMostRecent = 1000.0
	pointsPerAdd        = 3.0
	pointsPerSub        = 2.0
	pointsPerFile       = 1.0
)

// ParseActivityPerRepositories creates a map of repositoryName -> activityValue based on csv entries
func ParseActivityPerRepositories(entries []csv.Entry) RepoNameToActivityMap {
	//since the provided data isn't recent, i'm searching for the last commit. If it was recent, the current time and date would suffice
	mostRecentTimestamp := findMostRecentTimestamp(entries)
	activityMap := make(RepoNameToActivityMap)
	for _, entry := range entries {
		calculatedActivity := calculateActivityForCommit(entry, mostRecentTimestamp)
		// creates new entry if it doesn't exist, updates current one if it does
		currentActivity := 0.0
		if activity, ok := activityMap[entry.Repository]; ok {
			currentActivity = activity
		}
		activityMap[entry.Repository] = currentActivity + calculatedActivity
	}

	return activityMap
}

// findMostRecentTimestamp searches for the greatest timestmap value inside a CSVEntry slice
func findMostRecentTimestamp(entries []csv.Entry) time.Time {
	mostRecent := entries[0].Timestamp
	arrFromFirst := entries[1:]
	for _, entry := range arrFromFirst {
		if entry.Timestamp.After(mostRecent) {
			mostRecent = entry.Timestamp
		}
	}
	return mostRecent
}

// calculateActivityForCommit calculates an activity score for a csv entry (commit) based on previously defined
// arbitrary values
func calculateActivityForCommit(entry csv.Entry, mostRecent time.Time) float64 {
	activity := 0.0

	// creates an inversely proportional relationship between the points given and the difference
	// between the entry timestamp and the most recent commit recorded
	activity += pointsForMostRecent / (mostRecent.Sub(entry.Timestamp).Hours() + 1)
	activity += float64(pointsPerAdd * entry.Additions)
	activity += float64(pointsPerSub * entry.Deletions)
	activity += float64(pointsPerFile * entry.Files)

	return activity
}
