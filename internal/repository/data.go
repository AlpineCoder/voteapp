package repository

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Polls struct {
	mu *sync.Mutex
	db *sql.DB
}

func NewPolls(dataSourceName string) (*Polls, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Polls{
		db: db,
		mu: &sync.Mutex{},
	}, nil
}

func (p *Polls) Close() error {
	return p.db.Close()
}

// UpsertVote inserts a new vote or updates an existing one.
func (p *Polls) UpsertVote(poll_id, voter_id, choide_id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, err := p.db.Exec(`INSERT INTO votes (poll_id, voter_id, choice_id)
						 VALUES (?, ?, ?)
						 ON CONFLICT(poll_id, voter_id)
						 DO UPDATE SET choice_id=excluded.choice_id,
						               voted_at=datetime('now')`,
		poll_id, voter_id, choide_id)
	return err
}

// GetVote retrieves the choice_id for a given poll_id and voter_id.
func (p *Polls) GetVote(poll_id, voter_id string) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var choice_id string
	err := p.db.QueryRow(`SELECT choice_id FROM votes WHERE poll_id = ? AND voter_id = ?`,
		poll_id, voter_id).Scan(&choice_id)
	if err != nil {
		return "", err
	}
	return choice_id, nil
}
