package kafka

import (
	"context"
	"github.com/dtusupbaev/goProject/internal/message_broker"
	lru "github.com/hashicorp/golang-lru"
)

type Broker struct {
	brokers  []string
	clientID string

	cacheBroker message_broker.CacheBroker
	cache       *lru.TwoQueueCache
}

func NewBroker(brokers []string, cache *lru.TwoQueueCache, clientID string) message_broker.MessageBroker {
	return &Broker{brokers: brokers, cache: cache, clientID: clientID}
}

func (b *Broker) Connect(ctx context.Context) error {
	brokers := []message_broker.BrokerWithClient{b.Cache()}
	for _, broker := range brokers {
		if err := broker.Connect(ctx, b.brokers); err != nil {
			return err
		}
	}
	return nil
}

func (b *Broker) Close() error {
	brokers := []message_broker.BrokerWithClient{b.Cache()}
	for _, broker := range brokers {
		if err := broker.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Broker) Cache() message_broker.CacheBroker {
	if b.cacheBroker == nil {
		b.cacheBroker = NewCacheBroker(b.cache, b.clientID)
	}

	return b.cacheBroker
}
