package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	requesterImpl "github.com/vandensudarsono/bus-system/adaptors/clients"
	"github.com/vandensudarsono/bus-system/adaptors/controllers"
	"github.com/vandensudarsono/bus-system/adaptors/presenter"
	repositoryimpl "github.com/vandensudarsono/bus-system/adaptors/repositoryImpl"
	"github.com/vandensudarsono/bus-system/config"
	mongodb "github.com/vandensudarsono/bus-system/infrastructures/mongoDB"
	"github.com/vandensudarsono/bus-system/usecase"
)

func main() {
	config.LoadConfig(".")
	err := mongodb.InitDataStore(
		viper.GetString("mongoDB.user"),
		viper.GetString("mongoDB.pass"),
		viper.GetString("mongoDB.hostname"),
		viper.GetInt("mongoDB.port"),
	)

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	App := fiber.New()
	//create group api
	api := App.Group("/api")

	//repository
	md := repositoryimpl.NewMongoImpl(mongodb.GetDataStoreClient())
	//presenter
	p := presenter.NewPresenter()
	//client
	cl := requesterImpl.NewClients(&fiber.Client{})
	//usecase
	uc := usecase.NewBusInteractor(md, cl, p)
	//controller
	c := controllers.NewBusLinesController(uc)

	//create route api
	api.Get("/get-available-bus", c.GetAvailableLines)
	api.Get("/get-busline-detail", c.GetBuslineDetailById)
	api.Get("/get-busline-by-busstop-name", c.GetBuslineByBusStopName)
	api.Get("/get-busline-detail-by-busstop-id", c.GetBuslinesDetailByBusStopId)

	//api listen to given addres and port
	err = App.Listen(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")))
	if err != nil {
		fmt.Println(err)
	}
}
