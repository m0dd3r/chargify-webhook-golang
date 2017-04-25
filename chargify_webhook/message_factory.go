package chargify_webhook

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Message interface{}
type MessageType string
type MessageFactory func(PayloadMap) (Message, error)

var (
	messageFactories map[MessageType]MessageFactory
)

func init() {
	matcher = regexp.MustCompile(PATTERN)
	messageFactories = make(map[MessageType]MessageFactory)
	RegisterMessageFactory(TEST, NewTest)
	RegisterMessageFactory(SUBSCRIPTION_STATE_CHANGE, NewSubscriptionStateChange)
}

func RegisterMessageFactory(t MessageType, factory MessageFactory) {
	if factory == nil {
		log.Panicf("Message factory %s does not exist.", t)
	}
	_, registered := messageFactories[t]
	if registered {
		log.Printf("Message factory %s already registered. Ignoring.", t)
	}
	messageFactories[t] = factory
}

func CreateMessage(t MessageType, p PayloadMap) (Message, error) {
	factory, ok := messageFactories[t]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableMessageTypes := make([]string, len(messageFactories))
		for k, _ := range messageFactories {
			availableMessageTypes = append(availableMessageTypes, string(k))
		}
		return nil, errors.New(fmt.Sprintf("Invalid Message type. Must be one of: %s", strings.Join(availableMessageTypes, ", ")))
	}
	return factory(p)
}
