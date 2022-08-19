package client

import (
	"fmt"
	"net/http"

	"github.com/ferchuhill/form3-client/client/models"
)

// NewAccountService to create a new service to comunicate the the account data in form3.
func NewAccountService(c *Client) *AccountService {
	if c == nil {
		c = NewClient()
	}
	acc := AccountService{client: c}
	return &acc
}

// this Stuct has the client to make the api calls
type AccountService struct {
	client *Client
}

// AccountData is an envelope to parse the json recive and send by the api.
type AccountData struct {
	Data *models.Account `json:"data"`
}

// Fetch to find a account by Id
func (ac *AccountService) Fetch(id string) (*models.Account, error) {
	var account AccountData
	target := fmt.Sprintf("%s%s", "/v1/organisation/accounts/", id)
	err := ac.client.CallAPI(http.MethodGet, target, nil, &account)
	return account.Data, err
}

// Create to save a new account
func (ac *AccountService) Create(acData models.Account) (models.Account, error) {
	data := models.Account{}
	envelope := AccountData{Data: &acData}
	err := ac.client.CallAPI(http.MethodPost, "/v1/organisation/accounts", envelope, &data)
	return acData, err
}

// Delete to delete an account
func (ac *AccountService) Delete(id string, version int64) (bool, error) {
	data := models.Account{}
	target := fmt.Sprintf("%s%s?version=%d", "/v1/organisation/accounts/", id, version)
	err := ac.client.CallAPI(http.MethodDelete, target, nil, &data)
	if err != nil {
		return false, err
	}
	return true, nil
}
