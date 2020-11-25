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

	opts := db.UniqueOptions{
		IDs: bson.D{
			primitive.E{Key: "siglum", Value: "$sg_partido"},
			primitive.E{Key: "name", Value: "$nm_partido"},
			primitive.E{Key: "tseNumber", Value: "$nr_partido"},
		},
	}

	if err := collectorBase(&politicos.PoliticalParty{}, opts); err != nil {
		return err
	}

	return nil
}
