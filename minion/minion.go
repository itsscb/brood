package minion

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/itsscb/brood/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Minion struct {
	overlordURL string
	topics      []string
	hive        string
	brood       pb.BroodClient
}

func New(overlordURL, hive string, topics []string) (minion *Minion, err error) {
	minion = &Minion{
		overlordURL: overlordURL,
		topics:      topics,
		hive:        hive,
	}

	conn, err := grpc.Dial(overlordURL, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		return nil, err
	}

	minion.brood = pb.NewBroodClient(conn)

	go minion.brood.Join(context.Background(), &pb.JoinBrood{
		Hive:   hive,
		Topics: topics,
	})

	return minion, nil
}

func (m *Minion) Join(ctx context.Context) error {
	client, err := m.brood.Join(ctx, &pb.JoinBrood{
		Hive:   m.hive,
		Topics: m.topics,
	})
	if err != nil {
		return err
	}

	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return io.EOF
		}
		if err != nil {
			return err
		}
		log.Printf("Spore received: %#v", resp)
	}
}

func (m *Minion) Publish(ctx context.Context, topic, data string) error {
	resp, err := m.brood.SendSpore(ctx, &pb.Spore{
		Id: uuid.New().String(),
		Timestamp: timestamppb.New(
			time.Now(),
		),
		Topic: topic,
		Data:  data,
	})
	if err != nil || !resp.Ack {
		if !resp.Ack {
			return errors.New("spore not acknowledged")
		}
		return err
	}

	return nil
}
