package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config armazena todas as variáveis de ambiente carregadas
type Config struct {
	RabbitMQURL  string
	QueueName    string
	NestJSAPIURL string
	HTTPTimeout  time.Duration // Tempo limite para requisições HTTP
}

func LoadConfig() *Config {
	err := godotenv.Load() 
	if err != nil {
		log.Fatalf("❌ Erro ao carregar o arquivo .env: %v", err)
	}

	cfg := &Config{
		RabbitMQURL:  os.Getenv("RABBITMQ_URL"),
		QueueName:    os.Getenv("QUEUE_NAME"),
		NestJSAPIURL: os.Getenv("NESTJS_API_URL"),
		HTTPTimeout:  10 * time.Second,
	}

	if cfg.RabbitMQURL == "" || cfg.QueueName == "" || cfg.NestJSAPIURL == "" {
		log.Fatal("❌ Erro de Configuração: Uma ou mais variáveis de ambiente críticas não foram definidas.")
	}
	
	return cfg
}