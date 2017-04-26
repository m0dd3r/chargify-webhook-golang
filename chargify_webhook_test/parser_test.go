package chargify_webhook_test

import (
	"fmt"
	"testing"

	"github.com/m0dd3r/chargify-webhook-golang/chargify_webhook"
)

func TestUnmarshalText(t *testing.T) {
	body := "id=123456&event=test&payload[chargify]=testing"
	var (
		w   chargify_webhook.ChargifyWebhook
		err error
	)
	w, err = chargify_webhook.ParseChargifyWebhook(body)
	if err != nil {
		t.Error("Form could not be decoded: ", err)
	}
	if w.Id != 123456 {
		t.Error("Failed to unmarshal id, got: ", w.Id)
	}
	if w.Event != "test" {
		t.Error("Failed to unmarshal event, got: ", w.Event)
	}
	fmt.Println("w ", w)
	if w.Payload["chargify"] != "testing" {
		t.Error("Failed to unmarshal payload[chargify], got: ", w.Payload["chargify"])
	}

}
