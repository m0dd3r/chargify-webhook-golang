package chargify_webhook

import "time"

type Subscription struct {
	Id             int       `form:"id"`
	State          string    `form:"state"`
	TrialStartedAt time.Time `form:"trial_started_at"`
	Customer       Customer  `form:"customer"`
}

type Customer struct {
	FirstName string `form:"fisrt_name"`
	LastName  string `form:"last_name"`
}

type Site struct {
	Id        int
	Subdomain string
}
