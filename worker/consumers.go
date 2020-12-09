package worker

import (
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

type VideoTranscodingParams struct {
	MediaUUID uuid.UUID
}

func videoTranscodingQueueConsumer(w *Worker, msgs <-chan amqp.Delivery) {
	for d := range msgs {
		w.Logger.Info("Received a message", zap.ByteString("body", d.Body))
		time.Sleep(3 * time.Second)
		w.Logger.Info("Done")
	}
}
