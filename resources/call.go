package resources

import "github.com/exotel/goapi/assets/types"

//CallDetails struct for CallDetails request data
type CallDetails struct {
	ParentAccountSid     string
	AccountSid           string
	Userid               string
	From                 string `mandatory:"true"`
	To                   string `mandatory:"true"`
	URL                  string `mandatory:"true"`
	Method               string
	FallbackURL          string
	FallbackMethod       string
	StatusCallback       string
	StatusCallbackMethod string
	StatusCallbackEvents []string
	SendDigits           string
	Timeout              int
	Record               bool
	IfMachine            string
	RequestID            string
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
