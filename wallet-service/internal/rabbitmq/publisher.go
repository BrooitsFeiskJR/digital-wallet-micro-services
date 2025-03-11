package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewPublisher() (*Publisher, error) {
	connString := os.Getenv("RABBITMQ_CONN_STRING")
	if connString == "" {
		return nil, errors.New("env RABBITMQ_CONN_STRING is required")
	}

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

func (p *Publisher) Close() error {
	if err := p.channel.Close(); err != nil {
		return err
	}
	return nil
}

func (p *Publisher) DeclareQueue(name string) (amqp.Queue, error) {
	return p.channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (p *Publisher) PublishMessage(ctx context.Context, queueName string, message any) error {
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
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (p *Publisher) PulishTransactionCreated(eventType, userID string, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := map[string]any{
		"event_type": eventType,
		"user_id":    userID,
		"amount":     amount,
		"created_at": time.Now(),
	}

	return p.PublishMessage(ctx, "transactions_queue", message)
}
