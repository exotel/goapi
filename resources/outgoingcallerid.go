package resources

import (
	"encoding/xml"
	"time"

	"github.com/exotel/goapi/assets/types"
)

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

// OutgoingCallerID is a struct
type OutgoingCallerID struct {
	XMLName      xml.Name  `json:"-"`
	Sid          string    `json:"sid"`
	AccountSid   string    `json:"account_sid"`
	FriendlyName string    `json:"friendly_name"`
	PhoneNumber  string    `json:"phone_number"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	URI          string    `json:"uri"`
}

//CreateOutgoingCallerIDResponse defines the structure ofresponse
//on ccreate CreateOutgoingCallerIDResponse request
type CreateOutgoingCallerIDResponse struct {
	XMLName        xml.Name `json:"-"`
	AccountSid     string
	PhoneNumber    string
	FriendlyName   *string
	ValidationCode int
	CallSid        string
}

//UpdateOutgoingCallerIDResponse has the struct definition for OutgoingCallerID update request
type UpdateOutgoingCallerIDResponse struct {
	OutgoingCallerID
}

//GetOutgoingCallerIDResponse has the response structure for Get requests
type GetOutgoingCallerIDResponse struct {
	OutgoingCallerID
}

//BulkGetOutgoingCallerIDResponse has the response structure for bulk Get requests
type BulkGetOutgoingCallerIDResponse []OutgoingCallerID

//OutgoingCallerIDURLS saves the routes its in the format that the text/template library of
//go can accept
var OutgoingCallerIDURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs",
	types.READ:     "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs",
	types.DELETE:   "/v1/Accounts/{{.AccountSid}}}/OutgoingCallerIDs/{{.ResourceID}}",
}
