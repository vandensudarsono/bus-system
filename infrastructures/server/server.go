package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	requesterImpl "github.com/vandensudarsono/bus-system/adaptors/clients"
	"github.com/vandensudarsono/bus-system/adaptors/controllers"
	"github.com/vandensudarsono/bus-system/adaptors/presenter"
	repositoryimpl "github.com/vandensudarsono/bus-system/adaptors/repositoryImpl"
	mongodb "github.com/vandensudarsono/bus-system/infrastructures/mongoDB"
	handlers "github.com/vandensudarsono/bus-system/infrastructures/server/routers/v1"
	"github.com/vandensudarsono/bus-system/usecase"
)

type Server struct {
	Server *fiber.App
	// route  *fiber.Router
}

var S *fiber.App

func NewServer() {
	config := fiber.Config{
		AppName:      "Bus-System",
		Prefork:      false,
		ServerHeader: "bus-system",
	}

	S = fiber.New(config)

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

	//add router
	handlers.SetupBuslineHandler(S, c)

	//add middleware

}

func StartServer() error {
	//return S.Server().ListenAndServe(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")))
	return S.Listen(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")))
}

func StopServer() error {
	return S.Shutdown()
}
