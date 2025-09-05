package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	. "go-template-microservice-v2/internal/features/queries"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type OrderEndpoint struct {
	kafkaWriter *kafka.Writer
	kafkaReader *kafka.Reader
}

func NewOrderEndpoint() *OrderEndpoint {
	webOrderHandler := &OrderEndpoint{kafkaWriter: kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},
	}), kafkaReader: kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:29092"},
		Topic:       "order-responses",
		GroupID:     "order-endpoint",
		StartOffset: kafka.FirstOffset,
	})}

	return webOrderHandler
}

func (h *OrderEndpoint) OrderForm(c echo.Context) error {
	err := c.Render(http.StatusOK, "order_form.html", nil)
	if err != nil {
		log.Printf("Rendering error: %v\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Rendering error: %v", err))
	}

	return nil
}

func (h *OrderEndpoint) OrderShowResult(c echo.Context) error {
	orderID := c.FormValue("id")
	parsedID, err := uuid.Parse(orderID)

	if err != nil {
		log.Println(err)
		return c.Render(http.StatusOK, "order_form.html", map[string]interface{}{
			"Error": "Неверный формат ID заказа",
		})
	}

	correlationID := uuid.New().String()

	kafkaRequest := KafkaRequest{
		ID:            parsedID.String(),
		Type:          "get_entity",
		CorrelationID: correlationID,
	}

	requestBytes, err := json.Marshal(kafkaRequest)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal request"})
	}

	err = h.kafkaWriter.WriteMessages(c.Request().Context(),
		kafka.Message{
			Topic: "order-requests",
			Key:   []byte(correlationID),
			Value: requestBytes,
		},
	)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send request to Kafka"})
	}

	response, err := h.waitForKafkaResponse(correlationID, 10*time.Second, c.Request().Context())
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var entity GetOrderResponse
	err = json.Unmarshal(response, &entity)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse response"})
	}

	return c.Render(http.StatusOK, "order_info.html", map[string]interface{}{
		"Order": entity.Order,
	})
}

func (h *OrderEndpoint) waitForKafkaResponse(correlationID string, timeout time.Duration, ctx context.Context) ([]byte, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	messageChan := make(chan *kafka.Message, 1)
	errorChan := make(chan error, 1)

	go func() {
		for {
			msg, err := h.kafkaReader.ReadMessage(ctx)
			if err != nil {
				errorChan <- err
				return
			}

			if string(msg.Key) == correlationID {
				messageChan <- &msg
				return
			}
		}
	}()

	select {
	case msg := <-messageChan:
		return msg.Value, nil
	case err := <-errorChan:
		return nil, err
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout waiting for response")
	}
}
