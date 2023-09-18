package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/vandensudarsono/bus-system/adaptors/controllers"
	"github.com/vandensudarsono/bus-system/adaptors/presenter"
	repositoryimpl "github.com/vandensudarsono/bus-system/adaptors/repositoryImpl"
	"github.com/vandensudarsono/bus-system/config"
	mongodb "github.com/vandensudarsono/bus-system/infrastructures/mongoDB"
	"github.com/vandensudarsono/bus-system/usecase"
)

func main() {
	config.LoadConfig()
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

	//service interactor
	md := repositoryimpl.NewMongoImpl(mongodb.GetDataStoreClient())
	//presenter
	p := presenter.NewPresenter()
	uc := usecase.NewBusInteractor(md, p)
	c := controllers.NewBusLinesController(uc)

	App := fiber.New()
	App.Get("/api/get-available-bus", c.GetAvailableLines)

	err = App.Listen(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")))
	if err != nil {
		fmt.Println(err)
	}
}
