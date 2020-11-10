// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package db

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/olhoneles/politicos-go/politicos"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoSession struct {
	collection Collection
}

type DB interface {
	GetAll(q politicos.Queryable, p int) ([]politicos.Queryable, error)
	GetUnique(f politicos.Queryable, q politicos.Queryable, opts bson.D) ([]politicos.Queryable, error)
	InsertMany(q politicos.Queryable, d []interface{}) (*mongo.InsertManyResult, error)
	Ping() error
}

func newDBContext() (context.Context, context.CancelFunc) {
	timeout := viper.GetDuration("db.operation.timeout")
	return context.WithTimeout(context.Background(), timeout*time.Second)
}

func NewMongoSession() (DB, error) {
	log.Debug("[DB] New mongo session")
	dbURI := viper.GetString("mongodb.endpoint")
	dbName := viper.GetString("mongodb.name")
	ctx, _ := newDBContext()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Errorf("[DB] Error on create mongo session: %s", err)
	}
	mongo := &mongoSession{
		collection: &mongoCollection{
			session: client,
			dbName:  dbName,
		},
	}
	return mongo, err
}

func (m *mongoSession) Ping() error {
	log.Debug("[DB] Ping")
	return m.collection.Ping()
}

func (m *mongoSession) InsertMany(q politicos.Queryable, d []interface{}) (*mongo.InsertManyResult, error) {
	log.Debug("[DB] InsertMany")
	return m.collection.InsertMany(q.GetCollectionName(), d)
}

func (m *mongoSession) GetUnique(f politicos.Queryable, q politicos.Queryable, opts bson.D) ([]politicos.Queryable, error) {
	log.Debug("[DB] GetUnique")

	group := bson.D{
		primitive.E{
			Key:   "$group",
			Value: bson.D{primitive.E{Key: "_id", Value: opts}}},
	}
	pipeline := mongo.Pipeline{group}
	results, err := m.collection.Aggregate(f.GetCollectionName(), pipeline)
	if err != nil {
		return nil, err
	}

	operationList := []politicos.Queryable{}
	for _, result := range results {
		bsonBytes, err := bson.Marshal(result["_id"])
		if err != nil {
			return nil, err
		}
		p := q.Cast()
		if err := bson.Unmarshal(bsonBytes, p); err != nil {
			return nil, err
		}
		operationList = append(operationList, p)
	}

	return operationList, nil
}

func (m *mongoSession) GetAll(d politicos.Queryable, page int) ([]politicos.Queryable, error) {
	log.Debug("[DB] GetAll")

	query := bson.M{}
	perPage := viper.GetInt("db.operation.per-page")
	opts := options.Find()
	opts.SetSkip(int64((page - 1) * perPage))
	opts.SetLimit(int64(perPage))
	results, err := m.collection.FindAll(d.GetCollectionName(), query, opts)
	if err != nil {
		return nil, err
	}

	operationsList := []politicos.Queryable{}
	for _, result := range results {
		bsonBytes, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}
		cast := d.Cast()
		if err := bson.Unmarshal(bsonBytes, cast); err != nil {
			return nil, err
		}
		operationsList = append(operationsList, cast)
	}

	return operationsList, nil
}
