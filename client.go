package api

import (
	"log"

	"github.com/exotel/goapi/assets/types"
	"github.com/exotel/goapi/resources"
)

//Client is the iam client struct that calls all the other services
type Client struct {
	AccountSid   string `url:"AccountSid"`
	UserName     string
	ResourceID   string `url:"ResourceID"`
	url          string
	data         resources.IResource
	action       types.Action
	validActions types.Action
	baseURL      string
	logger       *log.Logger
	mode         types.Mode
	Credentials  types.Credentials
}

//New Creates an instance of iam client
func New(cr types.Credentials) *Client {
	return &Client{Credentials: cr, baseURL: "http://api.suffix.io/"}
}

//SetAccountSid sets the account sid of the user
func (c *Client) SetAccountSid(accountSid string) *Client {
	c.AccountSid = accountSid
	return c
}

//SetUserName sets the user name of the user
func (c *Client) SetUserName(userName string) *Client {
	c.UserName = userName
	return c
}

//Debug set the client in debug mode
func (c *Client) Debug(debug bool) *Client {
	if debug {
		c.mode = types.DEBUG
		return c
	}
	c.mode = types.PRODUCTION
	return c
}

//SetLogger sets logger for the class
func (c *Client) SetLogger(logger *log.Logger) *Client {
	c.logger = logger
	return c
}

//SetBaseURL sets the base url for making api requests
func (c *Client) SetBaseURL(url string) *Client {
	c.baseURL = url
	return c
}
