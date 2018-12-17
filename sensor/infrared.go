package sensor

import (
	"context"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"smartcar/logger"
	"smartcar/model"
	"time"
)

type InfraredInfo struct {
	Name     string
	Type     SensorType
	Output   rpio.Pin
	Duration time.Duration
	isStart  bool
	channel  chan model.DataInfo
}

func NewInfrared(name string,
	sersorType SensorType,
	output rpio.Pin,
	duration time.Duration) *InfraredInfo {
	infrared := InfraredInfo{
		Name:     name,
		Type:     sersorType,
		Output:   output,
		Duration: duration,
		isStart:  false,
	}

	return &infrared
}

func (infrared *InfraredInfo) Start(channel chan model.DataInfo,
	ctx context.Context) {
	if !infrared.isStart {
		infrared.Output.Input()
		infrared.channel = channel
		infrared.isStart = true

		go func(ctx context.Context) {
			logger.GetLogger(ctx).Info(
				fmt.Sprintf("InfraredSensor[%s] -> start", infrared.Name))
			defer close(infrared.channel)
			for ctx.Done() == nil {
				state := (infrared.Output.Read() == rpio.Low)
				infrared.channel <- model.DataInfo{
					Type: model.Sensor,
					Data: SensorData{
						Name: infrared.Name,
						Type: infrared.Type,
						Data: state,
					},
				}
				logger.GetLogger(ctx).Debug(
					fmt.Sprintf("InfraredSensor[%s].recived -> %d", infrared.Name, state))
				time.Sleep(infrared.Duration * time.Millisecond)
			}
		}(ctx)
	}
}

func (infrared *InfraredInfo) Stop(ctx context.Context) {
	infrared.isStart = false
	logger.GetLogger(ctx).Info(
		fmt.Sprintf("InfraredSensor[%s] -> stop", infrared.Name))
}
