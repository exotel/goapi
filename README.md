goexotel
=====
![Go-Exotel](http://homegrown.co.in/wp-content/uploads/2014/07/master-shifu-1024x436.jpg)
####What is it?
  GoExotel is a exotel platform api client written in golang

####resources
* Account
* IncomingPhoneNumber  
* OutgoingCallerID  
* AvailablePhoneNumber
* Call  
* Recording

####how to use
Install the library using `go get`
```
go get https://github.com/exotel/goapi
```
#####using the client
```
credential := goapi.Credential{
    UserCredentials: goapi.UserCredentials{AccessToken: "1212121"},
    ClientCredentials: goapi.ClientCredentials{ClientID: "asasa"}
  }
client := goapi.New(credential)
client.SetAccountSid("ACCONTSID")
```
for adding a sub account

```
 credential = gapi.Credential{
   goapi.UserCredentials{AccessToken: "1212121"},
   goapi.ClientCredentials{ClientID: ""},
 }

 accountDetails := resources.AccountDetails{FriendlyName: "SUB-ACCOUNTNAME", CountryCode: "SNG"}
 client := goapi.New(credential)
 client.SetAccountSid("ACCONTSID")
 result,err := client
 Account().
 Create().
 Details(accountDetails).
 Do()

if err != nil {
 log.Fatal(err)
 return
}

if result != nil  {
 //Do whatever
}

```


####Sample Code
```
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

```
####contributions
sarath@exotel.in
