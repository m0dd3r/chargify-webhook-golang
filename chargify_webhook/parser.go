package chargify_webhook

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type ChargifyWebhook struct {
	Id      int
	Event   string
	Payload map[string]interface{}
}

const (
	ID      = "id"
	EVENT   = "event"
	PAYLOAD = "payload"
	PATTERN = "\\[([^\\]]+)]"
)

var (
	matcher *regexp.Regexp
	err     error
)

func init() {
	matcher = regexp.MustCompile(PATTERN)
}

func ParseChargifyWebhook(body string) (ChargifyWebhook, error) {
	var (
		err error
		w   ChargifyWebhook
	)
	pairs, err := url.ParseQuery(body)
	if err != nil {
		return w, err
	}

	return parse(pairs)
}

func parse(pairs url.Values) (ChargifyWebhook, error) {
	var (
		genericMap map[string]interface{}
		w          ChargifyWebhook
	)
	genericMap = make(map[string]interface{})
	w = ChargifyWebhook{}

	for k, _ := range pairs {
		if strings.HasPrefix(k, PAYLOAD) {
			var levels []string
			matches := matcher.FindAllStringSubmatch(k, -1)
			levels = make([]string, len(matches))
			for i, l := range matches {
				levels[i] = l[1]
			}
			val := pairs.Get(k)
			debugf("LEVELS: %v, value: %v\n", levels, val)
			buildGenericMap(genericMap, val, levels)
		}
	}
	w.Id, err = strconv.Atoi(pairs.Get(ID))
	if err != nil {
		return w, err
	}
	w.Event = pairs.Get(EVENT)
	w.Payload = genericMap
	return w, nil

}

func buildGenericMap(result map[string]interface{}, val interface{}, keys []string) {
	key := keys[0]
	if len(keys) < 1 {
		return
	} else if len(keys) == 1 {
		result[key] = val
	} else {
		key := keys[0]
		result[key] = make(map[string]interface{})
		buildGenericMap(result[key].(map[string]interface{}), val, keys[1:])
	}
}
