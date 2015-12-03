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
####contributions
sarath@exotel.in
