package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var y *mongo.Client

func InitDataStore(user, password, host string, port int) error {
	var err error
	fmt.Println(user, password, host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	urlString := fmt.Sprintf("mongodb://%s:%s@%s:%d/", user, password, host, port)
	fmt.Println(urlString)
	y, err = mongo.Connect(ctx, options.Client().ApplyURI(urlString))
	if err != nil {
		return err
	}

	return err
}

func GetDataStoreClient() *mongo.Client {
	return y
}

func StopDataStoreClient() error {
	return y.Disconnect(context.TODO())
}
