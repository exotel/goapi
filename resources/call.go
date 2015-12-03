package resources

//CallDetails struct for CallDetails request data
type CallDetails struct {
	ParentAccountSid     string
	AccountSid           string
	Userid               string
	From                 string `mandatory:"true"`
	To                   string `mandatory:"true"`
	URL                  string `mandatory:"true"`
	Method               string
	FallbackURL          string
	FallbackMethod       string
	StatusCallback       string
	StatusCallbackMethod string
	StatusCallbackEvents []string
	SendDigits           string
	Timeout              int
	Record               bool
	IfMachine            string
	RequestID            string
}
