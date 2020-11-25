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

func ProcessAllPoliticalOffices() error {
	log.Debug("[Collector] ProcessAllPoliticalOffices")

	opts := db.UniqueOptions{
		IDs: bson.D{
			primitive.E{Key: "tseId", Value: "$cd_cargo"},
			primitive.E{Key: "name", Value: "$ds_cargo"},
		},
	}

	if err := collectorBase(&politicos.PoliticalOffice{}, opts); err != nil {
		return err
	}

	return nil
}
