package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Config struct{} // nolint unused

type Kafka struct { // nolint unused
	*Config
	client *kafka.Conn
	writer *kafka.Writer
	reader *kafka.Reader
}

func (mq *Kafka) Init(ctx context.Context) error {
	//var err error

	// to produce messages
	topic := "shortlink"
	partition := 0

	//if mq.client, err = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition); err != nil {
	//	return err
	//}

	mq.writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"192.168.99.1:9092"},
		Topic:   topic,
		//Dialer:            nil,
		//Balancer:          &kafka.LeastBytes{},
		//MaxAttempts:       0,
		//QueueCapacity:     0,
		//BatchSize:         0,
		//BatchBytes:        0,
		//BatchTimeout:      0,
		//ReadTimeout:       0,
		//WriteTimeout:      0,
		//RebalanceInterval: 0,
		//IdleConnTimeout:   0,
		//RequiredAcks:      0,
		//Async:             false,
		//CompressionCodec:  nil,
		//Logger:            nil,
		//ErrorLogger:       nil,
	})

	mq.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:                []string{"192.168.99.1:9092"},
		GroupID:                "",
		Topic:                  topic,
		Partition:              partition,
		Dialer:                 nil,
		QueueCapacity:          0,
		MinBytes:               10e3, // 10KB
		MaxBytes:               10e6, // 10MB
		MaxWait:                0,
		ReadLagInterval:        0,
		GroupBalancers:         nil,
		HeartbeatInterval:      0,
		CommitInterval:         0,
		PartitionWatchInterval: 0,
		WatchPartitionChanges:  false,
		SessionTimeout:         0,
		RebalanceTimeout:       0,
		JoinGroupBackoff:       0,
		RetentionTime:          0,
		StartOffset:            0,
		ReadBackoffMin:         0,
		ReadBackoffMax:         0,
		Logger:                 nil,
		ErrorLogger:            nil,
		IsolationLevel:         0,
		MaxAttempts:            0,
	})

	//if err = mq.client.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
	//	return err
	//}

	return nil
}

func (mq *Kafka) Close() error {
	var err error

	if mq.client != nil {
		err = mq.client.Close()
	}

	if mq.writer != nil {
		err = mq.writer.Close()
	}

	return err
}

func (mq *Kafka) Send(message []byte) error {
	err := mq.writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte("TEST"),
			Value: message,
		},
	)

	return err
}

func (mq *Kafka) Subscribe(message chan []byte) error {
	for {
		msg, err := mq.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		message <- msg.Value
	}

	return nil
}
