package resources

import "github.com/exotel/goapi/assets/types"

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

//IncomingPhoneNumberURLS saves the routes its in the format that the text/template library of
//go can accept
var IncomingPhoneNumberURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers",
	types.READ:     "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers",
	types.UPDATE:   "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers/{{.ResourceID}}",
	types.DELETE:   "/v1/Accounts/{{.AccountSid}}/IncomingPhoneNumbers/{{.ResourceID}}",
}
