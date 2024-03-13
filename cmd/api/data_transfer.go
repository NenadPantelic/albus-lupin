package main

// GradingRequest represents an input to grading; contains user's submission output
type GradingRequest struct {
	SubmissionID string `json:"submission_id"`
	AssignmentID string `json:"assignment_id"`
	Output       string `json:"output"`
	UserID       string `json:"user_id"`
}

// GradingResponse represents an output of the grading; contains the results
type GradingResponse struct {
	SubmissionID   string  `json:"submission_id"`
	AssignmentID   string  `json:"assignment_id"`
	Output         string  `json:"output"`
	ExpectedOutput string  `json:"expected_output"`
	UserID         string  `json:"user_id"`
	Score          float32 `json:"score"`
	Total          float32 `json:"total"`
}
