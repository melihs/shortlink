/*
Message Queue
*/
package mq

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/mq/kafka"
	"github.com/batazor/shortlink/internal/pkg/mq/nats"
	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/mq/rabbit"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain/link"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// Use return implementation of MQ
func (mq *DataBus) Use(ctx context.Context, log logger.Logger) (MQ, error) {
	// Set configuration
	mq.setConfig()

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, mq)

	switch mq.typeMQ {
	case "kafka":
		mq.mq = &kafka.Kafka{}
	case "nats":
		mq.mq = &nats.NATS{}
	case "rabbitmq":
		mq.mq = &rabbit.RabbitMQ{
			Log: log,
		}
	default:
		mq.mq = &kafka.Kafka{}
	}

	if err := mq.mq.Init(ctx); err != nil {
		return nil, err
	}

	log.Info("run MQ", field.Fields{
		"mq": mq.typeMQ,
	})

	return mq.mq, nil
}

// setConfig - set configuration
func (mq *DataBus) setConfig() { // nolint unused
	viper.SetDefault("MQ_TYPE", "rabbitmq") // Select: kafka, rabbitmq, nats
	mq.typeMQ = viper.GetString("MQ_TYPE")
}

// Notify ...
func (mq *DataBus) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response {
	switch event {
	case api_type.METHOD_ADD:
		// TODO: send []byte
		msg := payload.(*link.Link) // nolint errcheck
		data, err := proto.Marshal(msg)
		if err != nil {
			return notify.Response{
				Name:    "RESPONSE_MQ_ADD",
				Payload: nil,
				Error:   err,
			}
		}

		err = mq.mq.Publish(ctx, query.Message{
			Key:     nil,
			Payload: data,
		})
		return notify.Response{
			Name:    "RESPONSE_MQ_ADD",
			Payload: nil,
			Error:   err,
		}
	case api_type.METHOD_GET:
		panic("implement me")
	case api_type.METHOD_LIST:
		panic("implement me")
	case api_type.METHOD_UPDATE:
		panic("implement me")
	case api_type.METHOD_DELETE:
		panic("implement me")
	}

	return notify.Response{}
}
