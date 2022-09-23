package redismq

import (
	"github.com/woyow/example-module/internal/storage"

	"github.com/adjust/rmq/v5"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"time"
)

/*
	Example consumer on message in redismq queue
*/

type ConsumerTest struct {
	name   string
	count  int
	before time.Time
	log  *logrus.Logger
	storage *storage.Storage
}

func NewConsumerTest(name string, count int, logger *logrus.Logger, storage *storage.Storage) *ConsumerTest {
	return &ConsumerTest{
		name: name,
		count: count,
		before: time.Now(),
		log: logger,
		storage: storage,
	}
}

func (c *ConsumerTest) Consume(delivery rmq.Delivery) {
	payload := delivery.Payload()
	c.log.Infof("start consume %s", payload)
	time.Sleep(consumeDuration)

	c.count++
	if c.count%reportBatchSize == 0 {
		duration := time.Now().Sub(c.before)
		c.before = time.Now()
		perSecond := time.Second / (duration / reportBatchSize)
		c.log.Printf("%s consumed %d %s %d", c.name, c.count, payload, perSecond)
	}

	if c.count%reportBatchSize > 0 {
		if err := delivery.Ack(); err != nil {
			c.log.Debugf("failed to ack %s: %s", payload, err)
		} else {
			c.log.Debugf("acked %s", payload)
		}
	} else { // reject one per batch
		if err := delivery.Reject(); err != nil {
			c.log.Debugf("failed to reject %s: %s", payload, err)
		} else {
			c.log.Debugf("rejected %s", payload)
		}
	}

	// Test entity
	type test struct{
		ID 		string 		`json:"id"`
		Test 	string 		`json:"test"`
	}
	var newTest test

	err := json.Unmarshal([]byte(payload), &newTest)
	if err != nil {
		c.log.Println(err)
	}

	println(newTest.ID, newTest.Test)
}
