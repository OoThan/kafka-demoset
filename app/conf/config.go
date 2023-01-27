package conf

import (
	"kafka-demoset/app/internal/logger"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	kafka struct {
		Addr string `yaml:"addr"`
	}

	telegram struct {
		TokenID string `yaml:"tokenID"`
		GroupID string `yaml:"groupID"`
	}
)

var (
	_c struct {
		Kafka    kafka    `yaml:"kafka"`
		Telegram telegram `yaml:"telegram"`
	}
)

func init() {
	data, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		logger.Sugar.Error(err)
	}

	if err := yaml.Unmarshal(data, &_c); err != nil {
		logger.Sugar.Error(err)
	}
}

func Kafka() *kafka {
	return &_c.Kafka
}

func Telegram() *telegram {
	return &_c.Telegram
}
