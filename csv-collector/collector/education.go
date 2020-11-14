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

func ProcessAllEducations() error {
	log.Debug("[Collector] ProcessAllEducations")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	group := bson.D{
		primitive.E{Key: "tseId", Value: "$cd_grau_instrucao"},
		primitive.E{Key: "name", Value: "$ds_grau_instrucao"},
	}

	results, err := dbInstance.GetUnique(
		&politicos.Candidatures{},
		&politicos.Education{},
		group,
	)
	if err != nil {
		return err
	}

	// FIXME
	educations := []interface{}{}
	for _, p := range results {
		e := p.(*politicos.Education)
		e.Slug = slug.Make(e.Name)
		educations = append(educations, e)
	}

	_, err = dbInstance.InsertMany(&politicos.Education{}, educations)
	if err != nil {
		return err
	}

	return nil
}
