package resources

import "github.com/exotel/goapi/assets/types"

//IncomingPhoneNumberDetails struct for IncomingPhoneNumberDetails request data
type IncomingPhoneNumberDetails struct {
	PhoneNumber          string `mandatory:"true"`
	ParentAccountSid     string
	FriendlyName         string
	APIVersion           string
	VoiceURL             string
	VoiceMethod          string
	VoiceFallbackURL     string
	VoiceFallbackMethod  string
	StatusCallback       string
	StatusCallbackMethod string
	VoiceCallerIDLookup  string
	VoiceApplicationSid  string
	TrunkSid             string
	SmsURL               string
	SmsMethod            string
	SmsFallbackURL       string
	SmsFallbackMethod    string
	SmsApplicationSid    string
	Auth                 string
}

//IncomingPhoneNumberFilter struct for IncomingPhoneNumberFilter request data
type IncomingPhoneNumberFilter struct {
	PhoneNumber  string
	FriendlyName string
}

//IncomingPhoneNumberUpdatableDetails struct for IncomingPhoneNumberUpdatableDetails request data
type IncomingPhoneNumberUpdatableDetails struct {
	Sid                  string
	FriendlyName         string
	APIVersion           string
	VoiceURL             string
	VoiceMethod          string
	VoiceFallbackURL     string
	VoiceFallbackMethod  string
	StatusCallback       string
	StatusCallbackMethod string
	VoiceCallerIDLookup  string
	VoiceApplicationSid  string
	TrunkSid             string
	SmsURL               string
	SmsMethod            string
	SmsFallbackURL       string
	SmsFallbackMethod    string
	SmsApplicationSid    string
	AccountSid           string
	ParentAccountSid     string
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
