package internal

import (
	"log"
	"strings"
)

type Grader struct {
	AssigmentScoringParamsHandler AssignmentScoringParametersHandler
}

func (grader *Grader) Grade(assigmentID string, valueToEvaluate string) (GradingResult, error) {
	log.Printf("Grading assigmentID = %s, valueToEvaluate = %s", assigmentID, valueToEvaluate)

	assignmentScoringParameters, err := grader.AssigmentScoringParamsHandler.Get(assigmentID)
	if err != nil {
		return GradingResult{}, err
	}

	scorePercentage := doCompare(assignmentScoringParameters.Solution, valueToEvaluate, assignmentScoringParameters.ScoringPolicy)

	return GradingResult{
		AssignmentID:   assigmentID,
		Output:         valueToEvaluate,
		ExpectedOutput: assignmentScoringParameters.Solution,
		Score:          scorePercentage * assignmentScoringParameters.MaxScore,
		Total:          assignmentScoringParameters.MaxScore,
	}, nil
}

func doCompare(expectedValue string, actualValue string, scoringPolicy ScoringPolicy) float32 {
	if scoringPolicy == BINARY {
		if expectedValue == actualValue {
			return 1.0
		}

		return 0.0
	}

	expectedValueLines := strings.Split(expectedValue, "\n")
	actualValueLines := strings.Split(actualValue, "\n")

	noOflinesToProccess := len(expectedValueLines)
	if len(expectedValueLines) < len(actualValueLines) {
		noOflinesToProccess = len(actualValueLines)
	}

	score := 0
	for i := 0; i < noOflinesToProccess; i++ {
		valueA, presentA := getElementAt(i, expectedValueLines)
		valueB, presentB := getElementAt(i, actualValueLines)

		if presentA && presentB && valueA == valueB {
			score += 1
		}
	}

	return float32(score) / float32(len(expectedValueLines))
}

func getElementAt(index int, elements []string) (string, bool) {
	if index >= len(elements) {
		return "", false
	}

	return elements[index], true
}
