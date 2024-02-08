package main

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	must(err)

	defer nc.Close()

	js, err := jetstream.New(nc)
	must(err)

	ctx := context.Background()

	// TODO - Use bourgon errgroup

	log.Fatal(project(ctx, js))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func project(ctx context.Context, js jetstream.JetStream) error {
	kv, err := js.KeyValue(ctx, "projection-offset-store")
	if err != nil {
		return err
	}

	var offset uint64

	key := "projection-foo-offset"

	entry, err := kv.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, jetstream.ErrKeyNotFound) {
			return err
		}
	}

	if err == nil {
		offset = binary.BigEndian.Uint64(entry.Value())
	}

	log.Println("starting from offset ", offset)

	cons, err := js.OrderedConsumer(ctx, "event_store", jetstream.OrderedConsumerConfig{
		FilterSubjects:    nil,
		DeliverPolicy:     jetstream.DeliverByStartSequencePolicy,
		OptStartSeq:       offset,
		OptStartTime:      nil,
		ReplayPolicy:      jetstream.ReplayInstantPolicy,
		InactiveThreshold: 0,
		HeadersOnly:       false,
		MaxResetAttempts:  0,
	})
	if err != nil {
		return err
	}

	b := make([]byte, 8)

	cc, err := cons.Consume(func(msg jetstream.Msg) {
		offset++
		binary.BigEndian.PutUint64(b, offset)

		_, err = kv.Put(ctx, key, b)
		if err != nil {
			// TODO - Retries per event should be handled here
			log.Println(err)
			return
		}

		// TODO

		// TODO - grpc client to the runner eg. .net
		// Define the rpc contract here in this repo

		fmt.Println("Handling event: ", string(msg.Data()))

		// TODO - If message handling fails reset offset from kv
		// Is there a better strategy here in order to prevent duplicates ?
		// Give user a few configurable strategies for offset handling

		// TODO - Handle kill signals from ctx per message so we don't end up in funny state between
		// message handling and offset saving

		// But defer al of these things
	})

	defer cc.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}
}
