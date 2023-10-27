package mq

import (
	"context"
	"log"

	"github.com/go-redis/redis"
	"github.com/rmukhamet/core_test_task/internal/apperrors"
	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
)

type MessageQueue struct {
	channel string
	client  *redis.Client
}

func New(cfg *config.GatewayConfig) *MessageQueue {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS.Addr,
		Password: cfg.REDIS.Password,
		DB:       cfg.REDIS.DB,
	})

	return &MessageQueue{
		channel: cfg.REDIS.Channel,
		client:  redisClient,
	}
}

func (mq *MessageQueue) Ping(ctx context.Context) error {
	pong, err := mq.client.Ping().Result()
	log.Print(pong, err)

	return err
}
func (mq *MessageQueue) Publish(ctx context.Context, v interface{}) error {
	var data interface{}
	switch v {
	case v.(model.Retailer):
		data = NewRetailer(v.(model.Retailer))
	default:
		return apperrors.ErrorUnknownDataToQueue
	}

	return mq.client.Publish(mq.channel, data).Err()
}

func (mq *MessageQueue) Subscribe(ctx context.Context, v interface{}) (<-chan Retailer, error) {
	ch := make(chan Retailer)
	return ch, nil
}

func (mq *MessageQueue) Close(ctx context.Context) error {
	// close connection
	// close channel
	return nil
}
