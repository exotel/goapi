package main

import (
	"fmt"

	api "github.com/exotel/goapi"
	"github.com/exotel/goapi/assets/types"
	"github.com/exotel/goapi/resources"
)

func main() {
	credentials := types.Credentials{
		UserCredentials:   types.UserCredentials{AccessToken: "ACCESSTOKEN"},
		ClientCredentials: types.ClientCredentials{ClientID: "CLIENT_id", ClientSecret: "CLIENT_SECRET"},
	}
	client := api.New(credentials)
	client.SetAccountSid("ACCOUNT_SID") //parent working

	// CREATE
	c := client.Debug(true).
		Account().
		Create().
		Details(resources.AccountDetails{FriendlyName: "SARATH TESTING", CountryCode: "SNG"})
	if data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("No Error occured while doing the operation", data)
	}

	// GET - Single
	c = client.Debug(true).
		Account().
		ID("SUBACCOUNT_ID").
		Get()

	if data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("No Error occured while doing the operation", data)
	}

	// GET - bulk
	c = client.Debug(true).
		Account().
		Get()

	if data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("No Error occured while doing the operation", data)
	}

	//UPDATE
	c = client.
		Debug(true).
		Account().
		ID("SUBACCOUNT_ID").
		Update().
		UpdateDetails(resources.AccountUpdatableDetails{FriendlyName: "blinkasya"})
	if data, err := c.Do(); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("No Error occured while doing the operation", data)
	}
}
