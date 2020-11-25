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

func ProcessAllEducations() error {
	log.Debug("[Collector] ProcessAllEducations")

	opts := db.UniqueOptions{
		IDs: bson.D{
			primitive.E{Key: "tseId", Value: "$cd_grau_instrucao"},
			primitive.E{Key: "name", Value: "$ds_grau_instrucao"},
		},
	}

	if err := collectorBase(&politicos.Education{}, opts); err != nil {
		return err
	}

	return nil
}
