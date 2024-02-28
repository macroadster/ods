package backend

import (
  "context"
  "github.com/segmentio/kafka-go"
  "github.com/segmentio/kafka-go/sasl/plain"
	"ds/collector/pkg/common"
  "os"
)
// the topic and broker address are initialized as constants
const (
	topic = "message-log"
)

var user = os.Getenv("KAFKA_USER")
var password = os.Getenv("KAFKA_PASSWORD")
var brokerAddress = os.Getenv("KAFKA_BROKER")

func Produce(ctx context.Context, event common.Event) {

  mechanism := plain.Mechanism{Username:user, Password:password}

  sharedTransport := &kafka.Transport{
    SASL: mechanism,
  }
  // intialize the writer with the broker addresses, and the topic
  w := kafka.Writer{
    Addr:       kafka.TCP(brokerAddress),
    Topic:      topic,
    Balancer:   &kafka.Hash{},
    Transport:  sharedTransport,
  }

  err := w.WriteMessages(ctx, kafka.Message{
    Key: []byte("event"),
    // create an arbitrary message payload for the value
    Value: []byte(event.Data),
  })
  if err != nil {
    panic("could not write message " + err.Error())
  }
  w.Close()
}
