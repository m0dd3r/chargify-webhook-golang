# chargify-webhook-golang
Golang library to parse the request body of a Chargify Webhook. Inspired by https://github.com/prowave/chargify-webhook-java.

## Installation

`go get m0dd3r/chargify-webhook-golang`

### Usage

```golang 
import (
  cw "github.com/m0dd3r/chargify-webhook-golang"
)

body := "id=123456&event=test&payload[chargify]=testing"

wh, err = cw.ParseChargifyWebhook(body)
if err != nil {
  panic("Event Body could not be parsed: ", err)
}

msg, err = cw.CreateMessage(cw.EventName(wh.Event), wh.Payload)
if err != nil {
  panic("Failed to create message from webhook: ", err)
}

switch t := msg.(type) {
case cw.Test:
  fmt.Println("Test message: ", t)
case cw.SubscriptionStateChange:
  fmt.Println("SubscriptionStateChange message: ", t)
}
```
#### Provided

* `Test`
* `SubscriptionStateChange`

#### Backlog

* `SignupSuccess`
* `CustomerUpdate`
* `PaymentSuccess`
* `PaymentFailure`
* `SignupFailure`
* `RenewalSuccess`
* `RenewalFailure`
* `BillingDateChange`
* `SubscriptionProductChange`
* `ExpiringCard`
* `ComponentAllocationChange`
* `UpcomingRenewalNotice`
* `EndOfTrialNotice` - possible duplicate, or superceded by `TrialEndNotice`
* `TrialEndNotice`
* `UpgradeDowngradeSuccess`
* `UpgradeDowngradeFailure`
* `ExpirationDateChange`
* `SubscriptionCardUpdate`
* `MeteredUsage`
* `StatementClosed`
* `StatementSettled`
* `RefundSuccess`
* `RefundFailure`


## Other Considerations

> **NOTE**<br>
> The current implementation only supports a subset of the total webhook message types.
> PRs are more than welcome to add new message types and improve things.

