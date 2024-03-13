package main

import (
	"log"
	"net/http"
)

func (app *application) Status(w http.ResponseWriter, r *http.Request) {

	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Lupin grader service up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) Grade(w http.ResponseWriter, r *http.Request) {
	var gradingRequest GradingRequest

	err := app.readJSON(w, r, &gradingRequest)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("Grading request %v", gradingRequest)

	gradingResult, err := app.grader.Grade(gradingRequest.AssignmentID, gradingRequest.Output)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	gradingResponse := GradingResponse{
		SubmissionID:   gradingRequest.SubmissionID,
		AssignmentID:   gradingResult.AssignmentID,
		Output:         gradingResult.Output,
		ExpectedOutput: gradingResult.ExpectedOutput,
		UserID:         gradingRequest.UserID,
		Score:          gradingResult.Score,
		Total:          gradingResult.Total,
	}

	app.writeJSON(w, http.StatusCreated, gradingResponse)
}
