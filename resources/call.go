package resources

import (
	"time"

	"github.com/exotel/goapi/assets/types"
)

//CallDetails struct for CallDetails request data
type CallDetails struct {
	AccountSid           string   `queryparam:"AccountSid"`
	Userid               string   `queryparam:"Userid"`
	From                 string   `mandatory:"true" queryparam:"From"`
	To                   string   `mandatory:"true" queryparam:"To"`
	URL                  string   `mandatory:"true" queryparam:"Url"`
	Method               string   `queryparam:"Method"`
	FallbackURL          string   `queryparam:"FallbackURL"`
	FallbackMethod       string   `queryparam:"FallbackMethod"`
	StatusCallback       string   `queryparam:"StatusCallback"`
	StatusCallbackMethod string   `queryparam:"StatusCallbackMethod"`
	StatusCallbackEvents []string `queryparam:"StatusCallbackEvents"`
	SendDigits           string   `queryparam:"SendDigits"`
	Timeout              int      `queryparam:"Timeout"`
	Record               *bool    `queryparam:"Record"`
	IfMachine            string   `queryparam:"IfMachine"`
	RequestID            string   `queryparam:"RequestID"`
}

//CallFilter the filter struct for call search
type CallFilter struct {
	CallSid string `mandatory:"true" url:"CallSid"`
}

// Call is a struct for call
type Call struct {
	Sid             string      `json:"sid"`
	DateCreated     time.Time   `json:"date_created"`
	DateUpdated     time.Time   `json:"date_updated"`
	ParentCallSid   *string     `json:"parent_call_sid"`
	VirtualCallSid  string      `json:"virtual_call_sid"`
	AccountSid      string      `json:"account_sid"`
	To              string      `json:"to"`
	From            string      `json:"from"`
	PhoneNumberSid  *string     `json:"phone_number_sid"`
	Status          string      `json:"status"`
	StartTime       time.Time   `json:"start_time"`
	EndTime         time.Time   `json:"end_time"`
	Duration        *int64      `json:"duration"`
	Price           *float64    `json:"price"`
	PriceUnit       *string     `json:"price_unit"`
	Direction       string      `json:"direction"`
	AnsweredBy      *string     `json:"answered_by"`
	APIVersion      string      `json:"api_version"`
	ForwardedFrom   *string     `json:"forwarded_from"`
	CallerName      *string     `json:"caller_name"`
	URI             string      `json:"uri" xml:"Uri"`
	SubResourceURIs SubResource `json:"sub_resource_uris"`
}

//SubResource defines the fields that are linked to the call but need to be invoked separately
type SubResource struct {
	Recordings string `json:"recordings"`
}

//CreateCallResponse defines structure of create call response
type CreateCallResponse struct {
	Call
}

//GetCallResponse defines structure of get call details request response
type GetCallResponse struct {
	Call
}

//BulkGetCallResponse defines structure of bulk get call details request response
type BulkGetCallResponse []GetAccountResponse

//AvailablePhoneNumberURLS method => urls maping
var CallURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts/{{.AccountSid}}/Calls",
	types.READ:     "/v1/Accounts/{{.AccountSid}}/Calls/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/Calls/{{.CallSid}}",
}
