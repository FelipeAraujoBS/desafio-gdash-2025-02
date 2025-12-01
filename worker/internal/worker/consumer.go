package worker

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go-rabbitmq-worker/internal/config"
)

func StartWorker(cfg *config.Config) {
	for {
		func() {
			conn, err := amqp.Dial(cfg.RabbitMQURL)
			if err != nil {
				log.Printf("Failed to connect: %v", err)
				time.Sleep(5 * time.Second)
				return
			}
			// defer no escopo da função anônima: será executado quando a função anônima terminar
			defer func() {
				if err := conn.Close(); err != nil {
					log.Printf("Error closing conn: %v", err)
				}
			}()

			ch, err := conn.Channel()
			if err != nil {
				log.Printf("Failed to open channel: %v", err)
				return
			}
			defer func() {
				if err := ch.Close(); err != nil {
					log.Printf("Error closing channel: %v", err)
				}
			}()

			log.Println("Connected")

			if err := consumeMessages(ch, cfg); err != nil {
				log.Printf("Consume error: %v", err)
			}
		}()

		log.Println("Reconnecting in 5s...")
		time.Sleep(5 * time.Second)
	}
}

func consumeMessages(ch *amqp.Channel, cfg *config.Config) error {
	_, err := ch.QueueDeclarePassive(
		cfg.QueueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	if err := ch.Qos(1, 0, false); err != nil {
		return err
	}

	msgs, err := ch.Consume(
		cfg.QueueName,
		"",    // consumer tag
		false, // auto-ack = false (manual ack)
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Processa mensagens até msgs ser fechado (quando a conexão/chanel é fechado)
	for d := range msgs {
		log.Printf("Received a message: %s", string(d.Body))
		if ProcessAndSend(d.Body, cfg) {
			if err := d.Ack(false); err != nil {
				log.Printf("Ack error: %v", err)
			} else {
				log.Printf("Message processed and acknowledged")
			}
		} else {
			// simples: requeue (true). Em produção, prefira controlar tentativas e DLQ.
			if err := d.Nack(false, true); err != nil {
				log.Printf("Nack error: %v", err)
			} else {
				log.Printf("Message processing failed, message requeued")
			}
		}
	}

	// quando chegamos aqui msgs foi fechado -> encerramos consumeMessages para permitir reconexão
	return nil
}
