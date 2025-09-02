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
func (p *Polls) UpsertVote(pollId, voterId, choiceId string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, err := p.db.Exec(`INSERT INTO votes (poll_id, voter_id, choice_id)
						 VALUES (?, ?, ?)
						 ON CONFLICT(poll_id, voter_id)
						 DO UPDATE SET choice_id=excluded.choice_id,
						               voted_at=datetime('now')`,
		pollId, voterId, choiceId)
	return err
}

// GetVote retrieves the choice_id for a given poll_id and voter_id.
func (p *Polls) GetVote(pollId, voterId string) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var choiceId string
	err := p.db.QueryRow(`SELECT choice_id FROM votes WHERE poll_id = ? AND voter_id = ?`,
		pollId, voterId).Scan(&choiceId)
	if err != nil {
		return "", err
	}
	return choiceId, nil
}

func (p *Polls) GetResults(pollId string) (map[string]int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	rows, err := p.db.Query(`SELECT choice_id, COUNT(*) as vote_count
							 FROM votes
							 WHERE poll_id = ?
							 GROUP BY choice_id`, pollId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(map[string]int)
	for rows.Next() {
		var choiceId string
		var count int
		if err := rows.Scan(&choiceId, &count); err != nil {
			return nil, err
		}
		results[choiceId] = count
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
