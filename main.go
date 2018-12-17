package main

import (
	"smartcar/config"
	"smartcar/logger"
)

func main() {
	var conf config.Config
	config.NewConfig(&conf)
	log := logger.NewLogger(conf.Logger)
	log.Info("smartcar -> start")

	/*
	car,err := motor.NewController(
		motor.NewMotorPins(2,3,4),
		motor.NewMotorPins(17,27,22),
		motor.NewMotorPins(10,9,11),
		motor.NewMotorPins(14,15,18),
	)
	if(err != nil) {
	    fmt.Println("smartcar.error -> %s",err)
	    return
    	}
	defer car.Close()

	for i:= 0;i < 30;i++ {
		car.Forward()
		fmt.Println("smartcar -> forward")
		time.Sleep(1*time.Second)
	}
	*/
	log.Info("smartcar -> end")
}
