package resources

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
