package kafka

type KafkaRequest struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	CorrelationID string `json:"correlation_id"`
}
