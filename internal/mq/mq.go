package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
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

func (mq *MessageQueue) Publish(ctx context.Context, t model.Task) error {
	task := NewTask(t)
	err := mq.client.Publish(mq.channel, task).Err()
	if err != nil {
		return fmt.Errorf("failed publish: %+v with error: %w", t, err)
	}
	// debug
	log.Printf("published: %+v", t)

	return nil
}

func (mq *MessageQueue) Subscribe(ctx context.Context) (<-chan model.Task, error) {
	pubsub := mq.client.Subscribe(mq.channel)

	ch := make(chan model.Task)

	go func() {
		defer close(ch)
		defer pubsub.Close()

		for {
			select {
			case <-ctx.Done():

				break
			case msg := <-pubsub.Channel():
				if msg == nil {
					continue
				}

				var task Task

				log.Printf("DEBUG reseived msg: %+v\n", msg)

				err := json.Unmarshal([]byte(msg.Payload), &task) //todo new func
				if err != nil {
					log.Printf("wrong queue message format: %s, error: %s\n", msg.Payload, err.Error())
					continue
				}

				log.Print("DEBUG", task)

				ch <- task.ToDTO()
			}
		}
	}()

	return ch, nil
}

func (mq *MessageQueue) Close(ctx context.Context) error {
	return mq.client.Close()
}
