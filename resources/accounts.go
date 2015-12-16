package resources

import (
	"encoding/xml"
	"time"

	"github.com/exotel/goapi/assets/types"
)

//AccountDetails struct for AccountDetails request data
type AccountDetails struct {
	FriendlyName string `mandatory:"true" url:"FriendlyName"`
	CountryCode  string `mandatory:"true" queryparam:"CountryCode"`
}

// Account is a struct for Account details in the account response
type Account struct {
	XMLName         xml.Name  `json:"-"`
	Sid             string    `json:"sid"`
	DateCreated     time.Time `json:"date_created"`
	DateUpdated     time.Time `json:"date_updated"`
	FriendlyName    string    `json:"friendly_name"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	AuthToken       string    `json:"auth_token"`
	URI             string    `json:"uri"`
	OwnerAccountSid string    `json:"owner_account_sid,omitempty"`
}

//CreateAccountResponse has the struct definition for account creation rsponse
type CreateAccountResponse struct {
	Account
}

//UpdateAccountResponse has the struct definition for Account update request
type UpdateAccountResponse struct {
	Account
}

//GetAccountResponse has the response structure for Get requests
type GetAccountResponse struct {
	Account
}

//BulkGetAccountResponse has the response structure for bulk  Get requests
type BulkGetAccountResponse []GetAccountResponse

//AccountFilter struct for AccountFilter request data
type AccountFilter struct {
	FriendlyName string `queryparam:"FriendlyName"`
	Status       string `queryparam:"Status"`
}

//AccountUpdatableDetails struct for Update request data
type AccountUpdatableDetails struct {
	FriendlyName string `queryparam:"FriendlyName"`
	Status       string `queryparam:"Status"`
}

//AccountURLS saves the routes its in the format that the text/template library of
//go can accept
var AccountURLS = map[types.Action]string{
	types.CREATE:   "/v1/Accounts",
	types.READ:     "/v1/Accounts/{{.ResourceID}}",
	types.BULKREAD: "/v1/Accounts",
	types.UPDATE:   "/v1/Accounts/{{.ResourceID}}",
	types.DELETE:   "/v1/Accounts/{{.ResourceID}}",
}
