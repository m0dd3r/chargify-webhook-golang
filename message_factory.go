package chargify_webhook

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Message interface{}
type MessageFactory func(PayloadMap) (Message, error)

var (
	messageFactories map[EventName]MessageFactory
)

func init() {
	messageFactories = make(map[EventName]MessageFactory)
	RegisterMessageFactory(TEST, NewTest)
	RegisterMessageFactory(SUBSCRIPTION_STATE_CHANGE, NewSubscriptionStateChange)
}

func RegisterMessageFactory(t EventName, factory MessageFactory) {
	if factory == nil {
		log.Panicf("Message factory %s does not exist.", t)
	}
	_, registered := messageFactories[t]
	if registered {
		log.Printf("Message factory %s already registered. Ignoring.", t)
	}
	messageFactories[t] = factory
}

func CreateMessage(t EventName, p PayloadMap) (Message, error) {
	factory, ok := messageFactories[t]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available factories for logging.
		availableEventNames := make([]string, len(messageFactories))
		for k, _ := range messageFactories {
			availableEventNames = append(availableEventNames, string(k))
		}
		return nil, errors.New(fmt.Sprintf("Invalid Message type. Must be one of: %s", strings.Join(availableEventNames, ", ")))
	}
	return factory(p)
}
