package config

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"smartcar/sensor"
)

type Config struct {
	Controller struct {
		Motors []struct {
			In1 uint8
			In2 uint8
			Pwm uint8
		}
		Sensors []struct {
			Name string
			Type sensor.SensorType
			Output []uint8
		}
	}
	Logger zap.Config
}

func NewConfig(config *Config)  {
	path, err := os.Executable()

	if err != nil {
		panic(err)
	}
	folderPath := filepath.Dir(path)
	configPath := filepath.Join(folderPath,"config.json")

	if _,err := os.Stat(configPath);os.IsNotExist(err) {
		panic(err)
	}

	file,err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config);err != nil {
		panic(err)
	}
}

func GetConfig(ctx context.Context) *Config {
	return ctx.Value("config").(*Config)
}

func SetConfig(ctx context.Context,config *Config) context.Context {
	return context.WithValue(ctx,"config",config)
}