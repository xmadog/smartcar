package controller

import (
	"context"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"smartcar/config"
	"smartcar/logger"
	"smartcar/model"
	"smartcar/motor"
	"smartcar/sensor"
	"strings"
	"time"
)

type Controller struct {
	Sensors map[string]sensor.Sensor
	Motor   *motor.Motor
	Config  *config.Config
	context context.Context
	cancel  context.CancelFunc
	isStart bool
	channel chan model.DataInfo
}

func NewController(conf *config.Config) *Controller {
	if conf.Controller.Motors == nil || len(conf.Controller.Motors) == 0 {
		panic("controller.NewController -> conf.Controller.Motors is nil.")
	}

	var sensorMap map[string]sensor.Sensor = nil
	if conf.Controller.Sensors != nil && len(conf.Controller.Sensors) > 0 {
		sensorMap = make(map[string]sensor.Sensor)
		for _, item := range conf.Controller.Sensors {
			switch item.Type {
			case sensor.Infrared:
				sensorMap[item.Name] = sensor.NewInfrared(
					item.Name,
					item.Type,
					rpio.Pin(item.Output[0]),
					time.Duration(500))
			case sensor.Ultrasound:
				//do something
				break
			default:
				panic(fmt.Sprintf("controller.NewController -> SensorType[%d] is unknown.",
					item.Type))

			}
		}
	}

	controller := Controller{
		Config:  conf,
		Sensors: sensorMap,
		context: nil,
		isStart: false,
	}
	return &controller
}

func (this *Controller) Start(ctx context.Context) {
	if !this.isStart {
		this.context, this.cancel = context.WithCancel(ctx)
		this.channel = make(chan model.DataInfo)

		rpio.Open()
		this.Motor = motor.NewMotor(
			motor.NewMotorPins(
				this.Config.Controller.Motors[0].In1,
				this.Config.Controller.Motors[0].In2,
				this.Config.Controller.Motors[0].Pwm),
			motor.NewMotorPins(
				this.Config.Controller.Motors[1].In1,
				this.Config.Controller.Motors[1].In2,
				this.Config.Controller.Motors[1].Pwm),
			motor.NewMotorPins(
				this.Config.Controller.Motors[2].In1,
				this.Config.Controller.Motors[2].In2,
				this.Config.Controller.Motors[2].Pwm),
			motor.NewMotorPins(
				this.Config.Controller.Motors[3].In1,
				this.Config.Controller.Motors[3].In2,
				this.Config.Controller.Motors[3].Pwm),
		)

		if this.Sensors != nil {
			for _, item := range this.Sensors {
				item.Start(this.channel, this.context)
			}
		}

		go func(ctx context.Context) {
			logger.GetLogger(ctx).Info("controller -> start")
			select {
			case <-ctx.Done():
				close(this.channel)
				this.channel = nil
				this.context = nil
				logger.GetLogger(ctx).Info("controller -> end")
			case operation := <-this.channel:
				switch operation.Type {
				case model.Motor:
					this.setMotor(&operation)
				default:
					panic(fmt.Sprintf("controller.Start -> operation.Type[%d] is unknown.", operation.Type))
				}

			}
		}(this.context)
		this.isStart = true
		logger.GetLogger(this.context).Info("controller.Start -> called")
	}
}

func (this *Controller) Stop() {
	if this.isStart {
		this.cancel()
		rpio.Close()
		this.isStart = false
		logger.GetLogger(this.context).Info("controller.End -> called")
	}
}

func (this *Controller) Send(operaiton *model.DataInfo) {
	if this.isStart && operaiton != nil {
		this.channel <- *operaiton
	}
}

func (this *Controller) setMotor(operation *model.DataInfo) {
	value := strings.ToLower(operation.Data.(string))
	switch value {
	case "forward":
		this.Motor.Forward()
	case "turnright":
		this.Motor.TurnRight()
	case "backward":
		this.Motor.Backward()
	case "turnleft":
		this.Motor.TurnLeft()
	default:
		panic(fmt.Sprintf("controller.setMotoer -> operation[%s] is unknown.", value))
	}
}
