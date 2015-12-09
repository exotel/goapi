package resources

import "github.com/exotel/goapi/assets/types"

//OutgoingCallerIDDetails struct for OutgoingCallerIDDetails request data
type OutgoingCallerIDDetails struct {
	PhoneNumber          string `mandatory:"true" queryparam:"PhoneNumber"`
	FriendlyName         string `queryparam:"FriendlyName"`
	StatusCallback       string `queryparam:"StatusCallback"`
	StatusCallbackMethod string `queryparam:"StatusCallbackMethod"`
}

//OutgoingCallerIDFilter struct for OutgoingCallerIDFilter request data
type OutgoingCallerIDFilter struct {
	PhoneNumber  string `queryparam:"PhoneNumber"`
	FriendlyName string `queryparam:"FriendlyName"`
}

//OutgoingCallerIDUpdateableDetails struct for OutgoingCallerIDUpdateableDetails request data
type OutgoingCallerIDUpdateableDetails struct {
	FriendlyName string `queryparam:"FriendlyName"`
}

//OutgoingCallerIDURLS saves the routes its in the format that the text/template library of
//go can accept
var OutgoingCallerIDURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs",
	types.READ:     "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs",
	types.DELETE:   "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs/{{.ResourceID}}",
}
