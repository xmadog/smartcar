package main

import (
	"context"
	"smartcar/config"
	"smartcar/controller"
	"smartcar/logger"
	"smartcar/model"
	"time"
)

func main() {
	var conf config.Config
	config.NewConfig(&conf)
	log := logger.NewLogger(conf.Logger)
	ctx := logger.SetLogger(context.Background(), log)
	log.Info("smartcar -> start")

	carController := controller.NewController(&conf)
	carController.Start(ctx)
	defer carController.Stop()

	for {
		carController.Send(&model.DataInfo{
			Type: model.Motor,
			Data: "forward",
		})
		time.Sleep(500 * time.Millisecond)
	}

	log.Info("smartcar -> end")
}
