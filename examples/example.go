package main

import (
	"fmt"

	api "github.com/exotel/goapi"
	"github.com/exotel/goapi/assets/types"
	"github.com/exotel/goapi/resources"
)

func initClient() (client *api.Client) {
	credentials := types.Credentials{
		UserCredentials:   types.UserCredentials{AccessToken: "APIAccessTokenB1", UserName: "APIUsernameB1"},
		ClientCredentials: types.ClientCredentials{ClientID: "CLIENT_id", ClientSecret: "CLIENT_SECRET"},
	}
	client = api.New(credentials)
	client.SetAccountSid("AC2a002f82e0a34343be06063f2364c85d") //parent working
	return
}

//ExampleAccount shows all the account
func ExampleAccount(client *api.Client) {
	// CREATE
	c := client.Debug(true).
		Account().
		Create().
		Details(resources.AccountDetails{FriendlyName: "SARATH TESTING", CountryCode: "SNG"})
	if status, data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("No Error occured while doing the operation\nstatus : %d\ndata%v", status, data)
	}

	// GET - Single
	c = client.Debug(true).
		Account().
		ID("SUBACCOUNT_ID").
		Get()

	if status, data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("No Error occured while doing the operation\nstatus : %d\ndata%v", status, data)
	}

	// GET - bulk
	c = client.Debug(true).
		Account().
		Get()

	if status, data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("No Error occured while doing the operation\nstatus : %d\ndata%v", status, data)
	}

	//UPDATE
	c = client.
		Debug(true).
		Account().
		ID("SUBACCOUNT_ID").
		Update().
		UpdateDetails(resources.AccountUpdatableDetails{FriendlyName: "blinkasya"})
	if status, data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("No Error occured while doing the operation\nstatus : %d\ndata%v", status, data)
	}
}

//ExampleAvailablePhoneNumber has an example use of AvailablePhoneNumber resource
func ExampleAvailablePhoneNumber(client *api.Client) {
	var filter resources.AvailablePhoneNumberFilter
	filter.IsoCountry = "SG"
	if status, data, err := client.Debug(true).AvailablePhoneNumber().Get().Filter(filter).Do(); err != nil {
		fmt.Println("error occured", err.Error())
	} else {
		fmt.Printf("No Error occured while doing the operation\nstatus : %d\ndata%v", status, data)
	}
}

//ExampleCall Create call example
func ExampleCall(client *api.Client) {
	var callDetails resources.CallDetails
	callDetails.From = "+918030752401"
	callDetails.To = "+919742033616"
	callDetails.URL = "http://98ae7c6f.ngrok.io/dial/+919742033616"
	callDetails.Method = "GET"

	if status, data, err := client.Debug(true).Call().Create().Details(callDetails).Do(); err != nil {
		fmt.Println("error occured", err.Error())

	} else {
		fmt.Printf("No Error occured while doing the operation\nstatus : %d\ndata%v", status, data)
	}
}

func main() {
	client := initClient()
	ExampleCall(client)
}
