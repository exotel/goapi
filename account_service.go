package api

import (
	. "github.com/exotel/goapi/assets/types"
	"github.com/exotel/goapi/resources"
)

//AccountService is the struct that defines the account information and has fubnctions to perform account level
//Add ,Update ,getInformation etc
type AccountService struct {
	Client
}

//Create sets the action as create
func (accountService *AccountService) Create() *AccountService {
	accountService.action = CREATE
	accountService.url = resources.AccountURLS[CREATE]
	return accountService
}

//Update sets the action as update
func (accountService *AccountService) Update() *AccountService {
	accountService.action = UPDATE
	accountService.url = resources.AccountURLS[UPDATE]
	return accountService
}

//Delete sets the action as DELETE
func (accountService *AccountService) Delete() *AccountService {
	accountService.action = DELETE
	accountService.url = resources.AccountURLS[DELETE]
	return accountService
}

//Get sets the action as READ
func (accountService *AccountService) Get() *AccountService {
	accountService.action = BULKREAD
	accountService.url = resources.AccountURLS[READ]
	return accountService
}

//ID sets the id for the resource request
func (accountService *AccountService) ID(id string) *AccountService {
	accountService.ResourceID = id
	switch accountService.action {
	case BULKREAD:
		accountService.action = READ

	}
	return accountService
}

//Filter set the filter for read operation for read
func (accountService *AccountService) Filter(filter resources.AccountFilter) *AccountService {
	accountService.data = filter
	return accountService
}

//Details set the resource details for create resource requests
func (accountService *AccountService) Details(details resources.AccountDetails) *AccountService {
	accountService.data = details
	return accountService
}

//UpdateDetails sets the details to be updated for the resource
func (accountService *AccountService) UpdateDetails(details resources.AccountUpdatableDetails) *AccountService {
	accountService.data = details
	return accountService
}

//Account returns an instance of accountService
func (c *Client) Account() *AccountService {
	accountService := AccountService{Client: *c}
	accountService.validActions = CREATE | READ | UPDATE
	return &accountService
}
