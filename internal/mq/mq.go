package mq

import "context"

type MessageQueue struct {
}

func New() *MessageQueue {
	return &MessageQueue{}
}

func (mq *MessageQueue) Push(ctx context.Context, v interface{}) error {
	return nil
}

func (mq *MessageQueue) Pull(ctx context.Context, v interface{}) (<-chan retailer, error) {
	ch := make(chan retailer)
	return ch, nil
}

func (mq *MessageQueue) Close(ctx context.Context) error {
	// close connection
	// close channel
	return nil
}
