package nats

import (
	"github.com/nats-io/nats.go"

	"encoding/json"
)

/*
	Example publish message in queue
*/

func (n *Nats) PublishCreateTest() {

	// Test entity version 1
	//type test struct {
	//	Version 	int 		`json:"version"` // version of message
	//	ID 			string		`json:"id"`
	//	Test 		string 		`json:"test"`
	//}

	// Test entity version 2
	type testV2 struct {
		Version 	int 		`json:"version"` // version of message
		ID 			string		`json:"id"`
		Test 		string 		`json:"test"`
		NewTest 	string 		`json:"new_test"` // New on version 2
	}

	//newTest := test{
	//	Version: 1,
	//	ID: "222-333-444-555",
	//	Test: "test message",
	//}

	newTestV2 := testV2{
		Version: 1,
		ID: "222-333-444-555",
		Test: "test message",
		NewTest: "oh new field on message version 2",
	}

	//createTestPayload, err := json.Marshal(&newTest)
	//if err != nil {
	//	n.log.Error(err)
	//	return
	//}

	createTestPayloadV2, err := json.Marshal(&newTestV2)
	if err != nil {
		n.log.Error(err)
		return
	}

	js, err := n.conn.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		n.log.Error(err)
	}

	js.PublishAsync("TESTS.create", createTestPayloadV2)
}
