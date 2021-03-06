// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package collector

import (
	"github.com/gosimple/slug"
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProcessAllCandidaciesStatus() error {
	log.Debug("[Collector] ProcessAllCandidaciesStatus")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	opts := db.UniqueOptions{
		IDs: bson.D{
			primitive.E{Key: "tseId", Value: "$cd_sit_tot_turno"},
			primitive.E{Key: "name", Value: "$ds_sit_tot_turno"},
		},
	}

	results, err := dbInstance.GetUnique(
		&politicos.Candidatures{},
		&politicos.CandidacyStatus{},
		opts,
	)
	if err != nil {
		return err
	}

	// FIXME
	candidaciesStatus := []interface{}{}
	for _, p := range results {
		c := p.(*politicos.CandidacyStatus)
		c.Slug = slug.Make(c.Name)
		candidaciesStatus = append(candidaciesStatus, c)
	}

	_, err = dbInstance.InsertMany(&politicos.CandidacyStatus{}, candidaciesStatus)
	if err != nil {
		return err
	}

	return nil
}
