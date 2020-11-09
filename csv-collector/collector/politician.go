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

const ItemsPerInsert = 1000

func ProcessAllPoliticians() error {
	log.Debug("[Collector] ProcessAllPoliticians")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	log.Debugf("[Collector] Getting all unique CPF...")
	opts := db.UniqueOptions{
		IDs: bson.D{primitive.E{Key: "cpf", Value: "$nr_cpf_candidato"}},
	}
	// FIXME: paginate?
	cpfs, err := dbInstance.GetUnique(
		&politicos.Candidatures{},
		&politicos.Politician{},
		opts,
	)
	if err != nil {
		return err
	}

	log.Debugf("[Collector] Total of CPFs: %d", len(cpfs))

	politicians := []interface{}{}
	for _, candidate := range cpfs {
		c := candidate.(*politicos.Politician)
		log.Debugf("[Collector] Grouping %s data...", c.CPF)
		opts := db.UniqueOptions{
			Match: bson.D{primitive.E{Key: "nr_cpf_candidato", Value: c.CPF}},
			IDs:   bson.D{primitive.E{Key: "cpf", Value: "$nr_cpf_candidato"}},
			Data:  bson.D{primitive.E{Key: "$push", Value: "$$ROOT"}},
			Sort: bson.D{
				primitive.E{Key: "ano_eleicao", Value: -1},
				primitive.E{Key: "nr_turno", Value: -1},
			},
		}
		results, err := dbInstance.GetUnique(
			&politicos.Candidatures{},
			&politicos.Politician{},
			opts,
		)
		if err != nil {
			return err
		}

		// FIXME
		item := results[0].(*politicos.Politician)
		candidature := item.Data[0]
		item.CPF = candidature.NR_CPF_CANDIDATO
		item.CD_COR_RACA = candidature.CD_COR_RACA
		item.CD_ESTADO_CIVIL = candidature.CD_ESTADO_CIVIL
		item.CD_MUNICIPIO_NASCIMENTO = candidature.CD_MUNICIPIO_NASCIMENTO
		item.CD_NACIONALIDADE = candidature.CD_NACIONALIDADE
		item.DS_COR_RACA = candidature.DS_COR_RACA
		item.DS_ESTADO_CIVIL = candidature.DS_ESTADO_CIVIL
		item.DS_NACIONALIDADE = candidature.DS_NACIONALIDADE
		item.DT_NASCIMENTO = candidature.DT_NASCIMENTO
		item.NM_CANDIDATO = candidature.NM_CANDIDATO
		item.NM_EMAIL = candidature.NM_EMAIL
		item.NM_MUNICIPIO_NASCIMENTO = candidature.NM_MUNICIPIO_NASCIMENTO
		item.NM_SOCIAL_CANDIDATO = candidature.NM_SOCIAL_CANDIDATO
		item.NM_URNA_CANDIDATO = candidature.NM_URNA_CANDIDATO
		item.NR_TITULO_ELEITORAL_CANDIDATO = candidature.NR_TITULO_ELEITORAL_CANDIDATO
		item.SG_UF_NASCIMENTO = candidature.SG_UF_NASCIMENTO
		item.SQ_CANDIDATO = candidature.SQ_CANDIDATO

		politicians = append(politicians, item)

		if len(politicians) == ItemsPerInsert {
			log.Debugf("[Collector] Inserting %d...", ItemsPerInsert)
			_, err := dbInstance.InsertMany(&politicos.Politician{}, politicians)
			if err != nil {
				return err
			}
			politicians = []interface{}{}
		}
	}

	if len(politicians) > 0 {
		log.Debugf("[Collector] Inserting %d...", len(politicians))
		_, err = dbInstance.InsertMany(&politicos.Politician{}, politicians)
		if err != nil {
			return err
		}
	}

	return nil
}
