package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbRepository struct {
	host   string
	port   int64
	client *mongo.Client
}

func newMongoDbRepository(host string, port int64) MongoDbRepository {
	return MongoDbRepository{host: host, port: port}
}

func (m *MongoDbRepository) Connect() error {
	connectionError := errors.New("connection fail")
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", m.host, m.port))
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		return connectionError
	}
	m.client = client

	if err = m.client.Ping(context.TODO(), nil); err != nil {
		return connectionError
	}

	m.client = client
	return nil
}

func (m MongoDbRepository) Close() error {
	closeError := errors.New("close connection fail")

	if err := m.client.Disconnect(context.TODO()); err != nil {
		return closeError
	}
	return nil
}

func (m MongoDbRepository) Insert(database, collection string, documents []interface{}) error {
	type coll struct {
		userAlias string `bson:"user_alias"`
		userName  string `bson:"user_name"`
		userId    string `bson:"user_id"`
	}
	c := coll{
		userAlias: "rcrespo",
		userName:  "rodrigo crespo",
		userId:    "41589231",
	}

	err := m.Connect()
	if err != nil {
		return err
	}

	defer func() {
		err = m.Close()
	}()

	mongoCollection := m.client.Database(database).Collection(collection)

	if len(documents) > 1 {
		_, err := mongoCollection.InsertMany(context.TODO(), documents)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = mongoCollection.InsertOne(context.TODO(), c)
	if err != nil {
		return err
	}
	return nil
}
