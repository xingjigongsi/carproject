package mongodb

import (
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"github.com/xingjigongsi/carproject/framework/container"
	"time"
)

type MongoDbService struct {
}

func (mongodbService *MongoDbService) Register(appContainer container.InterfaceContainer) container.NewInstance {
	return MondbNewInstanse
}
func (mongodbService *MongoDbService) Params(appContainer container.InterfaceContainer) []interface{} {
	mustMake := appContainer.MustMake(parse.PASE_NAME).(*parse.ParseApply)
	mongodbUrl, _ := mustMake.GetString("database.mongodb.mongodbUrl")
	timeout, _ := mustMake.GetInt("database.mongodb.timeOut")
	timeoutstr := time.Duration(timeout) * time.Second
	return []interface{}{appContainer, mongodbUrl, timeoutstr}
}
func (mongodbService *MongoDbService) ApplyInit(appContainer container.InterfaceContainer) error {
	return nil
}
func (mongodbService *MongoDbService) Name() string {
	return MONDBAPP
}
func (mongodbService *MongoDbService) IsDefer() bool {
	return false
}
