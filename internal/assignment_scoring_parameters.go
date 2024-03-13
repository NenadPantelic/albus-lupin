package internal

type ScoringPolicy string

const (
	BINARY    ScoringPolicy = "BINARY"
	PARTIALLY ScoringPolicy = "PARTIALLY"
)

// AssignmentScoringParameters contains the data necessary for scoring some assigment
type AssignmentScoringParameters struct {
	AssignmentID  string        `json:"assignment_id"`
	Solution      string        `json:"solution"`
	MaxScore      float32       `json:"max_score"`
	ScoringPolicy ScoringPolicy `json:"scoring_policy"`
}
