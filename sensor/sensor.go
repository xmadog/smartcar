package sensor

import "context"

type Sensor interface {
	Start(chan interface{},context.Context)
	Stop(context.Context)
}
