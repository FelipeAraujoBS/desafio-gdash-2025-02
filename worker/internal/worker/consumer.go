package worker

import(
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"go-rabbitmq-worker/internal/config"
)

func StartWorker(cfg *config.Config){
	for{
		conn, err := amqp.Dial(cfg.RabbitMQ.URL)
		if err != nil {
			log.Printf("Failed to connect to RabbitMQ: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		defer conn.Close()

		log.Println("Connected to RabbitMQ")

		if err := consumeMessages(conn, cfg); err != nil {
			log.Printf("Error consuming messages: %v", err)
		}

		log.Println("Reconnecting to RabbitMQ...")
		time.Sleep(5 * time.Second)
	}
}

func consumeMessages(conn *amqp.Connection, cfg *config.Config) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		QueueName,
		true, //Duravél
		false, //auto-delete
		false, //Exclusivo
		false, //no-wait
		nil, // argumentos
	)
	if err != nil {
		return err
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		QueueName,
		"", // consumer tag
		false, // auto-ack - configuração manual	
		false, // exclusivo
		false, // não-local
		false, // no-wait
		nil, // argumentos
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// Simula o processamento da mensagem
			if processAndSend(d.Body){
				d.Ack(false)
				log.Printf("Message processed and acknowledged: %s", d.Body)
			} else{
				d.Nack(false, true)
				log.Printf("Message processing failed, message requeued: %s", d.Body)
			}
		}
	}()

	<-forever
	return nil
}
