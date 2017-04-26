package chargify_webhook

import "time"

type Subscription struct {
	Id                      int        `json:"id"`
	State                   string     `json:"state"`
	TrialStartedAt          time.Time  `json:"trial_started_at"`
	Customer                Customer   `json:"customer"`
	Product                 Product    `json:"product"`
	CreditCard              CreditCard `json:"credit_card"`
	TrialEndedAt            time.Time  `json:"trial_ended_at"`
	ActivatedAt             time.Time  `json:"activated_at"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
	ExpiresAt               time.Time  `json:"expires_at"`
	PreviousExpiresAt       time.Time  `json:"previous_expires_at"`
	BalanceInCents          int        `json:"balance_in_cents"`
	CurrentPeriodEndsAt     time.Time  `json:"current_period_ends_at"`
	NextAssessmentAt        time.Time  `json:"next_assessment_at"`
	CanceledAt              time.Time  `json:"canceled_at"`
	CancellationMessage     string     `json:"cancellation_message"`
	NextProductId           int        `json:"next_product_id"`
	CancelAtEndOfPeriod     bool       `json:"cancel_at_end_of_period"`
	PaymentCollectionMethod string     `json:"payment_collection_method"`
	SnapDay                 string     `json:"snap_day"`
	CancellationMethod      string     `json:"cancellation_method"`
	CurrentPeriodStartAt    time.Time  `json:"current_period_started_at"`
	PreviousState           string     `json:"previous_state"`
	SignupPaymentId         int        `json:"signup_payment_id"`
	SignupRevenue           float32    `json:"signup_revenue"`
	DelayedCancelAt         time.Time  `json:"delayed_cancel_at"`
	CouponCode              string     `json:"coupon_code"`
	TotalRevenueInCents     int        `json:"total_revenue_in_cents"`
	ProductPriceInCents     int        `json:"product_price_in_cents"`
	ProductVersionNumber    int        `json:"product_version_number"`
	PaymentType             string     `json:"payment_type"`
	ReferralCode            string     `json:"referral_code"`
	CouponUseCount          int        `json:"coupon_use_count"`
	CouponUsesAllowed       int        `json:"coupon_uses_allowed"`
}

type Customer struct {
	Id                         int       `json:"id"`
	FirstName                  string    `json:"first_name"`
	LastName                   string    `json:"last_name"`
	Organization               string    `json:"organization"`
	Email                      string    `json:"email"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
	Reference                  string    `json:"reference"`
	Address                    string    `json:"address"`
	Address2                   string    `json:"address_2"`
	City                       string    `json:"city"`
	State                      string    `json:"state"`
	Zip                        string    `json:"zip"`
	Country                    string    `json:"country"`
	Phone                      string    `json:"phone"`
	PortalInviteLastSentAt     time.Time `json:"portal_invite_last_sent_at"`
	PortalInviteLastAcceptedAt time.Time `json:"portal_invite_last_accepted_at"`
	Verified                   bool      `json:"verified"`
	PortalCustomerCreatedAt    time.Time `json:"portal_customer_created_at"`
	CcEmails                   string    `json:"cc_emails"`
	TaxExempt                  bool      `json:"tax_exempt"`
}

type Product struct {
	Id                      int                `json:"id"`
	Name                    string             `json:"name"`
	Handle                  string             `json:"handle"`
	Description             string             `json:"description"`
	AccountingCode          string             `json:"accounting_code"`
	RequestCreditCard       bool               `json:"request_credit_card"`
	ExpirationInterval      int                `json:"expiration_interval"`
	ExpirationIntervalUnit  string             `json:"expiration_interval_unit"`
	CreatedAt               time.Time          `json:"created_at"`
	UpdatedAt               time.Time          `json:"updated_at"`
	PriceInCents            int                `json:"price_in_cents"`
	Interval                int                `json:"interval"`
	IntervalUnit            string             `json:"interval_unit"`
	InitialChargeInCents    int                `json:"initial_charge_in_cents"`
	TrialPriceInCents       int                `json:"trial_price_in_cents"`
	TrialInterval           int                `json:"trial_interval"`
	TrialIntervalUnit       string             `json:"trial_interval_unit"`
	ArchivedAt              time.Time          `json:"archived_at"`
	RequireCreditCard       bool               `json:"require_credit_card"`
	ReturnParams            string             `json:"return_params"`
	Taxable                 bool               `json:"taxable"`
	UpdateReturnUrl         string             `json:"update_return_url"`
	InitialChargeAfterTrial bool               `json:"initial_charge_after_trial"`
	VersionNumber           int                `json:"version_number"`
	UpdateReturnParams      string             `json:"update_return_params"`
	ProductFamily           ProductFamily      `json:"product_family"`
	PublicSignupPages       []PublicSignupPage `json:"public_signup_pages"`
}

type ProductFamily struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Handle         string `json:"handle"`
	Description    string `json:"description"`
	AccountingCode string `json:"accounting_code"`
}

type PublicSignupPage struct {
	Id           int    `json:"id"`
	ReturnUrl    string `json:"return_url"`
	ReturnParams string `json:"return_params"`
	Url          string `json:"url"`
}

type CreditCard struct {
	Id                 int    `json:"id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	MaskedCardNumber   string `json:"masked_card_number"`
	CardType           string `json:"card_type"`
	ExpirationMonth    int    `json:"expiration_month"`
	ExpirationYear     int    `json:"expiration_year"`
	CustomerId         int    `json:"customer_id"`
	CurrentVault       string `json:"current_vault"`
	VaultToken         string `json:"vault_token"`
	BillingAddress     string `json:"billing_address"`
	BillingCity        string `json:"billing_city"`
	BillingState       string `json:"billing_state"`
	BillingZip         string `json:"billing_zip"`
	BillingCountry     string `json:"billing_country"`
	CustomerVaultToken string `json:"customer_vault_token"`
	BillingAddress2    string `json:"billing_address_2"`
	PaymentType        string `json:"payment_type"`
}

type Site struct {
	Id        int    `json:"id"`
	Subdomain string `json:"subdomain"`
}
