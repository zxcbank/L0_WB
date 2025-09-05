package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	. "go-template-microservice-v2/internal/features/get_order/queries"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type WebOrderHandler struct {
	kafkaWriter *kafka.Writer
	kafkaReader *kafka.Reader
}

func NewWebOrderHandler() *WebOrderHandler {
	webOrderHandler := &WebOrderHandler{kafkaWriter: kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},
	}), kafkaReader: kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:29092"},
		Topic:       "order-responses",
		GroupID:     "web-order-handler",
		StartOffset: kafka.FirstOffset,
	})}

	//go webOrderHandler.consumeKafkaMessages()

	return webOrderHandler
}

func (h *WebOrderHandler) WebOrderFormHandler(c echo.Context) error {
	fmt.Println("GET-FORM ORDER OPENED")
	err := c.Render(http.StatusOK, "order_form.html", nil)
	if err != nil {
		fmt.Printf("Rendering error: %v\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Rendering error: %v", err))
	}

	return nil
}

func (h *WebOrderHandler) WebGetOrderHandler(c echo.Context) error {
	orderID := c.FormValue("id")
	parsedID, err := uuid.Parse(orderID)

	if err != nil {
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal request"})
	}
	log.Printf("before writing message")
	err = h.kafkaWriter.WriteMessages(c.Request().Context(),
		kafka.Message{
			Topic: "order-requests",
			Key:   []byte(correlationID),
			Value: requestBytes,
		},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send request to Kafka"})
	}

	response, err := h.waitForKafkaResponse(correlationID, 10*time.Second)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	log.Printf("after writing message")

	var entity GetOrderResponse
	err = json.Unmarshal(response, &entity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse response"})
	}

	return c.Render(http.StatusOK, "order_info.html", map[string]interface{}{
		"Order": entity.Order,
	})
}

func (h *WebOrderHandler) waitForKafkaResponse(correlationID string, timeout time.Duration) ([]byte, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
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
