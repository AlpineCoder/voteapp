package model

const (
	PollID = "bday-2025"
)

type Poll struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description,omitempty"`
	Options     map[string]string `json:"options"` // option name to vote count
}

type PollOptions map[string]string

type Vote struct {
	ChoiceID string `json:"choiceId"`
}

type Results struct {
	PollID     string       `json:"poll_id"`
	TotalVotes int          `json:"total_votes"`
	Options    OptionCounts `json:"options"`
}

type OptionCounts struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Votes int    `json:"votes"`
}

// Predefined poll options
var DefinedPollOptions = PollOptions{
	"optionA": "Punk",
	"optionB": "Happy Hardcore",
	"optionC": "Schlager",
}
