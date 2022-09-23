package nats

import (
	"encoding/json"
)

/*
	Example subscriber on message in queue
*/

func (n *Nats) CreateTestSubscriber(data *[]byte) error {

	// Test entity
	type test struct {
		Version 	int 		`json:"version"` // version of message
		ID 			string		`json:"id"`
		Test 		string 		`json:"test"`
		Field 		string
		IsActive 	bool
		NewTest 	string 		`json:"new_test"` // New on version 2
	}

	var newTest test

	json.Unmarshal(*data, &newTest)

	if newTest.Version == 1 {
		newTest.Field = ""
		newTest.IsActive = true
		newTest.NewTest = ""
	}

	if newTest.Version == 2 {
		newTest.Field = "something new"
		newTest.IsActive = false
	}

	// Example: call storage method
	//if err := n.storage.Psql.Test.CreateTestQueue(&newTest); err != nil {
	//	n.log.Errorln(err)
	//	return err
	//}

	n.log.Debug("create test by queue: ", newTest)
	return nil
}