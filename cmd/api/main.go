package main

import (
	"fmt"
	"log"
	"net/http"
	"rs/ac/kg/fin/albus/lupin/internal"
	"time"
)

const port = 9881

type application struct {
	grader internal.Grader
}

func main() {
	assigmentScoringParamsHandler := internal.NewCachedAssignmentScoringParametersHandler(5*time.Minute, 10*time.Minute)

	grader := internal.Grader{
		AssigmentScoringParamsHandler: &assigmentScoringParamsHandler,
	}

	app := application{
		grader: grader,
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes()))
}
