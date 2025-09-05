package queries

import (
	"context"
	"encoding/json"
	"go-template-microservice-v2/internal/data/contracts"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type OrderService struct {
	Repository    contracts.IOrderRepository
	Ctx           context.Context
	kafkaReader   *kafka.Reader
	kafkaWriter   *kafka.Writer
	responseChans map[string]chan []byte
	mu            sync.RWMutex
}

func NewOrderService(
	repository contracts.IOrderRepository,
	ctx context.Context) *OrderService {

	handler := &OrderService{
		Repository: repository,
		Ctx:        ctx,
		kafkaReader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{"localhost:29092"},
			Topic:       "order-requests",
			GroupID:     "order-service",
			StartOffset: kafka.FirstOffset,
		}),
		kafkaWriter: kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{"localhost:29092"},
		}),
		responseChans: make(map[string]chan []byte),
	}

	go handler.consumeKafkaMessages()

	return handler
}

func (handler *OrderService) HandleGetRequest(ctx context.Context, id uuid.UUID) (*GetOrderResponse, error) {
	getOrdersResponse := &GetOrderResponse{}

	result, err := handler.Repository.GetOrder(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	getOrdersResponse.Order = GetOrderResponseItem{
		Id:                 result.Id,
		Track_number:       result.Track_number,
		Entry:              result.Entry,
		Locale:             result.Locale,
		Internal_signature: result.Internal_signature,
		Custromer_id:       result.Custromer_id,
		Delivery_service:   result.Delivery_service,
		Shardkey:           result.Shardkey,
		Sm_id:              result.Sm_id,
		Date_created:       result.Date_created,
		Oof_shard:          result.Oof_shard,
	}

	return getOrdersResponse, nil
}

func (handler *OrderService) consumeKafkaMessages() {

	if handler.kafkaReader == nil {
		log.Fatal("Kafka reader is not initialized")
		return
	}

	for {
		msg, err := handler.kafkaReader.ReadMessage(handler.Ctx)

		if err != nil {
			log.Printf("Error reading Kafka message: %v", err)
			continue
		}
		if msg.Topic != "order-requests" {
			continue
		}

		go func(message kafka.Message) {

			var request KafkaRequest
			err := json.Unmarshal(message.Value, &request)
			if err != nil {
				log.Printf("Kafka message: %v", err)
				return
			}

			id, err := uuid.Parse(request.ID)
			if err != nil {
				log.Printf("Error parsing Kafka message: %v", err)
				return
			}
			response, err := handler.HandleGetRequest(handler.Ctx, id)
			if err != nil {
				log.Printf("Error processing request: %v", err)
				return
			}

			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Printf("Error marshaling response: %v", err)
				return
			}

			err = handler.kafkaWriter.WriteMessages(handler.Ctx,
				kafka.Message{
					Topic: "order-responses",
					Key:   message.Key,
					Value: responseBytes,
				},
			)
			if err != nil {
				log.Printf("Error sending response to Kafka: %v", err)
			}
		}(msg)
	}
}

type KafkaRequest struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	CorrelationID string `json:"correlation_id"`
}
