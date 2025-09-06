package lru_cache_order

import (
	"errors"
	. "go-template-microservice-v2/internal/features/queries"
	"time"
)

type Lru_cache_order struct {
	CacheMap  map[string]Order_timestamp_pair
	CacheSize int
}

func (l *Lru_cache_order) Add(k string, v Order_timestamp_pair) {
	if len(l.CacheMap) == l.CacheSize {
		l.reduce()
	}
	l.CacheMap[k] = v
}

func (l *Lru_cache_order) Get(k string) (GetOrderResponse, error) {
	if (l.CacheMap[k] == Order_timestamp_pair{}) {
		return GetOrderResponse{}, errors.New("Order not in CacheMap")
	}
	return l.CacheMap[k].OrderResponse, nil
}

func (l *Lru_cache_order) reduce() {
	if len(l.CacheMap) == 0 {
		return
	}

	var minKey string
	var minTime time.Time
	firstIteration := true

	for key, pair := range l.CacheMap {
		if firstIteration {
			minKey = key
			minTime = pair.TimeStamp
			firstIteration = false
			continue
		}

		if pair.TimeStamp.Before(minTime) {
			minKey = key
			minTime = pair.TimeStamp
		}
	}

	delete(l.CacheMap, minKey)
}

type Order_timestamp_pair struct {
	OrderResponse GetOrderResponse
	TimeStamp     time.Time
}
