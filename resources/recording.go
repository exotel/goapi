package resources

import "github.com/exotel/goapi/assets/types"

//RecordingFilter struct for Get request data
type RecordingFilter struct {
	CallSid string `queryparam:"CallSid"`
}

//RecordingURLS saves the routes its in the format that the text/template library of
//go can accept
var RecordingURLS = map[types.Action]string{
	types.READ:     "/v1/Accounts/{{.AccountSid}}/Recordings/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/Recordings",
}
