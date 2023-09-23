package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vandensudarsono/bus-system/config"
	mongodb "github.com/vandensudarsono/bus-system/infrastructures/mongoDB"
	"github.com/vandensudarsono/bus-system/infrastructures/server"
)

var (
	text = `
	████████████████████████████████████████████████████
	█▄─▄─▀█▄─██─▄█─▄▄▄▄███▄─▄███▄─▄█▄─▀█▄─▄█▄─▄▄─█─▄▄▄▄█
	██─▄─▀██─██─██▄▄▄▄─████─██▀██─███─█▄▀─███─▄█▀█▄▄▄▄─█
	▀▄▄▄▄▀▀▀▄▄▄▄▀▀▄▄▄▄▄▀▀▀▄▄▄▄▄▀▄▄▄▀▄▄▄▀▀▄▄▀▄▄▄▄▄▀▄▄▄▄▄▀`
)

type Command struct {
	rootCmd *cobra.Command
}

// NewCommand the command line boot loader
func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:   "bus-system",
		Short: "bus system that enable user interactions to using bus",
		Long:  "Bus system that user can see the incomming bus in times",
	}

	return &Command{
		rootCmd: rootCmd,
	}
}

func (c *Command) Start() {
	var rootCommands = []*cobra.Command{
		{
			Use:   "start",
			Short: "start bus system api",
			Long:  "start bus system to provide informations to users",
			PreRun: func(cmd *cobra.Command, args []string) {
				fmt.Println(text)
				//init config
				config.LoadConfig(".")

				if err := mongodb.InitDataStore(
					viper.GetString("mongoDB.user"),
					viper.GetString("mongoDB.pass"),
					viper.GetString("mongoDB.hostname"),
					viper.GetInt("mongoDB.port"),
				); err != nil {
					fmt.Printf("Errot at init datastore: %+v", err)
					os.Exit(1)
				}

			},
			Run: func(cmd *cobra.Command, args []string) {
				server.NewServer()
				err := server.StartServer()
				if err != nil {
					fmt.Println("Error at star server: ", err)
					os.Exit(1)
				}
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				//stop database
				err := mongodb.StopDataStoreClient()
				if err != nil {
					fmt.Println("Error at stop datastore: ", err)
				}
				//stop server
				err = server.StopServer()
				if err != nil {
					fmt.Println("Errot at stop server: ", err)
				}
			},
		},
	}

	for _, command := range rootCommands {
		c.rootCmd.AddCommand(command)
	}

	err := c.rootCmd.Execute()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

// GetRoot the command line service
func (c *Command) GetRoot() *cobra.Command {
	return c.rootCmd
}
