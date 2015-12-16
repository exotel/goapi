goexotel [![GoDoc](https://godoc.org/github.com/exotel/goapi?status.svg)](https://godoc.org/github.com/exotel/goapi)
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
go get github.com/exotel/goapi
```

or for keeping the version number
```
go get gopkg.in/exotel/goapi.v0 //for version v0.x.x
```

#####using the client
initiate the client
```
credentials := types.Credentials{
  UserCredentials:   types.UserCredentials{AccessToken: "APIAccessToken", UserName: "APIUsername"},
  ClientCredentials: types.ClientCredentials{ClientID: "CLIENT_id", ClientSecret: "CLIENT_SECRET"},
}
client = api.New(credentials)
client.SetAccountSid("<ACCOUNT_SID>") //parent working
```

for adding a sub account

```
c := client.
  Account().
  Create().
  Details(resources.AccountDetails{FriendlyName: "SARATH TESTING", CountryCode: "SNG"})
if status, data, err := c.Do(); err != nil {
  fmt.Printf(err.Error())
} else {
  fmt.Printf("No Error occured while doing the operation\nstatus : %d\ndata%v", status, data)
}
```


#####For any resource do the following
consider resource is  `Resource` and the CRUD is allowed for `Resource`


1 . Create Request
```
client.Resource().Create().Details(resources.ResourceDetails{}).Do
```  
2 . Read Request
* bulk
```
client.Resource().Get().Filter(resources.ResourceFilter{}).Do()
```
* Single
```
client.Resource().ID(string).Get().Do
```

3 . Update Request
```
client.Resource().ID(string).Update().UpdateDetails(resources.ResourceUpdatableDetails{}).Do
```


4 . Delete Request
```
client.Resource().ID(string).Delete().Do()
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

func initClient() (client *api.Client) {
	credentials := types.Credentials{
		UserCredentials:   types.UserCredentials{AccessToken: "APIAccessToken", UserName: "APIUsername"},
		ClientCredentials: types.ClientCredentials{ClientID: "CLIENT_id", ClientSecret: "CLIENT_SECRET"},
	}
	client = api.New(credentials)
	client.SetAccountSid("<ACCOUNT_SID>") //parent working
	return
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
  //makes a call
	ExampleCall(client)
}

```



####What if i don't have go installed ?
This is how you can install go in linux [this is for version 1.5.2 of golang]
```
wget https://storage.googleapis.com/golang/go1.5.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.5.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

now you have go installed,try checking by checking the version
```
go version
```

Now you have to set the Workspace,say you are setting,~/go as the worksapce
this is how it is done
```
export GOPATH=$HOME/go
```

Add these two  export statements in ~/.bashrc so that its set all the time
Now you are ready to test goexoml,get the library using go get
ie,
```
go get github.com/exotel/goexoml
```

Now you can copy paste the sample codes from example folder to `path/to/filename.go` and test by running as
```
go run path/to/filename.go
````


####contributions
sarath@exotel.in
