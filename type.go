package chargify_webhook

import "time"

type FormattedTime struct {
	*time.Time
}

const timeFormat = `"2006-01-02 15:04:05 -0700"`

func (f *FormattedTime) MarshalJSON() ([]byte, error) {
	return []byte(f.Format(timeFormat)), nil
}

func (f *FormattedTime) UnmarshalJSON(input []byte) error {
	t, err := time.Parse(timeFormat, string(input))
	if err != nil {
		return err
	}
	*f = FormattedTime{&t}
	return nil
}

type Subscription struct {
	Id                      int            `json:"id,string"`
	State                   string         `json:"state"`
	TrialStartedAt          *FormattedTime `json:"trial_started_at"`
	Customer                *Customer      `json:"customer"`
	Product                 *Product       `json:"product"`
	CreditCard              *CreditCard    `json:"credit_card"`
	TrialEndedAt            *FormattedTime `json:"trial_ended_at"`
	ActivatedAt             *FormattedTime `json:"activated_at"`
	CreatedAt               *FormattedTime `json:"created_at"`
	UpdatedAt               *FormattedTime `json:"updated_at"`
	ExpiresAt               *FormattedTime `json:"expires_at"`
	PreviousExpiresAt       *FormattedTime `json:"previous_expires_at"`
	BalanceInCents          int            `json:"balance_in_cents,string"`
	CurrentPeriodEndsAt     *FormattedTime `json:"current_period_ends_at"`
	NextAssessmentAt        *FormattedTime `json:"next_assessment_at"`
	CanceledAt              *FormattedTime `json:"canceled_at"`
	CancellationMessage     string         `json:"cancellation_message"`
	NextProductId           int            `json:"next_product_id,string"`
	CancelAtEndOfPeriod     bool           `json:"cancel_at_end_of_period,string"`
	PaymentCollectionMethod string         `json:"payment_collection_method"`
	SnapDay                 string         `json:"snap_day"`
	CancellationMethod      string         `json:"cancellation_method"`
	CurrentPeriodStartAt    *FormattedTime `json:"current_period_started_at"`
	PreviousState           string         `json:"previous_state"`
	SignupPaymentId         int            `json:"signup_payment_id,string"`
	SignupRevenue           float32        `json:"signup_revenue,string"`
	DelayedCancelAt         *FormattedTime `json:"delayed_cancel_at"`
	CouponCode              string         `json:"coupon_code"`
	TotalRevenueInCents     int            `json:"total_revenue_in_cents,string"`
	ProductPriceInCents     int            `json:"product_price_in_cents,string"`
	ProductVersionNumber    int            `json:"product_version_number,string"`
	PaymentType             string         `json:"payment_type"`
	ReferralCode            string         `json:"referral_code"`
	CouponUseCount          int            `json:"coupon_use_count,string"`
	CouponUsesAllowed       int            `json:"coupon_uses_allowed,string"`
}

type Customer struct {
	Id                         int            `json:"id,string"`
	FirstName                  string         `json:"first_name"`
	LastName                   string         `json:"last_name"`
	Organization               string         `json:"organization"`
	Email                      string         `json:"email"`
	CreatedAt                  *FormattedTime `json:"created_at"`
	UpdatedAt                  *FormattedTime `json:"updated_at"`
	Reference                  string         `json:"reference"`
	Address                    string         `json:"address"`
	Address2                   string         `json:"address_2"`
	City                       string         `json:"city"`
	State                      string         `json:"state"`
	Zip                        string         `json:"zip"`
	Country                    string         `json:"country"`
	Phone                      string         `json:"phone"`
	PortalInviteLastSentAt     *FormattedTime `json:"portal_invite_last_sent_at"`
	PortalInviteLastAcceptedAt *FormattedTime `json:"portal_invite_last_accepted_at"`
	Verified                   bool           `json:"verified,string"`
	PortalCustomerCreatedAt    *FormattedTime `json:"portal_customer_created_at"`
	CcEmails                   string         `json:"cc_emails"`
	TaxExempt                  bool           `json:"tax_exempt,string"`
}

type Product struct {
	Id                      int                 `json:"id,string"`
	Name                    string              `json:"name"`
	Handle                  string              `json:"handle"`
	Description             string              `json:"description"`
	AccountingCode          string              `json:"accounting_code"`
	RequestCreditCard       bool                `json:"request_credit_card,string"`
	ExpirationInterval      int                 `json:"expiration_interval,string"`
	ExpirationIntervalUnit  string              `json:"expiration_interval_unit"`
	CreatedAt               *FormattedTime      `json:"created_at"`
	UpdatedAt               *FormattedTime      `json:"updated_at"`
	PriceInCents            int                 `json:"price_in_cents,string"`
	Interval                int                 `json:"interval,string"`
	IntervalUnit            string              `json:"interval_unit"`
	InitialChargeInCents    int                 `json:"initial_charge_in_cents,string"`
	TrialPriceInCents       int                 `json:"trial_price_in_cents,string"`
	TrialInterval           int                 `json:"trial_interval,string"`
	TrialIntervalUnit       string              `json:"trial_interval_unit"`
	ArchivedAt              *FormattedTime      `json:"archived_at"`
	RequireCreditCard       bool                `json:"require_credit_card,string"`
	ReturnParams            string              `json:"return_params"`
	Taxable                 bool                `json:"taxable,string"`
	UpdateReturnUrl         string              `json:"update_return_url"`
	InitialChargeAfterTrial bool                `json:"initial_charge_after_trial,string"`
	VersionNumber           int                 `json:"version_number,string"`
	UpdateReturnParams      string              `json:"update_return_params"`
	ProductFamily           *ProductFamily      `json:"product_family"`
	PublicSignupPages       []*PublicSignupPage `json:"public_signup_pages"`
}

type ProductFamily struct {
	Id             int    `json:"id,string"`
	Name           string `json:"name"`
	Handle         string `json:"handle"`
	Description    string `json:"description"`
	AccountingCode string `json:"accounting_code"`
}

type PublicSignupPage struct {
	Id           int    `json:"id,string"`
	ReturnUrl    string `json:"return_url"`
	ReturnParams string `json:"return_params"`
	Url          string `json:"url"`
}

type CreditCard struct {
	Id                 int    `json:"id,string"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	MaskedCardNumber   string `json:"masked_card_number"`
	CardType           string `json:"card_type"`
	ExpirationMonth    int    `json:"expiration_month,string"`
	ExpirationYear     int    `json:"expiration_year,string"`
	CustomerId         int    `json:"customer_id,string"`
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
	Id        int    `json:"id,string"`
	Subdomain string `json:"subdomain"`
}
