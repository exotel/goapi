package resources

import "github.com/exotel/goapi/assets/types"

//OutgoingCallerIDDetails struct for OutgoingCallerIDDetails request data
type OutgoingCallerIDDetails struct {
	PhoneNumber               string `mandatory:"true"`
	ParentOutgoingCallerIDSid string
	FriendlyName              string
	StatusCallback            string
	StatusCallbackMethod      string
}

//OutgoingCallerIDFilter struct for OutgoingCallerIDFilter request data
type OutgoingCallerIDFilter struct {
	PhoneNumber  string
	FriendlyName string
}

//OutgoingCallerIDUpdateableDetails struct for OutgoingCallerIDUpdateableDetails request data
type OutgoingCallerIDUpdateableDetails struct {
	FriendlyName string
}

//OutgoingCallerIDURLS saves the routes its in the format that the text/template library of
//go can accept
var OutgoingCallerIDURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs",
	types.READ:     "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs",
	types.DELETE:   "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs/{{.ResourceID}}",
}
