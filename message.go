package chargify_webhook

import "encoding/json"

const (
	TEST                      EventName = "test"
	SUBSCRIPTION_STATE_CHANGE EventName = "subscription_state_change"
	SIGNUP_SUCCESS            EventName = "signup_success"
)

type Test struct {
	Chargify string
}

type Payload struct {
	EventId      int
	Subscription Subscription
	Site         Site
}

type SubscriptionStateChange struct {
	Payload
}

type SignupSuccess struct {
	Payload
}

func NewTest(payload PayloadMap) (Message, error) {
	t := Test{}
	err := newMessage(&t, payload)
	if err != nil {
		return t, err
	}
	return t, nil
}

func NewSubscriptionStateChange(payload PayloadMap) (Message, error) {
	ssc := SubscriptionStateChange{}
	err := newMessage(&ssc, payload)
	if err != nil {
		return ssc, err
	}
	return ssc, nil
}

func NewSignupSuccess(payload PayloadMap) (Message, error) {
	ssc := SignupSuccess{}
	err := newMessage(&ssc, payload)
	if err != nil {
		return ssc, err
	}
	return ssc, nil
}

func newMessage(t interface{}, p PayloadMap) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, t)
	if err != nil {
		return err
	}
	return nil
}
