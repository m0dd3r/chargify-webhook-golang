package chargify_webhook_test

import (
	"testing"

	"github.com/m0dd3r/chargify-webhook-golang/chargify_webhook"
)

func TestUnmarshalText(t *testing.T) {
	//body := "id=123456&event=test&payload[chargify]=testing"
	body := "id=188397125&event=subscription_state_change&payload[subscription][activated_at]=2012-09-09%2011%3A38%3A33%20-0400&payload[subscription][balance_in_cents]=9900&payload[subscription][cancel_at_end_of_period]=false&payload[subscription][canceled_at]=&payload[subscription][cancellation_message]=&payload[subscription][coupon_code]=&payload[subscription][created_at]=2012-09-09%2011%3A38%3A32%20-0400&payload[subscription][credit_card][billing_address]=987%20Commerce%20St&payload[subscription][credit_card][billing_address_2]=Suite%20789&payload[subscription][credit_card][billing_city]=Greenberg&payload[subscription][credit_card][billing_country]=US&payload[subscription][credit_card][billing_state]=NC&payload[subscription][credit_card][billing_zip]=67890&payload[subscription][credit_card][card_type]=visa&payload[subscription][credit_card][current_vault]=bogus&payload[subscription][credit_card][customer_id]=0&payload[subscription][credit_card][expiration_month]=4&payload[subscription][credit_card][expiration_year]=2016&payload[subscription][credit_card][first_name]=Jane&payload[subscription][credit_card][id]=0&payload[subscription][credit_card][last_name]=Doe&payload[subscription][credit_card][masked_card_number]=XXXX-XXXX-XXXX-1111&payload[subscription][credit_card][vault_token]=1&payload[subscription][credit_card][customer_vault_token]=&payload[subscription][current_period_ends_at]=2012-10-09%2011%3A49%3A43%20-0400&payload[subscription][current_period_started_at]=2012-09-09%2011%3A49%3A43%20-0400&payload[subscription][customer][address]=123%20Main%20St&payload[subscription][customer][address_2]=Apt%20123&payload[subscription][customer][city]=Pleasantville&payload[subscription][customer][country]=US&payload[subscription][customer][created_at]=2012-09-09%2011%3A38%3A32%20-0400&payload[subscription][customer][email]=john%40example.com&payload[subscription][customer][first_name]=John&payload[subscription][customer][id]=0&payload[subscription][customer][last_name]=Doe&payload[subscription][customer][organization]=Acme%2C%20Inc.&payload[subscription][customer][phone]=555-555-1234&payload[subscription][customer][reference]=johndoe&payload[subscription][customer][state]=NC&payload[subscription][customer][updated_at]=2012-09-09%2011%3A38%3A32%20-0400&payload[subscription][customer][zip]=12345&payload[subscription][delayed_cancel_at]=&payload[subscription][expires_at]=&payload[subscription][id]=0&payload[subscription][next_assessment_at]=2012-10-09%2011%3A49%3A43%20-0400&payload[subscription][previous_state]=active&payload[subscription][payment_type]=credit_card&payload[subscription][product][accounting_code]=pro1234&payload[subscription][product][archived_at]=&payload[subscription][product][created_at]=2012-09-06%2010%3A09%3A35%20-0400&payload[subscription][product][description]=Vel%20soluta%20nihil%20qui%20accusamus%20quidem.&payload[subscription][product][expiration_interval]=&payload[subscription][product][expiration_interval_unit]=never&payload[subscription][product][handle]=handle_6a9273b8a&payload[subscription][product][id]=0&payload[subscription][product][initial_charge_in_cents]=&payload[subscription][product][interval]=1&payload[subscription][product][interval_unit]=month&payload[subscription][product][name]=Pro&payload[subscription][product][price_in_cents]=9900&payload[subscription][product][product_family][accounting_code]=aopf1234&payload[subscription][product][product_family][description]=Lorem%20ipsum%20dolor%20sit%20amet.&payload[subscription][product][product_family][handle]=acme-online&payload[subscription][product][product_family][id]=0&payload[subscription][product][product_family][name]=Acme%20Online&payload[subscription][product][request_credit_card]=true&payload[subscription][product][require_credit_card]=true&payload[subscription][product][return_params]=&payload[subscription][product][return_url]=&payload[subscription][product][trial_interval]=&payload[subscription][product][trial_interval_unit]=month&payload[subscription][product][trial_price_in_cents]=&payload[subscription][product][update_return_url]=&payload[subscription][product][updated_at]=2012-09-09%2011%3A36%3A53%20-0400&payload[subscription][signup_payment_id]=30&payload[subscription][signup_revenue]=99.00&payload[subscription][state]=active&payload[subscription][total_revenue_in_cents]=4200&payload[subscription][trial_ended_at]=&payload[subscription][trial_started_at]=&payload[subscription][updated_at]=2012-09-09%2011%3A49%3A44%20-0400&payload[site][id]=44871&payload[site][subdomain]=mashapp2"
	var (
		w   chargify_webhook.ChargifyWebhook
		err error
	)
	w, err = chargify_webhook.ParseChargifyWebhook(body)
	if err != nil {
		t.Error("Form could not be decoded: ", err)
	}
	if w.Id != 188397125 {
		t.Error("Failed to unmarshal id, got: ", w.Id)
	}
	if w.Event != "subscription_state_change" {
		t.Error("Failed to unmarshal event, got: ", w.Event)
	}

	s := w.Payload["subscription"].(map[string]interface{})
	if s["balance_in_cents"] != "9900" {
		t.Error("Failed to unmarshal payload[subscription][balance_in_cents], got: ", s["balance_in_cents"])
	}

	if s["activated_at"] != "2012-09-09 11:38:33 -0400" {
		t.Error("Failed to unmarshal payload[subscription][activated_at], got: ", s["activated_at"])
	}

	if _, ok := s["canceled_at"]; ok {
		t.Error("Failed to unmarshal payload[subscription][canceled_at], should ignore empty value, got: ", s["canceled_at"])
	}
}

func TestParseAndFactory(t *testing.T) {
	body := "id=188397125&event=subscription_state_change&payload[subscription][activated_at]=2012-09-09%2011%3A38%3A33%20-0400&payload[subscription][balance_in_cents]=9900&payload[subscription][cancel_at_end_of_period]=false&payload[subscription][canceled_at]=&payload[subscription][cancellation_message]=&payload[subscription][coupon_code]=&payload[subscription][created_at]=2012-09-09%2011%3A38%3A32%20-0400&payload[subscription][credit_card][billing_address]=987%20Commerce%20St&payload[subscription][credit_card][billing_address_2]=Suite%20789&payload[subscription][credit_card][billing_city]=Greenberg&payload[subscription][credit_card][billing_country]=US&payload[subscription][credit_card][billing_state]=NC&payload[subscription][credit_card][billing_zip]=67890&payload[subscription][credit_card][card_type]=visa&payload[subscription][credit_card][current_vault]=bogus&payload[subscription][credit_card][customer_id]=0&payload[subscription][credit_card][expiration_month]=4&payload[subscription][credit_card][expiration_year]=2016&payload[subscription][credit_card][first_name]=Jane&payload[subscription][credit_card][id]=0&payload[subscription][credit_card][last_name]=Doe&payload[subscription][credit_card][masked_card_number]=XXXX-XXXX-XXXX-1111&payload[subscription][credit_card][vault_token]=1&payload[subscription][credit_card][customer_vault_token]=&payload[subscription][current_period_ends_at]=2012-10-09%2011%3A49%3A43%20-0400&payload[subscription][current_period_started_at]=2012-09-09%2011%3A49%3A43%20-0400&payload[subscription][customer][address]=123%20Main%20St&payload[subscription][customer][address_2]=Apt%20123&payload[subscription][customer][city]=Pleasantville&payload[subscription][customer][country]=US&payload[subscription][customer][created_at]=2012-09-09%2011%3A38%3A32%20-0400&payload[subscription][customer][email]=john%40example.com&payload[subscription][customer][first_name]=John&payload[subscription][customer][id]=0&payload[subscription][customer][last_name]=Doe&payload[subscription][customer][organization]=Acme%2C%20Inc.&payload[subscription][customer][phone]=555-555-1234&payload[subscription][customer][reference]=johndoe&payload[subscription][customer][state]=NC&payload[subscription][customer][updated_at]=2012-09-09%2011%3A38%3A32%20-0400&payload[subscription][customer][zip]=12345&payload[subscription][delayed_cancel_at]=&payload[subscription][expires_at]=&payload[subscription][id]=0&payload[subscription][next_assessment_at]=2012-10-09%2011%3A49%3A43%20-0400&payload[subscription][previous_state]=active&payload[subscription][payment_type]=credit_card&payload[subscription][product][accounting_code]=pro1234&payload[subscription][product][archived_at]=&payload[subscription][product][created_at]=2012-09-06%2010%3A09%3A35%20-0400&payload[subscription][product][description]=Vel%20soluta%20nihil%20qui%20accusamus%20quidem.&payload[subscription][product][expiration_interval]=&payload[subscription][product][expiration_interval_unit]=never&payload[subscription][product][handle]=handle_6a9273b8a&payload[subscription][product][id]=0&payload[subscription][product][initial_charge_in_cents]=&payload[subscription][product][interval]=1&payload[subscription][product][interval_unit]=month&payload[subscription][product][name]=Pro&payload[subscription][product][price_in_cents]=9900&payload[subscription][product][product_family][accounting_code]=aopf1234&payload[subscription][product][product_family][description]=Lorem%20ipsum%20dolor%20sit%20amet.&payload[subscription][product][product_family][handle]=acme-online&payload[subscription][product][product_family][id]=0&payload[subscription][product][product_family][name]=Acme%20Online&payload[subscription][product][request_credit_card]=true&payload[subscription][product][require_credit_card]=true&payload[subscription][product][return_params]=&payload[subscription][product][return_url]=&payload[subscription][product][trial_interval]=&payload[subscription][product][trial_interval_unit]=month&payload[subscription][product][trial_price_in_cents]=&payload[subscription][product][update_return_url]=&payload[subscription][product][updated_at]=2012-09-09%2011%3A36%3A53%20-0400&payload[subscription][signup_payment_id]=30&payload[subscription][signup_revenue]=99.00&payload[subscription][state]=active&payload[subscription][total_revenue_in_cents]=4200&payload[subscription][trial_ended_at]=&payload[subscription][trial_started_at]=&payload[subscription][updated_at]=2012-09-09%2011%3A49%3A44%20-0400&payload[site][id]=44871&payload[site][subdomain]=mashapp2"
	var (
		w   chargify_webhook.ChargifyWebhook
		err error
	)
	w, err = chargify_webhook.ParseChargifyWebhook(body)
	if err != nil {
		t.Error("Form could not be decoded: ", err)
	}
	_, err = chargify_webhook.CreateMessage(chargify_webhook.MessageType(w.Event), w.Payload)
	if err != nil {
		t.Errorf("Failed to create %s", err)
	}
}
