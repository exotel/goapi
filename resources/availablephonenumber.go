package resources

import "github.com/exotel/goapi/assets/types"

//AvailablePhoneNumberFilter struct for AvailablePhoneNumberFilter request data
type AvailablePhoneNumberFilter struct {
	ParentAccountSid string  `queryparam:"ParentAccountSid"`
	IsoCountry       string  `mandatory:"true" url:"IsoCountry"`
	Contains         *string `queryparam:"Contains"`
	SmsEnabled       bool    `queryparam:"SmsEnabled"`
	MmsEnabled       bool    `queryparam:"MmsEnabled"`
	VoiceEnabled     bool    `queryparam:"VoiceEnabled"`
	Region           string  `queryparam:"Region"`
}

// AvailablePhoneNumber is a struct for APN
type AvailablePhoneNumber struct {
	FriendlyName string       `json:"friendly_name"`
	PhoneNumber  string       `json:"phone_number"`
	Lata         *int         `json:"lata"`
	RateCenter   *string      `json:"rate_center"`
	Latitude     *float32     `json:"latitude"`
	Longitude    *float32     `json:"longitude"`
	Region       string       `json:"region"`
	PostalCode   *string      `json:"postal_code"`
	IsoCountry   string       `json:"iso_country"`
	Capabilities Capabilities `json:"capabilities"`
}

//GetAvailablePhoneNumberResponse  is the response structure of available phone numbers
//single get  request
type GetAvailablePhoneNumberResponse struct {
	AvailablePhoneNumber
}

//BulkGetAvailablePhoneNumberResponse  is the response structure of available phone numbers
//bulk get request
type BulkGetAvailablePhoneNumberResponse []GetAvailablePhoneNumberResponse

//AvailablePhoneNumberURLS method => urls maping
var AvailablePhoneNumberURLS = map[types.Action]string{
	types.READ:     "/v1/Accounts/{{.AccountSid}}/AvailablePhoneNumbers/{{.IsoCountry}}/Local",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/AvailablePhoneNumbers/{{.IsoCountry}}/Local",
}
