package mq

import (
	"context"
	"encoding/json"
	"fmt"
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

func New(cfg *config.REDIS) *MessageQueue {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &MessageQueue{
		channel: cfg.Channel,
		client:  redisClient,
	}
}

func (mq *MessageQueue) ping(ctx context.Context) error {
	_, err := mq.client.Ping().Result()
	if err != nil {
		return fmt.Errorf("error ping redis with error: %w", err)
	}

	return nil
}
func (mq *MessageQueue) Publish(ctx context.Context, v interface{}) error {
	var data interface{}
	switch v {
	case v.(model.Retailer):
		data = NewRetailer(v.(model.Retailer))
	default:
		return apperrors.ErrorUnknownDataToQueue
	}

	err := mq.client.Publish(mq.channel, data).Err()
	if err != nil {
		return fmt.Errorf("failed publish: %+v with error: %w", data, err)
	}
	// debug
	log.Printf("published: %+v", data)

	return err
}

func (mq *MessageQueue) Subscribe(ctx context.Context, v interface{}) (<-chan model.Retailer, error) {
	pubsub := mq.client.Subscribe(mq.channel)
	ch := make(chan model.Retailer)
	defer close(ch)
	defer pubsub.Close()

	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case msg := <-pubsub.Channel():
				var retailer Retailer

				err := json.Unmarshal([]byte(msg.Payload), &retailer)
				if err != nil {
					log.Printf("wrong message format: %s, error: %s\n", msg.Payload, err.Error())
					continue
				}

				log.Print(retailer)

				ch <- retailer.ToDTO()
			}
		}
	}()

	return ch, nil
}

func (mq *MessageQueue) Close(ctx context.Context) error {
	// close connection
	// close channel
	return nil
}
