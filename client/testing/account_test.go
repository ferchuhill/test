package testingF3UnitTesting

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ferchuhill/form3-client/client"
	model "github.com/ferchuhill/form3-client/client/models"
	"github.com/stretchr/testify/assert"
)

const expected = `{"data": {"type": "%s","id": "%s","organisation_id": "%s","version": "%d","attributes": {"account_number": "%s","country": "%s","base_currency": "%s","bank_id": "%s","bank_id_code": "%s","name": ["%s"]}}}`
const expectedRecive = `{"data": {"type": "%s","id": "%s","organisation_id": "%s","version": %d,"attributes": {"account_number": "%s","country": "%s","base_currency": "%s","bank_id": "%s","bank_id_code": "%s","name": ["%s"]}}}`

const types = "accounts"
const id = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
const idOrg = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
const accountNumber = "41426819"
const baseCurrency = "GBP"
const bankId = "400300"
const bankCode = "GBDSC"
const personName = "Samantha Holder"

var version = int64(0)
var country = "GB"

// TestAccountCreate is the unitTest for create a new resource with a stub account services
func TestAccountCreate(t *testing.T) {
	assert := assert.New(t)

	clientF3 := client.NewClient()

	//returnJson in the server
	expectedJson := fmt.Sprintf(expected, types, id, idOrg, version, accountNumber, country, baseCurrency, bankId, bankCode, personName)

	//stub server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(404)
		}
		assert.Equal(r.URL.Path, "/v1/organisation/accounts", "it should be calling the accounts url")
		fmt.Fprintf(w, expectedJson)
	}))
	defer svr.Close()

	//set the stub Server in the client
	clientF3.SetURL(svr.URL)
	clientAccount := client.NewAccountService(clientF3)

	//prepare the data
	names := []string{personName}
	attr := model.AccountAttributes{AccountNumber: accountNumber, Country: &country, BankID: bankId, BankIDCode: bankCode, BaseCurrency: baseCurrency, Name: names}
	account := model.Account{Attributes: &attr, ID: id, OrganisationID: idOrg, Type: types, Version: version}

	//call
	accountReturn, err := clientAccount.Create(account)

	//assert
	assert.Nil(err, "No error should be")
	assert.Equal(account, accountReturn, "The account send and recive should be equals")
}

// TestAccountFetch is the unitTest for find by id in stub account services
func TestAccountFetch(t *testing.T) {
	assert := assert.New(t)

	clientF3 := client.NewClient()

	//returnJson in the server
	expectedJson := fmt.Sprintf(expectedRecive, types, id, idOrg, version, accountNumber, country, baseCurrency, bankId, bankCode, personName)
	target := fmt.Sprintf("%s%s", "/v1/organisation/accounts/", id)

	//stub server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(404)
		}
		assert.Equal(r.URL.Path, target)
		fmt.Fprintf(w, expectedJson)
	}))
	defer svr.Close()

	//set the stub Server in the client
	clientF3.SetURL(svr.URL)

	//call
	clientAccount := client.NewAccountService(clientF3)
	accountReturn, err := clientAccount.Fetch(id)

	//assert
	assert.Nil(err, "No error should be")
	assert.Equal(accountReturn.ID, id, "The account id send and recive should be equals")
}

// TestAccountDelete is the unitTest for delete a resource with a stub account services
func TestAccountDelete(t *testing.T) {
	assert := assert.New(t)

	clientF3 := client.NewClient()

	target := fmt.Sprintf("%s%s", "/v1/organisation/accounts/", id)
	//stub server
	query := fmt.Sprintf("version=%d", version)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(404)
		}
		assert.Equal(r.URL.Path, target, "it should be calling the  accounts url")
		assert.Equal(r.URL.Query().Encode(), query, "it should be calling the correct query")
		fmt.Fprintf(w, "")
	}))
	defer svr.Close()

	//set the stub Server in the client
	clientF3.SetURL(svr.URL)

	//call
	clientAccount := client.NewAccountService(clientF3)
	status, err := clientAccount.Delete(id, version)

	//assert
	assert.Nil(err, "No error should be")
	assert.Equal(status, true, "If the account was delete, the status should be true")
}
