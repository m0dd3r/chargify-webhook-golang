package chargify_webhook_test

import (
	"errors"
	"fmt"
	"testing"

	"gitlab.com/batchblue-team/chargify-webhook-golang/chargify_webhook"
)

type verifier func(interface{}) error
type testInput struct {
	Payload     map[string]interface{}
	MessageType chargify_webhook.MessageType
	Verifier    verifier
}

var inputs []testInput = []testInput{
	testInput{
		map[string]interface{}{"chargify": "testing"},
		chargify_webhook.TEST,
		func(i interface{}) error {
			test := i.(chargify_webhook.Test)
			if test.Chargify != "testing" {
				return errors.New(fmt.Sprintf("Failed to properly populate TEST: %v", test))
			}
			return nil
		}},
	testInput{
		map[string]interface{}{"site": map[string]interface{}{"id": 5, "subdomain": "chargify"}},
		chargify_webhook.SUBSCRIPTION_STATE_CHANGE,
		func(i interface{}) error {
			ssc := i.(chargify_webhook.SubscriptionStateChange)
			if ssc.Site.Id != 5 || ssc.Site.Subdomain != "chargify" {
				return errors.New(fmt.Sprintf("Failed to properly populate SSC: %v", ssc))
			}
			return nil
		}},
}

func TestCreateTMessage(t *testing.T) {
	for _, input := range inputs {
		input := input
		msg, err := chargify_webhook.CreateMessage(input.MessageType, input.Payload)
		if err != nil {
			t.Errorf("Failed to create %s: %s", input.MessageType, err)
		}
		t.Logf("Verifying %s, %v", input.MessageType, msg)
		err = input.Verifier(msg)
		if err != nil {
			t.Error(err)
		}
	}
}

//func TestCreateTMessage(t *testing.T) {
//	payload := map[string]interface{}{"chargify": "testing"}
//	testI, err := chargify_webhook.CreateMessage(chargify_webhook.TEST, payload)
//	if err != nil {
//		t.Error("Failed to create TEST: ", err)
//	}
//	test, ok := testI.(chargify_webhook.Test)
//	if !ok {
//		t.Error("Failed to assert TEST: ", test)
//	}
//
//	if test.Chargify != "testing" {
//		t.Error("Failed to properly populate TEST: ", test.Chargify)
//	}
//}
