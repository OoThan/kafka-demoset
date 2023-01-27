package internal

import (
	"kafka-demoset/app/conf"
	"kafka-demoset/app/internal/services"
)

var (
	TelegramBot *services.TelegramBot
)

type KafkaService struct {
	telegram *services.TelegramBot
}

func NewKafkaService() *KafkaService {
	telegramBot := services.NewTelegramBot(conf.Telegram().TokenID, conf.Telegram().GroupID)
	TelegramBot = telegramBot

	return &KafkaService{
		telegram: TelegramBot,
	}
}
