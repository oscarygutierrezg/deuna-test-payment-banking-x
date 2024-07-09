package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BootstrapServers string
	GroupID          string
	ConsumerTopic    string
	ProducerTopic    string
	BaseUrl          string
	ContentType      string
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("bootstrapServers", "localhost:9092")
	viper.SetDefault("groupId", "payments-group")
	viper.SetDefault("consumerTopic", "com.deuna.payment.payment_banking_x.v1.payments")
	viper.SetDefault("producerTopic", "com.deuna.payment.payment.v1.payments.updated")
	viper.SetDefault("baseUrl", "http://localhost:5001/v1")
	viper.SetDefault("contentType", "application/json")

	viper.AutomaticEnv()

	config := &Config{
		BootstrapServers: viper.GetString("bootstrapServers"),
		GroupID:          viper.GetString("groupId"),
		ConsumerTopic:    viper.GetString("consumerTopic"),
		ProducerTopic:    viper.GetString("producerTopic"),
		BaseUrl:          viper.GetString("baseUrl"),
		ContentType:      viper.GetString("contentType"),
	}

	return config, nil
}
