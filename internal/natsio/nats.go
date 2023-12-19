package natsio

import (
	"context"
	"errors"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"time"
)

func Init(ctx context.Context, js jetstream.JetStream) error {
	_, err := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket:       "projection-offset-store",
		Description:  "",
		MaxValueSize: 100,
		History:      1,
		MaxBytes:     100,
		Storage:      jetstream.FileStorage,
	})
	if err != nil {
		return err
	}

	_, err = js.CreateStream(ctx, jetstream.StreamConfig{
		Name:                 "event_store",
		Description:          "Main event stream",
		Subjects:             []string{"ingest.events"},
		Retention:            jetstream.LimitsPolicy,
		MaxConsumers:         -1,
		MaxMsgs:              922337203685477,
		MaxBytes:             -1,
		Discard:              jetstream.DiscardOld,
		DiscardNewPerSubject: false,
		MaxAge:               time.Hour * 99999,
		MaxMsgsPerSubject:    -1,
		MaxMsgSize:           -1,
		Replicas:             1,
		Duplicates:           time.Minute,
	})

	if errors.Is(err, jetstream.ErrStreamNameAlreadyInUse) {
		log.Println("stream exists")
		return nil
	}

	return err
}
