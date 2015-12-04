package resources

import "github.com/exotel/goapi/assets/types"

//CallDetails struct for CallDetails request data
type CallDetails struct {
	AccountSid           string   `queryparam:"AccountSid"`
	Userid               string   `queryparam:"Userid"`
	From                 string   `mandatory:"true" queryparam:"From"`
	To                   string   `mandatory:"true" queryparam:"To"`
	URL                  string   `mandatory:"true" queryparam:"URL"`
	Method               string   `queryparam:"Method"`
	FallbackURL          string   `queryparam:"FallbackURL"`
	FallbackMethod       string   `queryparam:"FallbackMethod"`
	StatusCallback       string   `queryparam:"StatusCallback"`
	StatusCallbackMethod string   `queryparam:"StatusCallbackMethod"`
	StatusCallbackEvents []string `queryparam:"StatusCallbackEvents"`
	SendDigits           string   `queryparam:"SendDigits"`
	Timeout              int      `queryparam:"Timeout"`
	Record               bool     `queryparam:"Record"`
	IfMachine            string   `queryparam:"IfMachine"`
	RequestID            string   `queryparam:"RequestID"`
}

//CallFilter the filter struct for call search
type CallFilter struct {
	CallSid string `mandatory:"true" url:"CallSid"`
}

//AvailablePhoneNumberURLS method => urls maping
var CallURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts/{{.AccountSid}}/Calls",
	types.READ:     "/v1/Accounts/{{.AccountSid}}/Calls/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/Calls/{{.CallSid}}",
}
