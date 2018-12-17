package sensor

import (
	"context"
	"smartcar/model"
)

type Sensor interface {
	Start(chan model.DataInfo, context.Context)
	Stop(context.Context)
}
