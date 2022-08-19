package testingF3UnitTesting

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ferchuhill/form3-client/client"
	"github.com/stretchr/testify/assert"
)

// check if the constructor has the correct parameters
func TestConstructor(t *testing.T) {
	assert := assert.New(t)

	client := client.NewClient()
	assert.Equal(client.GetURL(), client.GetDefaultUrl(), "they should be equal")

	newUrl := "test.com"
	client.SetURL(newUrl)
	assert.Equal(client.GetURL(), newUrl, "they should be equal")
}

type DummyData struct {
	Data string `json:"data"`
}

// check if the call Api make a call in the stub server
func TestCallAPI(t *testing.T) {
	assert := assert.New(t)

	client := client.NewClient()

	//prepare
	expected := `{ "data": "dummy" }`
	dummyData := DummyData{}

	//stub server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	//set the stub Server in the client
	client.SetURL(svr.URL)

	//call
	err := client.CallAPI(http.MethodGet, "/test", nil, &dummyData)

	//assert
	assert.Nil(err, "No error should be")
	assert.Equal(dummyData.Data, "dummy")
}

// check the error when a parse error occurs by the stub server
func TestCallAPIErrorParse(t *testing.T) {
	assert := assert.New(t)

	client := client.NewClient()

	//prepare
	expected := `{ "error": "test"`
	dummyData := DummyData{}

	//stub server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	//set the stub Server in the client
	client.SetURL(svr.URL)

	//call
	err := client.CallAPI(http.MethodGet, "/test", nil, &dummyData)

	//assert
	assert.NotNil(err, "Error should have throw")
	assert.Contains(err.Error(), "unexpected end of JSON input")
}

// check when a empty response is recive by the stub server
func TestCallAPIEmptyResponse(t *testing.T) {
	assert := assert.New(t)

	client := client.NewClient()

	expected := `{ "error": "test" }`

	dummyData := DummyData{}
	dummyDataEmpty := DummyData{}
	//stub server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	client.SetURL(svr.URL)
	err := client.CallAPI(http.MethodGet, "/test", nil, &dummyData)

	assert.Nil(err, "No error should be")
	assert.Equal(dummyData, dummyDataEmpty, "should be equal to a empty state")
}

// check when the stub api return an error
func TestCallAPIErrorHTTP(t *testing.T) {
	assert := assert.New(t)

	client := client.NewClient()

	expected := `{ "message": "resourse not found" }`

	dummyData := DummyData{}
	//stub server
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	client.SetURL(svr.URL)
	err := client.CallAPI(http.MethodGet, "/test", nil, &dummyData)

	assert.NotNil(err, "Error should have throw")
	assert.Equal(err.Error(), `Error 500: "resourse not found"`, "an 500 error should be throw")
}
