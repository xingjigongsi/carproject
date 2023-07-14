package mongodb

import "go.mongodb.org/mongo-driver/mongo"

const MONDBAPP = "app:monodb"

type MongodbInterface interface {
	MongodbClient() (*mongo.Client, error)
}
