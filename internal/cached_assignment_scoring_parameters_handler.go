package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

const scoringParamsFilesRoot = "scoring"

type CachedAssignmentScoringParametersHandler struct {
	ASPCache cache.Cache
}

func NewCachedAssignmentScoringParametersHandler(defaultExpiration time.Duration, cleanupInterval time.Duration) CachedAssignmentScoringParametersHandler {
	cache := cache.New(defaultExpiration, cleanupInterval)

	return CachedAssignmentScoringParametersHandler{
		ASPCache: *cache,
	}
}

func (handler *CachedAssignmentScoringParametersHandler) Get(assignmentID string) (AssignmentScoringParameters, error) {
	log.Printf("Fetching scoring parameters of an assignment %s", assignmentID)

	if element, found := handler.ASPCache.Get(assignmentID); found {
		return element.(AssignmentScoringParameters), nil
	} else {
		filepath := fmt.Sprintf("%s/%s.json", scoringParamsFilesRoot, assignmentID)
		scoringParams, err := GetFileContent(filepath)

		if err != nil {
			return AssignmentScoringParameters{}, nil
		}

		handler.ASPCache.Set(assignmentID, scoringParams, 30*time.Minute)
		return scoringParams, nil
	}

}
