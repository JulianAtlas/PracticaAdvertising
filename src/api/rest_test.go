package main

import (
	"net/http"
	"testing"

	"bytes"
	"strconv"

	"github.com/stretchr/testify/assert"
	"github.com/PracticaAdvertising/src/api/rest"
)

func TestWhenCreateProductGetOkResponseCode(t *testing.T) {
	//setup
	router := rest.SetupRouter()
	asserter := assert.New(t)


	//operations

	var jsonStr = []byte(`{"Name":"pelota"}`)
	req, err := http.NewRequest("POST", ServerPath+"/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	//assserts

	resp := MakeRequestToRouter(router, req)
	asserter.Equal(http.StatusCreated, resp.HttpStatusCode)
	asserter.Equal(resp.Body, `{"Id":1,"Name":"pelota"}`)

}

func TestWhenCreateProductWithInvalidJsonThenResponseStatusIsBadRequest(t *testing.T) {
	//setup
	router := rest.SetupRouter()
	asserter := assert.New(t)

	//operations
	var jsonStr = []byte(`{"Name":1,"Id":"manuel"}`)
	req, err := http.NewRequest("POST", ServerPath+"/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")



	//asserts
	resp := MakeRequestToRouter(router, req)
	asserter.Equal(http.StatusBadRequest, resp.HttpStatusCode)
}

func TestWhenCreateProductWithNameEmptyTheResponseStatusIsBadRequest(t *testing.T) {
	//setup
	router := rest.SetupRouter()
	asserter := assert.New(t)

	//operations
	var jsonStr = []byte(`{"Name":""}`)
	req, err := http.NewRequest("POST", ServerPath+"/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	//asserts

	resp := MakeRequestToRouter(router, req)
	asserter.Equal(http.StatusBadRequest, resp.HttpStatusCode)
}

func TestWhenDeleteProductThatDoesntExistThenResponseStatusIsNotFound(t *testing.T) {
	//setup
	router := rest.SetupRouter()
	asserter := assert.New(t)

	//operations
	id := 6876
	req, err := http.NewRequest("DELETE", ServerPath+"/products"+"/"+strconv.Itoa(id), nil)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	//asserts

	resp := MakeRequestToRouter(router, req)
	asserter.Equal(http.StatusNotFound, resp.HttpStatusCode)
}

func TestWhenDeleteDoesntGetAndIntThenResponseStatusIsBadRequest(t *testing.T) {
	//setup
	router := rest.SetupRouter()
	asserter := assert.New(t)

	//operations
	req, err := http.NewRequest("DELETE", ServerPath+"/products"+"/"+"pelota", nil)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	//asserts

	resp := MakeRequestToRouter(router, req)
	asserter.Equal(http.StatusBadRequest, resp.HttpStatusCode)
}

func TestWhenDeleteAExistingProductThenResponseIsStatusOk(t *testing.T) {
	//setup
	router := rest.SetupRouter()
	asserter := assert.New(t)



	var jsonStr = []byte(`{"Name":"pelota"}`)
	req, err := http.NewRequest("POST", ServerPath+"/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	MakeRequestToRouter(router, req)

	// functtion

	req, err = http.NewRequest("DELETE", ServerPath+"/products"+"/"+"1", nil)
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	//asserts

	resp := MakeRequestToRouter(router, req)
	asserter.Equal(http.StatusOK, resp.HttpStatusCode)
}

func TestWhenListProductThenAllProductsAreReturned(t *testing.T) {
	//setup
	router := rest.SetupRouter()
	asserter := assert.New(t)

	var req *http.Request
	var req2 *http.Request
	var req3 *http.Request
	var err error
	//opertation

	var jsonStr = []byte(`{"Name":"pelota"}`)
	req, err = http.NewRequest("POST", ServerPath+"/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	MakeRequestToRouter(router, req)

	var jsonStr2 = []byte(`{"Name":"botella"}`)
	req2, err = http.NewRequest("POST", ServerPath+"/products", bytes.NewBuffer(jsonStr2))
	if err != nil {
		t.Error(err.Error())
	}
	req2.Header.Set("Content-Type", "application/json")
	resp := MakeRequestToRouter(router, req2)

	if resp.HttpStatusCode != http.StatusCreated {
		t.Error("No se creo el objeto")
	}

	req3, err = http.NewRequest("GET", ServerPath+"/products", nil)
	if err != nil {
		t.Error(err.Error())
	}
	resp = MakeRequestToRouter(router, req3)

	asserter.Equal(resp.Body, `{"1":{"Id":1,"Name":"pelota"},"2":{"Id":2,"Name":"botella"}}`)

}
