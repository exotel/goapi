goexotel
=====
![Go-Exotel](http://homegrown.co.in/wp-content/uploads/2014/07/master-shifu-1024x436.jpg)
####What is it?
  GoExotel is a exotel platform api client written in golang

####services
  * account management
  * Call
  * Sms
  * Billing

####how to use
Install the library using `go get`
```
go get bitbucket.org/Exotel/rohan/platform
```
_**PS**:Now go get tool can not get private repositories initially so you need to clone the repo for the first time and then do go get every other time to update_

Use this [gist](
https://gist.github.com/sarathsp06/30710cfa1749eaf7c4f8
) to do the same


#####using the client
```
credential := iam.Credential{
    UserCredentials: iam.UserCredentials{AccessToken: "1212121"},
    ClientCredentials: iam.ClientCredentials{ClientID: "asasa"}
  }
client := iam.New(credential)
```
The client is initialized using the credentials and for any services there ,ie:for access checking service

```
output,err := client.Check().SetScopes(scopeSet).Do()
if err != nil {
  log.Fatal(err)
  return
}
if output.Success && output.Authorized  {
  //The user is authorized with the set of scopes provided
}

```
for adding a user

```

//for creating a user only AccessToken is required later if
//the same client has to be used for any other services
//set the Credentials by using client.Credentals.[whatever]
userInfo = UserInfo{
  "exotel",
  gandalf.BasicUserInfo{
    Name:  "Sarath",
    Email: "sarath@exotel.in"},
  gandalf.DetailedUserInfo{},
}
credential = Credential{
   iam.UserCredentials{AccessToken: "1212121"},
   iam.ClientCredentials{ClientID: ""},
 }

 user,err := New(credential).
 User().
 SetUser(userInfo).
 Insert().
 Do()

if err != nil {
 log.Fatal(err)
 return
}

if user.Suceess {
 //Do whatever userid is user.UserID
}

```




####contributions
sarath@exotel.in
