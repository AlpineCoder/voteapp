//go:build test_all || func_test
// +build test_all func_test

package repository

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var polls *Polls

func init() {
	polls, _ = NewPolls("../../data/poll.sqlite")
}

func TestPolls_AddVote(t *testing.T) {
	type args struct {
		poll_id   string
		voter_id  string
		choide_id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test_AddVote",
			args:    args{poll_id: "1", voter_id: "voter1", choide_id: "choice1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := polls
			if err := p.InsertVote(tt.args.poll_id, tt.args.voter_id, tt.args.choide_id); (err != nil) != tt.wantErr {
				t.Errorf("Polls.AddVote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
