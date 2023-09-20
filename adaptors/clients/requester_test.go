package clients

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vandensudarsono/bus-system/config"
)

func TestGetRunningBusInABusLines(t *testing.T) {
	config.LoadConfig("../../.")
	client := fiber.AcquireClient()

	req := NewClients(client)

	t.Run("test running bus", func(t *testing.T) {
		result, err := req.GetRunningBusInABusLines(context.TODO(), "44478")
		if err != nil {
			log.Fatalf("error %+v", err)
		}

		fmt.Println(result)

		assert.Equal(t, err, nil)
	})

}
