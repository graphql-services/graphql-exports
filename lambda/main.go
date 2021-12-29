package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/graphql-services/graphql-exports/gen"
	"github.com/graphql-services/graphql-exports/src"
)

func main() {
	db := gen.NewDBFromEnvVars()

	eventController, err := gen.NewEventController()
	if err != nil {
		panic(err)
	}

	handler := gen.GetHTTPServeMux(src.New(db, &eventController), db, src.GetMigrations(db))
	algnhsa.ListenAndServe(handler, nil)
}
