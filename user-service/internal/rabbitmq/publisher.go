package rabbitmq

import (
	"context"
	"encoding/json"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher manages the RabbitMQ connection and provides methods to publish messages
type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewPublisher creates and returns a new Publisher instance
func NewPublisher() (*Publisher, error) {
	connString := os.Getenv("RABBITMQ_CONN_STRING")
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Publisher{
		conn:    conn,
		channel: ch,
	}, nil
}

// Close closes the channel and connection
func (p *Publisher) Close() error {
	if err := p.channel.Close(); err != nil {
		return err
	}
	return p.conn.Close()
}

// DeclareQueue declares a queue if it doesn't exist
func (p *Publisher) DeclareQueue(name string) (amqp.Queue, error) {
	return p.channel.QueueDeclare(
		name,  // queue name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

// PublishMessage publishes a message to a queue
func (p *Publisher) PublishMessage(ctx context.Context, queueName string, message interface{}) error {
	q, err := p.DeclareQueue(queueName)
	if err != nil {
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.channel.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// PublishUserCreated publishes a message when a new user is created
func (p *Publisher) PublishUserCreated(userID string, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := map[string]any{
		"event_type": "user_created",
		"user_id":    userID,
		"email":      email,
		"timestamp":  time.Now(),
	}

	return p.PublishMessage(ctx, "user_queue", message)
}
