#!/usr/bin/env bash
set -euo pipefail

# --- config ---
DB_FILE="${DB_FILE:-data/poll.sqlite}"
POLL_ID="${POLL_ID:-bday-2025}"
POLL_TITLE="${POLL_TITLE:-Birthday Poll}"
POLL_DESC="${POLL_DESC:-Pick one option. You can change your mind.}"

mkdir -p "$(dirname "$DB_FILE")"

sqlite3 "$DB_FILE" <<SQL
PRAGMA journal_mode=WAL;
PRAGMA foreign_keys=ON;

BEGIN;

-- Poll metadata
CREATE TABLE IF NOT EXISTS poll (
  id           TEXT PRIMARY KEY,
  title        TEXT NOT NULL,
  description  TEXT,
  created_at   DATETIME NOT NULL DEFAULT (datetime('now'))
);

-- Votes (1 per voter per poll)
CREATE TABLE IF NOT EXISTS votes (
  poll_id   TEXT NOT NULL,
  voter_id  TEXT NOT NULL,
  choice_id TEXT NOT NULL,
  voted_at  DATETIME NOT NULL DEFAULT (datetime('now')),
  PRIMARY KEY (poll_id, voter_id),
  FOREIGN KEY (poll_id) REFERENCES poll(id) ON DELETE CASCADE
);

-- Helpful index for results aggregation
CREATE INDEX IF NOT EXISTS votes_poll_choice_idx ON votes (poll_id, choice_id);

-- Insert or update poll metadata
INSERT INTO poll (id, title, description) VALUES (
  '$POLL_ID', '$POLL_TITLE', '$POLL_DESC'
)
ON CONFLICT(id) DO UPDATE SET
  title = excluded.title,
  description = excluded.description;

COMMIT;
SQL

echo "Initialized $DB_FILE with poll=$POLL_ID"
