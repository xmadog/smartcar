package model

type DataType byte

const (
	Motor DataType = iota
	Sensor
)

type DataInfo struct {
	Type DataType
	Data interface{}
}
