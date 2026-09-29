package componenteconexioncola

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

type RabbitPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

type NotificacionCancion struct {
	Titulo  string `json:"titulo"`
	Mensaje string `json:"mensaje"`
}

func NewRabbitPublisher() (*RabbitPublisher, error) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://admin:1234@10.150.10.73/"
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("error abriendo canal de RabbitMQ: %w", err)
	}

	q, err := ch.QueueDeclare("notificaciones_audios", true, false, false, false, nil)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("error declarando cola: %w", err)
	}

	return &RabbitPublisher{conn: conn, channel: ch, queue: q}, nil
}

func (p *RabbitPublisher) PublicarNotificacion(msg NotificacionCancion) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error convirtiendo mensaje a JSON: %w", err)
	}

	if err := p.channel.Publish("", p.queue.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}); err != nil {
		return fmt.Errorf("error publicando mensaje: %w", err)
	}

	fmt.Printf("Notificación enviada a RabbitMQ: %s\n", body)
	return nil
}

func (p *RabbitPublisher) Cerrar() {
	if p == nil {
		return
	}
	if p.channel != nil {
		_ = p.channel.Close()
	}
	if p.conn != nil {
		_ = p.conn.Close()
	}
}
