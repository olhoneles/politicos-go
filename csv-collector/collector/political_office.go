// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package collector

import (
	"github.com/labstack/gommon/log"
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProcessAllPoliticalOffices() error {
	log.Debug("[Collector] ProcessAllPoliticalOffices")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	group := bson.D{
		primitive.E{Key: "tseId", Value: "$cd_cargo"},
		primitive.E{Key: "name", Value: "$ds_cargo"},
	}

	results, err := dbInstance.GetUnique(
		&politicos.Candidatures{},
		&politicos.PoliticalOffice{},
		group,
	)
	if err != nil {
		return err
	}

	// FIXME
	politicalOffices := []interface{}{}
	for _, p := range results {
		politicalOffices = append(politicalOffices, p)
	}

	_, err = dbInstance.InsertMany(&politicos.PoliticalOffice{}, politicalOffices)
	if err != nil {
		return err
	}

	return nil
}
