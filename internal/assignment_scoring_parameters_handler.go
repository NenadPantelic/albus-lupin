package internal

// treba da dohvati resenja i poene za assignment; jedna implementacija Cached.....; sredi naming
type AssignmentScoringParametersHandler interface {
	Get(assignmentID string) (AssignmentScoringParameters, error)
}
