package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandler func([]byte) error

type Consumer struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	handlers  map[string]MessageHandler
	queueName string
}

func NewConsumer(queueName string) (*Consumer, error) {
	connString := os.Getenv("RABBITMQ_CONN_STRING")
	if connString == "" {
		return nil, errors.New("env RABBITMQ_CONN_STRING is required")
	}

	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	consumer := &Consumer{
		conn:      conn,
		channel:   ch,
		handlers:  make(map[string]MessageHandler),
		queueName: queueName,
	}

	return consumer, nil
}

func (c *Consumer) Close() error {
	if err := c.channel.Close(); err != nil {
		return err
	}
	return c.conn.Close()
}

func (c *Consumer) DeclareQueue() (amqp.Queue, error) {
	return c.channel.QueueDeclare(
		c.queueName,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (c *Consumer) RegisterHandler(eventType string, handler MessageHandler) {
	c.handlers[eventType] = handler
}

func (c *Consumer) Start(ctx context.Context) error {
	q, err := c.DeclareQueue()
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	msgs, err := c.channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Consumer stopped due to context cancellation")
				return
			case msg, ok := <-msgs:
				if !ok {
					log.Println("Consumer channel closed")
					return
				}

				if err := c.processMessage(msg); err != nil {
					log.Printf("Error processing message: %v", err)
					msg.Nack(false, true)
				} else {
					msg.Ack(false)
				}
			}
		}
	}()

	log.Printf("Started consuming from queue: %s", q.Name)
	return nil
}

func (c *Consumer) processMessage(msg amqp.Delivery) error {
	var message map[string]any
	if err := json.Unmarshal(msg.Body, &message); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	eventType, ok := message["event_type"].(string)
	if !ok {
		return errors.New("message missing event_type field")
	}

	handler, exists := c.handlers[eventType]
	if !exists {
		return fmt.Errorf("no handler registered for event type: %s", eventType)
	}

	return handler(msg.Body)
}
