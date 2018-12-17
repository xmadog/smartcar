package controller

type ControllerType byte
const (
	Motor ControllerType = iota
	Sensor
)

type ControllerData struct {
	Type ControllerType
	Data interface{}
}
