package github

import (
	"encoding/json"
	"time"
)

type GitHubTimestamp struct {
	*time.Time
}

func (t *GitHubTimestamp) MarshalJSON() ([]byte, error) {
	return t.Time.MarshalJSON()
}

func (t *GitHubTimestamp) UnmarshalJSON(data []byte) error {
	// Try deserializing as timestamp then time
	var x int64 = 0
	newt := time.Time{}
	if err := json.Unmarshal(data, &x); err != nil {
		if err := json.Unmarshal(data, &newt); err != nil {
			return err
		}
	} else {
		newt = time.Unix(x, 0)
	}

	// From timestamp
	t.Time = &newt

	return nil
}
