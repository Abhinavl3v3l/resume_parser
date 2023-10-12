package types

// Unmarshal JSON string to this.
type CandidateInfo struct {
	Email           string   `json:"email"`
	Skills          []string `json:"Skills"`
	ExperienceLevel int      `json:"Experience Level"`
}
