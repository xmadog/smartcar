package main

import (
	"context"
	"smartcar/config"
	"smartcar/controller"
	"smartcar/logger"
)

func main() {
	var conf config.Config
	config.NewConfig(&conf)
	log := logger.NewLogger(conf.Logger)
	log.Info("smartcar -> start")

	carController := controller.NewController(&conf)
	carController.Start(context.Background())
	defer carController.Stop()

	log.Info("smartcar -> end")
}
