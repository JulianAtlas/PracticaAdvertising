package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"bytes"
	"io/ioutil"
	//"fmt"
	"fmt"
)

func TestWhenCreateProductPassGetOkResponseCode(t *testing.T) {
	asserter := assert.New(t)

	var jsonStr = []byte(`{"Name":"pelota"}`)
	req, err := http.NewRequest("POST", ServerPath+ "/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err.Error())
	}

	asserter.Equal(http.StatusCreated, resp.StatusCode)
	buff ,_:= ioutil.ReadAll(resp.Body)
	asserter.Equal(string(buff), `{"Id":1,"Name":"pelota"}`)

}

func TestWhenListProductThenAllProductsAreReturned(t *testing.T) {
	asserter := assert.New(t)

	var jsonStr = []byte(`{"Name":"pelota"}`)
	req, _ := http.NewRequest("POST", ServerPath+ "/products", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)

	var jsonStr2 = []byte(`{"Name":"botella"}`)
	req2, _ := http.NewRequest("POST", ServerPath+ "/products", bytes.NewBuffer(jsonStr2))
	req2.Header.Set("Content-Type", "application/json")
	resp,_ :=client.Do(req2)

	if resp.StatusCode != http.StatusCreated{
		t.Error("No se creo el objeto")
	}

	req3, _ := http.NewRequest("GET", ServerPath+ "/products", nil)
	resp,_ = client.Do(req3)

	fmt.Println(resp)
	buff ,_ := ioutil.ReadAll(resp.Body)

	asserter.Equal(string(buff), `{"1":{"Id":1,"Name":"pelota"},"2":{"Id":2,"Name":"botella"}}`)

}


