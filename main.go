package main

import (
	"fmt"

	f3 "github.com/ferchuhill/form3-client/client"
	model "github.com/ferchuhill/form3-client/client/models"
	"github.com/google/uuid"
)

func main() {

	// check API
	client := f3.NewClient()
	if client.IsHealthy() {
		fmt.Println("API is available")
	} else {
		fmt.Println("API is not available")
	}

	// Create
	acService := f3.NewAccountService(nil)

	var (
		uuidId  string
		uuidOrg string
		attr    model.AccountAttributes
	)
	uuidId = uuid.New().String()
	uuidOrg = uuid.New().String()
	country := "GB"
	names := []string{"Samantha Holder"}
	version := int64(0)

	attr = model.AccountAttributes{AccountNumber: "123123123123213", Country: &country, BankID: "400300", BankIDCode: "GBDSC", BaseCurrency: "GBP", Name: names}
	accountData := model.Account{Attributes: &attr, ID: uuidId, OrganisationID: uuidOrg, Type: "accounts", Version: version}

	// Create
	acData, err := acService.Create(accountData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Create : %s - %d\n", acData.ID, acData.Version)

	//Fetch
	acDataFectch, err := acService.Fetch(acData.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acDataFectch.ID)

	//Delete
	del, errDel := acService.Delete(acDataFectch.ID, int64(acDataFectch.Version))
	if errDel != nil {
		fmt.Println(errDel)
	} else {
		fmt.Println(del)
	}

}
