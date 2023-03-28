package kafka

import (
	"testing"

	"github.com/r1cebucket/gopkg/kafka"
	"github.com/r1cebucket/gopkg/log"
)

func TestConsumer(t *testing.T) {
	handlers := map[string]func([]byte) error{"test": func(msg []byte) error {
		log.Info().Msg(string(msg))
		return nil
	}}
	c := kafka.NewConsumer("test_group", []string{"test"})
	kafka.Consume(c, handlers)
}
