package resources

import (
	"time"

	"github.com/exotel/goapi/assets/types"
)

//RecordingFilter struct for Get request data
type RecordingFilter struct {
	CallSid string `queryparam:"CallSid"`
}

// Recording is a response for get recording resource.
type Recording struct {
	Sid         string    `json:"sid"`
	AccountSid  string    `json:"account_sid"`
	CallSid     string    `json:"call_sid"`
	Duration    int64     `json:"duration"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	APIVersion  string    `json:"api_version"`
	URI         string    `json:"uri" xml:"Uri"`
}

//GetRecordingResponse has the response structure for Get requests
type GetRecordingResponse struct {
	Recording
}

//BulkGetRecordingResponse has the response structure for bulk  Get requests
type BulkGetRecordingResponse []GetRecordingResponse

//RecordingURLS saves the routes its in the format that the text/template library of
//go can accept
var RecordingURLS = map[types.Action]string{
	types.READ:     "/v1/Accounts/{{.AccountSid}}/Recordings/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/Recordings",
}
