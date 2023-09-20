package repositoryimpl

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"github.com/vandensudarsono/bus-system/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoImpl struct {
	client *mongo.Client
}

func NewMongoImpl(client *mongo.Client) *MongoImpl {
	return &MongoImpl{client: client}
}

func (m *MongoImpl) GetBuslines(ctx context.Context, offset, limit int64) []models.BusLine {
	var (
		result []models.BusLine
		err    error
		cur    *mongo.Cursor
	)

	cur, err = m.client.Database(viper.GetString("mongoDB.database")).
		Collection("available-bus").Find(
		ctx, bson.M{}, options.Find().SetLimit(limit).SetSkip(offset).SetProjection(bson.M{"_id": 0}))

	fmt.Println("error at get available bus lines: ", err)
	if err == nil {
		err = cur.All(ctx, &result)
		fmt.Println("error cursor at get bus lines: ", err)
		defer cur.Close(ctx)
	}

	return result
}

func (m *MongoImpl) GetBuslineById(ctx context.Context, busLineID string) models.BusLine {
	var (
		result models.BusLine
		err    error
	)

	err = m.client.Database(viper.GetString("mongoDB.database")).
		Collection("available-bus").FindOne(ctx, bson.M{"id": busLineID},
		options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&result)

	if err != nil {
		fmt.Println("error get bus line by id: ", err)
	}

	return result
}

func (m *MongoImpl) GetBuslineByBusStopId(ctx context.Context, busStopID string) ([]models.BusLine, error) {
	var (
		result []models.BusLine
		err    error
		cur    *mongo.Cursor
	)

	cur, err = m.client.Database(viper.GetString("mongoDB.database")).
		Collection("available-bus").Find(ctx, bson.M{"busStops.id": busStopID},
		options.Find().SetProjection(bson.M{"_id": 0}))

	fmt.Println("error at get available bus lines by bus stop id: ", err)
	if err == nil {
		err = cur.All(ctx, &result)
		fmt.Println("error cursor at get bus lines by bus stop id: ", err)
		defer cur.Close(ctx)
	}

	return result, err
}

func (m *MongoImpl) InsertBusLine(ctx context.Context, input []models.BusLine) bool {
	var (
		result    bool = true
		inputMany []interface{}
	)

	for _, data := range input {
		inputMany = append(inputMany, data)
	}

	_, err := m.client.Database(viper.GetString("mongoDB.database")).
		Collection("available-bus").InsertMany(ctx, inputMany)
	if err != nil {
		fmt.Println("error at insert many: ", err)
		result = false
	}

	return result
}

func (m *MongoImpl) InsertBusStops(ctx context.Context, input []models.BusStop) bool {
	var (
		result    bool = true
		inputMany []interface{}
	)

	for _, data := range input {
		inputMany = append(inputMany, data)
	}

	_, err := m.client.Database(viper.GetString("mongoDB.database")).
		Collection("bus-stops").InsertMany(ctx, inputMany)

	if err != nil {
		fmt.Println("error at insert bus stops: ", err)
	}

	return result

}

func (m *MongoImpl) GetBusStops(ctx context.Context, limit, offset int64) []models.BusStop {
	var (
		result []models.BusStop
	)

	cur, err := m.client.Database(viper.GetString("mongoDB.database")).
		Collection("bus-stops").Find(ctx, bson.D{}, options.Find().SetAllowDiskUse(true).
		SetLimit(limit).SetSkip(offset).SetProjection(bson.M{"busStops": 1}))

	if err != nil {
		fmt.Println("error at get bus stops: ", err)
		return result
	}

	err = cur.All(ctx, &result)
	fmt.Println("error cursor at get bus stops: ", err)
	defer cur.Close(ctx)

	return result
}

func (m *MongoImpl) GetBuslineByBusStopName(ctx context.Context, name string, offset, limit int64) []models.BusLine {
	var result []models.BusLine

	if name == "" {
		return result
	}

	cur, err := m.client.Database(viper.GetString("mongoDB.database")).
		Collection("available-bus").Find(ctx, bson.M{"busStops.name": bson.M{"$regex": name, "$options": "i"}},
		options.Find().SetAllowDiskUse(true).SetLimit(limit).SetSkip(offset).SetProjection(bson.M{"_id": 0}))

	if err != nil {
		fmt.Println("error at get bus stops: ", err)
		return result
	}

	err = cur.All(ctx, &result)
	fmt.Println("error cursor at get bus lines by bus stop name: ", err)
	defer cur.Close(ctx)

	return result
}
