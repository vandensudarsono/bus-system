package main

import (
	"github.com/vandensudarsono/bus-system/cmd"
	_ "github.com/vandensudarsono/bus-system/swagger"
)

// @title Bus Line Sample API
// @version 1.0
// @description This is a sample swagger for Busline Sample API
// @contact.name edi abdullah
// @contatct.email this.email@yahoo.com
// @host localhost:8081
// @BasePath /api/
func main() {
	command := cmd.NewCommand()
	command.Start()
}
