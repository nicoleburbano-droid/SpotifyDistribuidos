package componenteconexioncola

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

const colaNotificaciones = "notificaciones_audios"

type NotificacionCancion struct {
	Titulo  string `json:"titulo"`
	Mensaje string `json:"mensaje"`
}

type RabbitPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	mu      sync.Mutex
}

func NewRabbitPublisher() (*RabbitPublisher, error) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://admin:1234@localhost:5672/"
	}
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %w", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("error abriendo canal de RabbitMQ: %w", err)
	}
	queue, err := channel.QueueDeclare(colaNotificaciones, true, false, false, false, nil)
	if err != nil {
		_ = channel.Close()
		_ = conn.Close()
		return nil, fmt.Errorf("error declarando la cola %s: %w", colaNotificaciones, err)
	}
	return &RabbitPublisher{conn: conn, channel: channel, queue: queue}, nil
}

func (p *RabbitPublisher) PublicarNotificacion(mensaje NotificacionCancion) error {
	body, err := json.Marshal(mensaje)
	if err != nil {
		return fmt.Errorf("error convirtiendo la notificación a JSON: %w", err)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if err := p.channel.Publish("", p.queue.Name, false, false, amqp.Publishing{ContentType: "application/json", Body: body}); err != nil {
		return fmt.Errorf("error publicando notificación en RabbitMQ: %w", err)
	}
	fmt.Printf("Notificación enviada a RabbitMQ: %s\n", body)
	return nil
}

func (p *RabbitPublisher) Cerrar() {
	if p == nil {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.channel != nil {
		_ = p.channel.Close()
	}
	if p.conn != nil {
		_ = p.conn.Close()
	}
}
