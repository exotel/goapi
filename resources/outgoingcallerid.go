package resources

//OutgoingCallerIDDetails struct for OutgoingCallerIDDetails request data
type OutgoingCallerIDDetails struct {
	PhoneNumber          string `mandatory:"true"`
	ParentAccountSid     string
	FriendlyName         string
	StatusCallback       string
	StatusCallbackMethod string
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
