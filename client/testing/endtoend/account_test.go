package testingF3EndToEnd

import (
	"os"
	"testing"
	"time"

	"github.com/ferchuhill/form3-client/client"
	model "github.com/ferchuhill/form3-client/client/models"
	"github.com/stretchr/testify/assert"
)

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

// This test is for testing the comunication with the api, and verify if everithing is correct
func TestEndToEnd(t *testing.T) {

	// sleep to wait to the api is ready
	time.Sleep(3 * time.Second)

	assert := assert.New(t)

	//creaing a new client
	clientF3 := client.NewClient()
	url := os.Getenv("F3_API_URL")
	if url != "" {
		//setting the Docker url of the API
		clientF3.SetURL(url)
	}

	//Creating the new Account Service with the new client
	acService := client.NewAccountService(clientF3)

	//creating data to store
	names := []string{personName}
	attr := model.AccountAttributes{AccountNumber: accountNumber, Country: &country, BankID: bankId, BankIDCode: bankCode, BaseCurrency: baseCurrency, Name: names}
	account := model.Account{Attributes: &attr, ID: id, OrganisationID: idOrg, Type: types, Version: version}

	//create
	acData, errCreate := acService.Create(account)
	assert.Nil(errCreate)
	assert.Equal(acData.ID, id, "The ids should be equals")

	//Fetch
	acDataFectch, errFetch := acService.Fetch(acData.ID)
	assert.Nil(errFetch)
	assert.Equal(acData.ID, acDataFectch.ID, "The ids should be equals")

	//Delete
	del, errDel := acService.Delete(acDataFectch.ID, int64(acDataFectch.Version))
	assert.Nil(errDel, "No error should be")
	assert.Equal(del, true, "If the account was delete, the status should be true")

}
