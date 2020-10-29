package db

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoCollection struct {
	dbName  string
	session *mongo.Client
}

type Collection interface {
	Aggregate(c string, pipeline interface{}) ([]bson.M, error)
	Distinct(c string, q string, f interface{}) ([]interface{}, error)
	FindAll(c string, q bson.M, o ...*options.FindOptions) ([]bson.M, error)
	InsertMany(c string, d []interface{}) (*mongo.InsertManyResult, error)
	Ping() error
}

func newCollectionContext() (context.Context, context.CancelFunc) {
	timeout := viper.GetDuration("collection.operation.timeout")
	return context.WithTimeout(context.Background(), timeout*time.Second)
}

func (m *mongoCollection) Aggregate(c string, pipeline interface{}) ([]bson.M, error) {
	log.Debug("[Collection] Aggregate...")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if cursor == nil {
		return nil, err
	}
	var results []bson.M
	ctx, _ = newCollectionContext()
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *mongoCollection) Distinct(c string, q string, f interface{}) ([]interface{}, error) {
	log.Debug("[Collection] Distinct")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	return collection.Distinct(ctx, q, f)
}

func (m *mongoCollection) FindAll(c string, q bson.M, o ...*options.FindOptions) ([]bson.M, error) {
	log.Debug("[Collection] FindAll")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	cur, err := collection.Find(ctx, q, o...)
	if err != nil {
		log.Errorf("[Collection] Find: %v", err)
		return nil, err
	}
	var results []bson.M
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			log.Errorf("[Collection] Decode: %v", err)
			return nil, err
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Errorf("[Collection] Cursor: %v", err)
		return nil, err
	}
	return results, nil
}

func (m *mongoCollection) InsertMany(c string, d []interface{}) (*mongo.InsertManyResult, error) {
	log.Debug("[Collection] InsertMany...")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	return collection.InsertMany(ctx, d)
}

func (m *mongoCollection) Ping() error {
	log.Debug("[Collection] Ping")
	ctx, _ := newCollectionContext()
	return m.session.Ping(ctx, readpref.Primary())
}
