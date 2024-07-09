package main

import (
	"context"
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	banking_x "payment-banking-x/internal/client/banking-x"
	"payment-banking-x/internal/client/http"
	"payment-banking-x/internal/config"
	"payment-banking-x/internal/kafka/consumer"
	"payment-banking-x/internal/kafka/producer"
	"payment-banking-x/internal/service"
	"payment-banking-x/pkg/util"
)

func init() {
	flag.Int("port", 5001, "kafka port.")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.AutomaticEnv()
	_ = viper.BindPFlags(pflag.CommandLine)
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		panic(err)
	}

	err = util.CreateTopic(context.Background(), cfg, cfg.ProducerTopic)
	if err != nil {
		log.Fatalf("Failed to create topic: %s", cfg.ProducerTopic)
		panic(err)
	}
	err = util.CreateTopic(context.Background(), cfg, cfg.ConsumerTopic)
	if err != nil {
		log.Fatalf("Failed to create topic: %s", cfg.ConsumerTopic)
		panic(err)
	}

	restClient := http.NewRestClient(cfg)
	bankingXClient := banking_x.NewClient(restClient)

	paymentProducer, err := producer.NewPaymentProducer(cfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka paymentProducer: %v", err)
		panic(err)
	}

	paymentConsumer, err := consumer.NewPaymentConsumer(cfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka paymentConsumer: %v", err)
		panic(err)
	}

	paymentService := service.NewPaymentService(bankingXClient, paymentProducer)
	refundService := service.NewRefundService(bankingXClient, paymentProducer)

	services := &service.Services{
		Payment: paymentService,
		Refund:  refundService,
	}

	paymentConsumer.Consume(services)

}
