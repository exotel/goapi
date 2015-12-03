package resources

import "github.com/exotel/goapi/assets/types"

//AvailablePhoneNumberFilter struct for AvailablePhoneNumberFilter request data
type AvailablePhoneNumberFilter struct {
	ParentAccountSid string
	IsoCountry       string `mandatory:"true"`
	Contains         *string
	SmsEnabled       bool
	MmsEnabled       bool
	VoiceEnabled     bool
	Region           *string
}

//AvailablePhoneNumberURLS method => urls maping
var AvailablePhoneNumberURLS = map[types.Action]string{
	types.READ:     "/v1/Accounts/{{.AccountSid}}/AvailablePhoneNumbers/{{.IsoCountry}}/Local",
	types.BULKREAD: "/v1/Accounts/{{.AccountSid}}/AvailablePhoneNumbers/{{.IsoCountry}}/Local",
}
