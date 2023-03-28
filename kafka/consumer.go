package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
)

// ref: https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/consumer_example/consumer_example.go

func NewConsumer(group string, topics []string) *kafka.Consumer {
	c, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers":  strings.Join(config.Kafka.Servers, ","),
			"group.id":           group,
			"session.timeout.ms": 6000,
			// Start reading from the first message of each assigned
			// partition if there are no previously committed offsets
			// for this group.
			"auto.offset.reset": "earliest",
			// Whether or not we store offsets automatically.
			"enable.auto.offset.store": false,
		})

	if err != nil {
		log.Panic().Msg("faile to create consumer: " + err.Error())
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		log.Panic().Msg("faile to subscribe topics: " + err.Error())
	}

	return c
}

type Handler interface {
	Handle([]byte) error
}

func Consume(c *kafka.Consumer, handlers map[string]func([]byte) error) {
	// A signal handler or similar could be used to set this to false to break the loop.
	defer c.Close()
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigchan:
			log.Info().Msg(fmt.Sprintf("Caught signal %v: terminating\n", sig))
			return
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				// Process the message received.
				log.Info().Msg(fmt.Sprintf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value)))
				err := handlers[*e.TopicPartition.Topic](e.Value)
				if err != nil {
					log.Err(err)
				}
				if e.Headers != nil {
					log.Info().Msg(fmt.Sprintf("%% Headers: %v\n", e.Headers))
				}

				// We can store the offsets of the messages manually or let
				// the library do it automatically based on the setting
				// enable.auto.offset.store. Once an offset is stored, the
				// library takes care of periodically committing it to the broker
				// if enable.auto.commit isn't set to false (the default is true).
				// By storing the offsets manually after completely processing
				// each message, we can ensure atleast once processing.
				_, err = c.StoreMessage(e)
				if err != nil {
					log.Err(err).Msg(fmt.Sprintf("%% Error storing offset after message %s:\n", e.TopicPartition))
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				log.Err(e).Msg("Kafka error")
				// log.Err(e).Msg(fmt.Sprintf("%% Error: %v: %v\n", e.Code(), e))
				if e.Code() == kafka.ErrAllBrokersDown {
					return
				}
			default:
				log.Info().Msg(fmt.Sprintf("Ignored %v\n", e))
			}
		}
	}
}
