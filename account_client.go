// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING;
// Autogenerated by clientgenerator -r Account -m CRU DO NOT EDIT

package api

import (
	"github.com/exotel/goapi/assets/types"
	"github.com/exotel/goapi/resources"
)

//AccountService is the struct that defines the account information and has fubnctions to perform account level
type AccountService struct {
	Client
}

//Create sets the action as create
func (__receiver_AService *AccountService) Create() *AccountService {
	__receiver_AService.action = types.CREATE
	__receiver_AService.data = resources.AccountDetails{}
	__receiver_AService.url = resources.AccountURLS[types.CREATE]
	return __receiver_AService
}

//Update sets the action as update
func (__receiver_AService *AccountService) Update() *AccountService {
	__receiver_AService.action = types.UPDATE
	__receiver_AService.data = resources.AccountUpdatableDetails{}
	__receiver_AService.url = resources.AccountURLS[types.UPDATE]
	return __receiver_AService
}

//Get sets the action as types.READ
func (__receiver_AService *AccountService) Get() *AccountService {
	if len(__receiver_AService.ResourceID) == 0 {
		__receiver_AService.action = types.BULKREAD
		__receiver_AService.data = resources.AccountFilter{}
		__receiver_AService.url = resources.AccountURLS[types.BULKREAD]
	} else {
		__receiver_AService.data = struct{}{}
		__receiver_AService.url = resources.AccountURLS[types.READ]
		__receiver_AService.action = types.READ
	}
	return __receiver_AService
}

//ID sets the id for the resource request
func (__receiver_AService *AccountService) ID(id string) *AccountService {
	__receiver_AService.ResourceID = id
	switch __receiver_AService.action {
	case types.BULKREAD:
		__receiver_AService.data = struct{}{}
		__receiver_AService.url = resources.AccountURLS[types.READ]
		__receiver_AService.action = types.READ

	}
	return __receiver_AService
}

//Filter set the filter for read operation for read
func (__receiver_AService *AccountService) Filter(filter resources.AccountFilter) *AccountService {
	__receiver_AService.data = filter
	return __receiver_AService
}

//Details set the resource details for create resource requests
func (__receiver_AService *AccountService) Details(details resources.AccountDetails) *AccountService {
	__receiver_AService.data = details
	return __receiver_AService
}

//UpdateDetails sets the details to be updated for the resource
func (__receiver_AService *AccountService) UpdateDetails(details resources.AccountUpdatableDetails) *AccountService {
	__receiver_AService.data = details
	return __receiver_AService
}

//Account returns an instance of AccountService
func (c *Client) Account() *AccountService {
	accountService := AccountService{Client: *c}
	accountService.validActions = types.CREATE | types.READ | types.BULKREAD | types.UPDATE | 0x00
	return &accountService
}
