package folio

import "time"

type Node struct {
	// Meta
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Created      time.Time `json:"created"`
	LastModified time.Time `json:"last_modified"`
	// Content
	Content  string   `json:"content,omitempty"`
	Outbound []string `json:"outbound,omitempty"`
	Inbound  []string `json:"inbound,omitempty"`
	// Node-specific
	Status      string `json:"status,omitempty"`
	Confidence  string `json:"confidence,omitempty"`
	Origin      string `json:"origin,omitempty"`
	PublishDate string `json:"publish_date,omitempty"`
}

type Folio struct {
	// Meta
	ID            string    `json:"id"`
	SchemaVersion string    `json:"schema_version"`
	Name          string    `json:"name"`
	Created       time.Time `json:"created"`
	LastModified  time.Time `json:"last_modified"`
	// Content
	Sources   []Node `json:"sources,omitempty"`
	Cores     []Node `json:"cores,omitempty"`
	Questions []Node `json:"questions,omitempty"`
	Notes     []Node `json:"notes,omitempty"`
	Tasks     []Node `json:"tasks,omitempty"`
	Dailies   []Node `json:"dailies,omitempty"`
	Ghosts    []Node `json:"ghosts,omitempty"`
}
