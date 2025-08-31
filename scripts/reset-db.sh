#!/usr/bin/env bash
set -euo pipefail

# --- config ---
DB_FILE="${DB_FILE:-data/poll.sqlite}"
POLL_ID="${POLL_ID:-bday-2025}"
POLL_TITLE="${POLL_TITLE:-Birthday Poll}"
POLL_DESC="${POLL_DESC:-Pick one option. You can change your mind.}"

# remove old db and aux files
rm -f "$DB_FILE" "$DB_FILE-shm" "$DB_FILE-wal"

mkdir -p "$(dirname "$DB_FILE")"

sqlite3 "$DB_FILE" <<SQL
PRAGMA journal_mode=WAL;
PRAGMA foreign_keys=ON;

BEGIN;

-- Poll metadata
CREATE TABLE poll (
  id           TEXT PRIMARY KEY,
  title        TEXT NOT NULL,
  description  TEXT,
  created_at   DATETIME NOT NULL DEFAULT (datetime('now'))
);

-- Votes (1 per voter per poll)
CREATE TABLE votes (
  poll_id   TEXT NOT NULL,
  voter_id  TEXT NOT NULL,
  choice_id TEXT NOT NULL,
  voted_at  DATETIME NOT NULL DEFAULT (datetime('now')),
  PRIMARY KEY (poll_id, voter_id),
  FOREIGN KEY (poll_id) REFERENCES poll(id) ON DELETE CASCADE
);

CREATE INDEX votes_poll_choice_idx ON votes (poll_id, choice_id);

-- Insert poll metadata
INSERT INTO poll (id, title, description) VALUES (
  '$POLL_ID', '$POLL_TITLE', '$POLL_DESC'
);

COMMIT;
SQL

echo "Reset $DB_FILE with fresh poll=$POLL_ID"
