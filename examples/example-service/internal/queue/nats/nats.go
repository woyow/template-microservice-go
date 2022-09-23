package nats

import (
	"github.com/woyow/example-module/config"
	"github.com/woyow/example-module/internal/storage"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const (
	durableName = "YOUR-SERVICE-NAME"
)


type Nats struct {
	cfg *config.Nats
	conn *nats.Conn
	log *logrus.Logger
	storage *storage.Storage
}

func NewNats(cfg *config.Nats, storage *storage.Storage, logger *logrus.Logger) (*Nats, error) {

	natsProto := "nats://"
	natsURL := natsProto + cfg.Host + ":" + cfg.Port

	// Connect to a server
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	return &Nats{
		cfg: cfg,
		conn: nc,
		storage: storage,
		log: logger,
	}, nil
}

func (n *Nats) Run() {
	n.StartSubscribers()
	n.log.Debug("Nuts subscribers has been started")
}

func (n *Nats) StartSubscribers() {

	n.conn.QueueSubscribe("tests.create", "job_workers", func(m *nats.Msg) {
		n.CreateTestSubscriber(&m.Data)
	})

	js, err := n.conn.JetStream()
	if err != nil {
		n.log.Error(err)
		return
	}

	streamName := "TESTS"
	streamSubjects := "TESTS.*"

	stream, err := js.StreamInfo(streamName)
	if err != nil {
		n.log.Println(err)
	}

	if stream == nil {
		n.log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			n.log.Println(err)
		}
	}

	js.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable: durableName,
		DeliverPolicy: nats.DeliverAllPolicy,
		AckPolicy: nats.AckExplicitPolicy,
		ReplayPolicy: nats.ReplayInstantPolicy,
		Replicas: 1,
	})

	for {
		sub, err := js.PullSubscribe("tests.create", durableName)
		if err != nil {
			n.log.Println(err)
		}
		msgs, err := sub.Fetch(1)

		for _, msg := range msgs {
			err := msg.InProgress()
			if err != nil {
				n.log.Println("Unable to In Progress Msg", err)
			}
			n.log.Debug("Received a message: ", string(msg.Data))
			if err := n.CreateTestSubscriber(&msg.Data); err != nil {
				n.log.Println(err)
			}
			err = msg.Ack()
			if err != nil {
				n.log.Println("Unable to Ack Msg", err)
			} else {
				n.log.Infof("Msg from subject %s has been acknowledged", msg.Subject)
			}
		}
	}
}

func (n *Nats) Shutdown() {
	n.conn.Close()
}