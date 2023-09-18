package repositoryimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vandensudarsono/bus-system/config"
	"github.com/vandensudarsono/bus-system/domain/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewMongoImpl(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://user:pass@172.19.0.4:27017/BusSistem"))
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Run("test_mongo", func(t *testing.T) {
		got := NewMongoImpl(client)
		assert.Equal(t, client, got.client)
	})
}

func TestGetBuslines(t *testing.T) {
	config.LoadConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://user:pass@172.19.0.4:27017/BusSistem"))
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.ReadFile("./sample_data/BusSistem.available-bus.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var expect []models.BusLine
	json.Unmarshal(file, &expect)

	db := NewMongoImpl(client)
	t.Run("test_get_buslines", func(t *testing.T) {
		got := db.GetBuslines(ctx, 0, 10)
		fmt.Println(got)
		assert.Equal(t, expect, got)
	})
}
