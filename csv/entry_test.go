package csv

import (
	"reflect"
	"testing"
	"time"
)

func TestNewEntry(t *testing.T) {
	type args struct {
		csvLine []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Entry
		wantErr bool
	}{
		{
			name: "simple case",
			args: args{csvLine: []string{"1734107460", "user2", "repo3", "1", "3", "1"}},
			want: &Entry{
				Timestamp:  time.Unix(1734107460, 0),
				Username:   "user2",
				Repository: "repo3",
				Files:      1,
				Additions:  3,
				Deletions:  1,
			},
			wantErr: false,
		},
		{
			name:    "missing fields",
			args:    args{csvLine: []string{"1734107460", "repo3", "1", "3", "1"}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "fields not numbers",
			args:    args{csvLine: []string{"1734107460", "user1", "repo3", "notANumber", "3", "1"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEntry(tt.args.csvLine)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntry() got = %v, want %v", got, tt.want)
			}
		})
	}
}
