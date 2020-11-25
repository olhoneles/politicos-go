// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package collector

import (
	"github.com/gosimple/slug"
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	log "github.com/sirupsen/logrus"
)

func collectorBase(q politicos.Queryable, opts db.UniqueOptions) error {
	log.Debug("[Collector] CollectorBase")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	results, err := dbInstance.GetUnique(&politicos.Candidatures{}, q, opts)
	if err != nil {
		return err
	}

	// FIXME
	items := []interface{}{}
	for _, p := range results {
		p.SetSlug(slug.Make(p.GetID()))
		items = append(items, p)
	}

	_, err = dbInstance.InsertMany(q, items)
	if err != nil {
		return err
	}

	return nil
}
