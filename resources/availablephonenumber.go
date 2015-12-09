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

//AvailablePhoneNumberURLS method => urls maping
var AvailablePhoneNumberURLS = map[types.Action]string{
	types.READ:     "/v1/Accounts/{{.AccountSid}}/AvailablePhoneNumbers/{{.IsoCountry}}/Local",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/AvailablePhoneNumbers/{{.IsoCountry}}/Local",
}
