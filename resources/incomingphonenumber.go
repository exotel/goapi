package resources

import (
	"time"

	"github.com/exotel/goapi/assets/types"
)

//IncomingPhoneNumberDetails struct for IncomingPhoneNumberDetails request data
type IncomingPhoneNumberDetails struct {
	PhoneNumber          string `mandatory:"true" queryparam:"PhoneNumber"`
	FriendlyName         string `queryparam:"FriendlyName"`
	APIVersion           string `queryparam:"APIVersion"`
	VoiceURL             string `queryparam:"VoiceURL"`
	VoiceMethod          string `queryparam:"VoiceMethod"`
	VoiceFallbackURL     string `queryparam:"VoiceFallbackURL"`
	VoiceFallbackMethod  string `queryparam:"VoiceFallbackMethod"`
	StatusCallback       string `queryparam:"StatusCallback"`
	StatusCallbackMethod string `queryparam:"StatusCallbackMethod"`
	VoiceCallerIDLookup  string `queryparam:"VoiceCallerIDLookup"`
	VoiceApplicationSid  string `queryparam:"VoiceApplicationSid"`
	TrunkSid             string `queryparam:"TrunkSid"`
	SmsURL               string `queryparam:"SmsURL"`
	SmsMethod            string `queryparam:"SmsMethod"`
	SmsFallbackURL       string `queryparam:"SmsFallbackURL"`
	SmsFallbackMethod    string `queryparam:"SmsFallbackMethod"`
	SmsApplicationSid    string `queryparam:"SmsApplicationSid"`
}

//IncomingPhoneNumberFilter struct for IncomingPhoneNumberFilter request data
type IncomingPhoneNumberFilter struct {
	PhoneNumber  string
	FriendlyName string
}

//IncomingPhoneNumberUpdatableDetails struct for IncomingPhoneNumberUpdatableDetails request data
type IncomingPhoneNumberUpdatableDetails struct {
	Sid                  string `queryparam:"Sid"`
	FriendlyName         string `queryparam:"FriendlyName"`
	APIVersion           string `queryparam:"APIVersion"`
	VoiceURL             string `queryparam:"VoiceURL"`
	VoiceMethod          string `queryparam:"VoiceMethod"`
	VoiceFallbackURL     string `queryparam:"VoiceFallbackURL"`
	VoiceFallbackMethod  string `queryparam:"VoiceFallbackMethod"`
	StatusCallback       string `queryparam:"StatusCallback"`
	StatusCallbackMethod string `queryparam:"StatusCallbackMethod"`
	VoiceCallerIDLookup  string `queryparam:"VoiceCallerIDLookup"`
	VoiceApplicationSid  string `queryparam:"VoiceApplicationSid"`
	TrunkSid             string `queryparam:"TrunkSid"`
	SmsURL               string `queryparam:"SmsURL"`
	SmsMethod            string `queryparam:"SmsMethod"`
	SmsFallbackURL       string `queryparam:"SmsFallbackURL"`
	SmsFallbackMethod    string `queryparam:"SmsFallbackMethod"`
	SmsApplicationSid    string `queryparam:"SmsApplicationSid"`
}

// IncomingPhoneNumber is a struct for IPN Details in response
type IncomingPhoneNumber struct {
	Sid                  string       `json:"sid"`
	AccountSid           string       `json:"account_sid"`
	FriendlyName         string       `json:"friendly_name"`
	PhoneNumber          string       `json:"phone_number"`
	VoiceURL             *string      `json:"voice_url"`
	VoiceMethod          string       `json:"voice_method"`
	VoiceFallbackURL     *string      `json:"voice_fallback_url"`
	VoiceFallbackMethod  string       `json:"voice_fallback_method"`
	VoiceCallerIDLookup  *string      `json:"voice_called_id_lookup"`
	DateCreated          time.Time    `json:"date_created"`
	DateUpdated          time.Time    `json:"date_updated"`
	SmsURL               *string      `json:"sms_url"`
	SmsMethod            string       `json:"sms_method"`
	SmsFallbackURL       *string      `json:"sms_fallback_url"`
	SmsFallbackMethod    string       `json:"sms_fallback_method"`
	Capabilities         Capabilities `json:"capabilities"`
	StatusCallback       *string      `json:"status_callback"`
	StatusCallbackMethod string       `json:"status_callback_method"`
	APIVersion           string       `json:"api_version"`
	URI                  string       `json:"uri"`
}

//CreateIncomingPhoneNumberResponse has the struct definition for account creation rsponse
type CreateIncomingPhoneNumberResponse struct {
	IncomingPhoneNumber
}

//UpdateIncomingPhoneNumberResponse has the struct definition for IncomingPhoneNumber update request
type UpdateIncomingPhoneNumberResponse struct {
	IncomingPhoneNumber
}

//GetIncomingPhoneNumberResponse has the response structure for Get requests
type GetIncomingPhoneNumberResponse struct {
	IncomingPhoneNumber
}

//BulkGetIncomingPhoneNumberResponse defines structure of bulk get call details request response
type BulkGetIncomingPhoneNumberResponse []GetAccountResponse

//IncomingPhoneNumberURLS saves the routes its in the format that the text/template library of
//go can accept
var IncomingPhoneNumberURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers",
	types.READ:     "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers",
	types.UPDATE:   "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers/{{.ResourceID}}",
	types.DELETE:   "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers/{{.ResourceID}}",
}
