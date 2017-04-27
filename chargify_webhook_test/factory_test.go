package chargify_webhook_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/m0dd3r/chargify-webhook-golang/chargify_webhook"
)

type verifier func(interface{}) error
type testInput struct {
	Payload     map[string]interface{}
	MessageType chargify_webhook.MessageType
	Verifier    verifier
}

const timeFormat = "2006-01-02 15:04:05 -0700"

var now = "2012-09-09 11:38:33 -0400"

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
		map[string]interface{}{
			"site": map[string]interface{}{
				"id":        "5",
				"subdomain": "chargify",
			},
			"subscription": map[string]interface{}{
				"id":               "55",
				"state":            "active",
				"trial_started_at": now,
				"customer": map[string]interface{}{
					"first_name": "bob",
					"last_name":  "dobbler",
				},
			},
		},
		chargify_webhook.SUBSCRIPTION_STATE_CHANGE,
		func(i interface{}) error {
			ssc := i.(chargify_webhook.SubscriptionStateChange)
			if ssc.Site.Id != 5 || ssc.Site.Subdomain != "chargify" {
				return errors.New(fmt.Sprintf("Failed to properly populate SSC.Site: %v", ssc))
			}

			if ssc.Subscription.Id != 55 ||
				ssc.Subscription.State != "active" {
				return errors.New(fmt.Sprintf("Failed to properly populate SSC.Subscription: %v", ssc.Subscription))
			}
			if ssc.Subscription.TrialStartedAt.Format(timeFormat) != now {
				return errors.New(fmt.Sprintf("Failed to properly parse times: %v, %v", ssc.Subscription.TrialStartedAt.Format(timeFormat), now))
			}

			if ssc.Subscription.Customer.FirstName != "bob" ||
				ssc.Subscription.Customer.LastName != "dobbler" {
				return errors.New(fmt.Sprintf("Failed to properly populate SSC.Subscription.Customer: %v", ssc.Subscription.Customer))
			}
			return nil
		},
	},
}

func TestCreateMessage(t *testing.T) {
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
