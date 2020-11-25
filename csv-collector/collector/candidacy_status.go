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

func ProcessAllCandidaciesStatus() error {
	log.Debug("[Collector] ProcessAllCandidaciesStatus")

	opts := db.UniqueOptions{
		IDs: bson.D{
			primitive.E{Key: "tseId", Value: "$cd_sit_tot_turno"},
			primitive.E{Key: "name", Value: "$ds_sit_tot_turno"},
		},
	}

	if err := collectorBase(&politicos.CandidacyStatus{}, opts); err != nil {
		return err
	}

	return nil
}
