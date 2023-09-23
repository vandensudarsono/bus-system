package main

import "github.com/vandensudarsono/bus-system/cmd"

func main() {
	command := cmd.NewCommand()
	command.Start()
}
