package rabbitmq

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) {
	// Ensure RABBITMQ_CONN_STRING is set for tests
	if os.Getenv("RABBITMQ_CONN_STRING") == "" {
		os.Setenv("RABBITMQ_CONN_STRING", "amqp://guest:guest@localhost:5672/")
	}
}

func cleanup(t *testing.T, queueName string) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_STRING"))
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	// Delete the test queue
	_, err = ch.QueueDelete(queueName, false, false, false)
	if err != nil {
		t.Logf("Failed to delete queue %s: %v", queueName, err)
	}
}

func TestNewPublisher(t *testing.T) {
	setupTest(t)

	pub, err := NewPublisher()
	require.NoError(t, err)
	require.NotNil(t, pub)
	require.NotNil(t, pub.conn)
	require.NotNil(t, pub.channel)

	err = pub.Close()
	require.NoError(t, err)
}

func TestDeclareQueue(t *testing.T) {
	setupTest(t)

	pub, err := NewPublisher()
	require.NoError(t, err)
	defer pub.Close()

	queueName := "test_queue"
	queue, err := pub.DeclareQueue(queueName)
	require.NoError(t, err)
	assert.Equal(t, queueName, queue.Name)

	defer cleanup(t, queueName)
}

func TestPublishMessage(t *testing.T) {
	setupTest(t)

	pub, err := NewPublisher()
	require.NoError(t, err)
	defer pub.Close()

	queueName := "test_publish_queue"
	message := map[string]string{
		"test_key": "test_value",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pub.PublishMessage(ctx, queueName, message)
	require.NoError(t, err)

	// Verify the message was published
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_STRING"))
	require.NoError(t, err)
	defer conn.Close()

	ch, err := conn.Channel()
	require.NoError(t, err)
	defer ch.Close()

	msg, ok, err := ch.Get(queueName, true)
	require.NoError(t, err)
	require.True(t, ok, "No message received from queue")

	var receivedMsg map[string]string
	err = json.Unmarshal(msg.Body, &receivedMsg)
	require.NoError(t, err)
	assert.Equal(t, message["test_key"], receivedMsg["test_key"])

	defer cleanup(t, queueName)
}

func TestPublishUserCreated(t *testing.T) {
	setupTest(t)

	pub, err := NewPublisher()
	require.NoError(t, err)
	defer pub.Close()

	userID := "test-user-123"
	email := "test@example.com"

	err = pub.PublishUserCreated(userID, email)
	require.NoError(t, err)

	// Verify the message was published with correct format
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_STRING"))
	require.NoError(t, err)
	defer conn.Close()

	ch, err := conn.Channel()
	require.NoError(t, err)
	defer ch.Close()

	msg, ok, err := ch.Get("user_queue", true)
	require.NoError(t, err)
	require.True(t, ok, "No message received from queue")

	var receivedMsg map[string]interface{}
	err = json.Unmarshal(msg.Body, &receivedMsg)
	require.NoError(t, err)
	assert.Equal(t, "user_created", receivedMsg["event_type"])
	assert.Equal(t, userID, receivedMsg["user_id"])
	assert.Equal(t, email, receivedMsg["email"])
	assert.NotNil(t, receivedMsg["created_at"])

	defer cleanup(t, "user_queue")
}
