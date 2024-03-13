package internal

// GradingResult represents an output of the grading; contains the results
type GradingResult struct {
	AssignmentID   string
	Output         string
	ExpectedOutput string
	Score          float32
	Total          float32
}
