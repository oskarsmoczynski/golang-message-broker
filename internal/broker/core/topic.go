package core

import (
	"errors"
	"fmt"

	"github.com/oskarsmoczynski/golang-message-broker/internal/broker/queue"
	"github.com/oskarsmoczynski/golang-message-broker/internal/broker/utils"
)

var ErrAlreadyExists = errors.New("already exists")

type Topic struct {
	Queue *queue.Queue
}

type TopicManager struct {
	topics map[string]*Topic
}

func NewTopicManager() *TopicManager {
	return &TopicManager{
		topics: make(map[string]*Topic, 0),
	}
}

func (tm *TopicManager) GetTopicByName(name string) (*Topic, bool) {
	if !utils.IsValidTopicName(name) {
		return nil, false
	}
	topic, ok := tm.topics[name]
	return topic, ok
}

func (tm *TopicManager) CreateTopic(name string) error {
	if _, exists := tm.GetTopicByName(name); exists {
		return fmt.Errorf("failed to create topic %s: %w", name, ErrAlreadyExists)
	}
	tm.topics[name] = &Topic{
		Queue: queue.NewQueue(),
	}
	return nil
}
