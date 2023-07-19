package mongodb

import (
	"context"
	"github.com/xingjigongsi/carproject/framework/container"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongodbApply struct {
	Container  container.InterfaceContainer
	TimeOut    time.Duration
	MongodbUrl string
}

func MondbNewInstanse(params ...interface{}) (interface{}, error) {
	Container := params[0].(container.InterfaceContainer)
	mongdbUrl := params[1].(string)
	timeout := params[2].(time.Duration)
	return &MongodbApply{
		Container:  Container,
		TimeOut:    timeout,
		MongodbUrl: mongdbUrl,
	}, nil
}

func (mongodbApply *MongodbApply) MongodbClient() (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongodbApply.TimeOut)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongodbApply.MongodbUrl))
	return
}
