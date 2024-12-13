package activity

import (
	"github.com/joao-silva-1007/repo-activity/csv"
	"reflect"
	"testing"
	"time"
)

func TestParseActivityPerRepositories(t *testing.T) {
	now := time.Now()
	type args struct {
		entries []csv.Entry
	}
	tests := []struct {
		name string
		args args
		want RepoNameToActivityMap
	}{
		{name: "timestamp same as most recent",
			args: args{
				entries: []csv.Entry{
					{
						Timestamp:  now,
						Username:   "user10",
						Repository: "repo1",
						Files:      5,
						Additions:  10,
						Deletions:  20,
					},
					{
						Timestamp:  now,
						Username:   "user13",
						Repository: "repo3",
						Files:      25,
						Additions:  12,
						Deletions:  0,
					},
					{
						Timestamp:  now,
						Username:   "user18",
						Repository: "repo2",
						Files:      1,
						Additions:  150,
						Deletions:  22,
					}, {
						Timestamp:  now,
						Username:   "user1",
						Repository: "repo3",
						Files:      8,
						Additions:  250,
						Deletions:  0,
					},
					{
						Timestamp:  now,
						Username:   "",
						Repository: "repo1",
						Files:      15,
						Additions:  250,
						Deletions:  200,
					},
				},
			},
			want: RepoNameToActivityMap{"repo1": 3240, "repo3": 2819, "repo2": 1495},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseActivityPerRepositories(tt.args.entries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseActivityPerRepositories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateActivityForCommit(t *testing.T) {
	now := time.Now()
	type args struct {
		entry      csv.Entry
		mostRecent time.Time
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "timestamp same as most recent",
			args: args{
				entry: csv.Entry{
					Timestamp:  now,
					Username:   "user1",
					Repository: "repo1",
					Files:      10,
					Additions:  25,
					Deletions:  20,
				},
				mostRecent: now,
			},
			want: 1125.0,
		},
		{name: "timestamp different than most recent",
			args: args{
				entry: csv.Entry{
					Timestamp:  now.Add(-9 * time.Hour),
					Username:   "user1",
					Repository: "repo1",
					Files:      10,
					Additions:  25,
					Deletions:  20,
				},
				mostRecent: now,
			},
			want: 225.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateActivityForCommit(tt.args.entry, tt.args.mostRecent); got != tt.want {
				t.Errorf("calculateActivityForCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMostRecentTimestamp(t *testing.T) {
	type args struct {
		entries []csv.Entry
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "one of the values is greater than the others",
			args: args{
				entries: []csv.Entry{
					{Timestamp: now.Add(-20 * time.Hour)},
					{Timestamp: now},
					{Timestamp: now.Add(20 * time.Hour)},
				},
			},
			want: now.Add(20 * time.Hour),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMostRecentTimestamp(tt.args.entries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMostRecentTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
