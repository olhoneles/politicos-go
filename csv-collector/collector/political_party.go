// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package collector

import (
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProcessAllPoliticalParties() error {
	log.Debug("[Collector] ProcessAllPoliticalParties")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	group := bson.D{
		primitive.E{Key: "siglum", Value: "$sg_partido"},
		primitive.E{Key: "name", Value: "$nm_partido"},
		primitive.E{Key: "tseNumber", Value: "$nr_partido"},
	}

	results, err := dbInstance.GetUnique(
		&politicos.Candidatures{},
		&politicos.PoliticalParty{},
		group,
	)
	if err != nil {
		return err
	}

	// FIXME
	politicalParties := []interface{}{}
	for _, p := range results {
		politicalParties = append(politicalParties, p)
	}

	_, err = dbInstance.InsertMany(&politicos.PoliticalParty{}, politicalParties)
	if err != nil {
		return err
	}

	return nil
}
