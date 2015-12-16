package types

//Credentials user credentials and client credentials
type Credentials struct {
	UserCredentials
	ClientCredentials
}

//UserCredentials the users credentials
type UserCredentials struct {
	UserName    string
	AccessToken string
}

//ClientCredentials The clients credetials  the users credentials
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

//Mode The possible client running action
type Mode uint8

const (
	//PRODUCTION production mode says its not in debug mode
	PRODUCTION Mode = iota
	//DEBUG mode says that the client is in debug mode
	DEBUG
)

// Error is a struct for platform error
type Error struct {
	Code     int    `json:"code"`
	HTTPCode int    `json:"status"`
	Message  string `json:"message"`
}

//Action Has the set of actions possible
type Action uint8

const (
	//CREATE create resource request
	CREATE Action = 0X01
	//BULKREAD get request on a resource
	//READ is by default bulk
	BULKREAD Action = 0X02
	//READ flag on action
	READ Action = 0x04
	//UPDATE update request on a resource
	UPDATE Action = 0X08
	//DELETE delete request on a resource
	DELETE Action = 0X10
)

//String implements the fmt.Stringer interface
func (a Action) String() (action string) {
	switch a {
	case CREATE:
		action = "CREATE"
	case BULKREAD:
		action = "READ"
	case READ:
		action = "READ"
	case UPDATE:
		action = "UPDATE"
	case DELETE:
		action = "DELETE"
	default:
		action = "INVALIDMETHOD"
	}
	return
}
