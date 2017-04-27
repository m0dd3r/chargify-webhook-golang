package test

import (
	"errors"
	"fmt"
	"testing"

	cw "github.com/m0dd3r/chargify-webhook-golang"
)

type verifier func(interface{}) error
type testInput struct {
	Payload     map[string]interface{}
	MessageType cw.MessageType
	Verifier    verifier
}

const (
	timeFormat = "2006-01-02 15:04:05 -0700"
	now        = "2012-09-09 11:38:33 -0400"
)

var inputs []testInput = []testInput{
	testInput{
		map[string]interface{}{"chargify": "testing"},
		cw.TEST,
		func(i interface{}) error {
			test := i.(cw.Test)
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
		cw.SUBSCRIPTION_STATE_CHANGE,
		func(i interface{}) error {
			ssc := i.(cw.SubscriptionStateChange)
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
		msg, err := cw.CreateMessage(input.MessageType, input.Payload)
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
