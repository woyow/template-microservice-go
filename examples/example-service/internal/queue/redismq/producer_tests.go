package redismq

import (
	"encoding/json"
)

/*
	Example publish message in redismq queue
*/

func (r *RedisMQ) NewAccount() {

	// Test entity
	type test struct{
		ID 		string 		`json:"id"`
		Test 	string 		`json:"test"`
	}

	newTest := &test{
		ID: "111-222-333-444",
		Test: "test-message",
	}

	// Open queue
	tests, err := r.connection.OpenQueue("tests")
	if err != nil {
		panic(err)
	}

	delivery, err := json.Marshal(&newTest)
	if err != nil {
		r.log.Error(err)
	}

	// Publish message to queue
	if err := tests.Publish(string(delivery)); err != nil {
		r.log.Printf("failed to publish: %s", err)
	}
}