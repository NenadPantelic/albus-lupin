package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetFileContent(filepath string) (AssignmentScoringParameters, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return AssignmentScoringParameters{}, err
	}

	var scoringParams AssignmentScoringParameters
	err = json.Unmarshal(content, &scoringParams)

	log.Printf("DBG err %v", err)
	if err != nil {
		return AssignmentScoringParameters{}, err
	}

	return scoringParams, nil
}
