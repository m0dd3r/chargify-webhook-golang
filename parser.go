package chargify_webhook

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type EventName string
type PayloadMap map[string]interface{}

type ChargifyWebhook struct {
	Id      int
	Event   EventName
	Payload PayloadMap
}

type Parser struct {
	matcher *regexp.Regexp
}

const (
	ID      = "id"
	EVENT   = "event"
	PAYLOAD = "payload"
	PATTERN = "\\[([^\\]]+)]"
)

func NewParser() Parser {
	return Parser{
		matcher: regexp.MustCompile(PATTERN),
	}
}

func ParseChargifyWebhook(body string) (ChargifyWebhook, error) {
	w := ChargifyWebhook{}
	err := NewParser().ParseChargifyWebhook(body, &w)
	if err != nil {
		return w, err
	}
	return w, nil
}

func (parser Parser) ParseChargifyWebhook(body string, w *ChargifyWebhook) error {
	var (
		err error
	)
	pairs, err := url.ParseQuery(body)
	if err != nil {
		return err
	}

	return parser.parseKeyValuePairs(pairs, w)
}

func (parser Parser) parseKeyValuePairs(pairs url.Values, w *ChargifyWebhook) error {
	var err error

	payloadMap := make(map[string]interface{})

	for k, _ := range pairs {
		if strings.HasPrefix(k, PAYLOAD) {
			matches := parser.matcher.FindAllStringSubmatch(k, -1)
			levels := make([]string, len(matches))
			for i, l := range matches {
				levels[i] = l[1]
			}

			val := pairs.Get(k)
			if val != "" {
				buildNestedMap(payloadMap, val, levels)
			}
		}
	}

	w.Id, err = strconv.Atoi(pairs.Get(ID))
	if err != nil {
		return err
	}
	w.Event = EventName(pairs.Get(EVENT))
	w.Payload = payloadMap
	return nil
}

func buildNestedMap(result map[string]interface{}, val string, keys []string) {
	if len(keys) < 1 {
		return
	}
	key := keys[0]
	if len(keys) == 1 {
		result[key] = val
	} else {
		if _, ok := result[key]; !ok {
			result[key] = make(map[string]interface{})
		}
		buildNestedMap(result[key].(map[string]interface{}), val, keys[1:])
	}
}
